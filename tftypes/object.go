package tftypes

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// Object is a Terraform type representing an unordered collection of
// attributes, potentially of differing types, each identifiable with a unique
// string name. The number of attributes, their names, and their types are part
// of the type signature for the Object, and so two Objects with different
// attribute names or types are considered to be distinct types.
type Object struct {
	AttributeTypes map[string]Type

	// used to make this type uncomparable
	// see https://golang.org/ref/spec#Comparison_operators
	// this enforces the use of Is, instead
	_ []struct{}
}

// Is returns whether `t` is an Object type or not. If `t` is an instance of
// the Object type and its AttributeTypes property is not nil, it will only
// return true the AttributeTypes are considered the same. To be considered
// equal, the same set of keys must be present in each, and each key's value
// needs to be considered the same type between the two Objects.
func (o Object) Is(t Type) bool {
	v, ok := t.(Object)
	if !ok {
		return false
	}
	if v.AttributeTypes != nil {
		if len(o.AttributeTypes) != len(v.AttributeTypes) {
			return false
		}
		for k, typ := range o.AttributeTypes {
			if _, ok := v.AttributeTypes[k]; !ok {
				return false
			}
			if !typ.Is(v.AttributeTypes[k]) {
				return false
			}
		}
	}
	return ok
}

func (o Object) String() string {
	var res strings.Builder
	res.WriteString("tftypes.Object[")
	keys := make([]string, 0, len(o.AttributeTypes))
	for k := range o.AttributeTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for pos, key := range keys {
		if pos != 0 {
			res.WriteString(", ")
		}
		res.WriteString(`"` + key + `":`)
		res.WriteString(o.AttributeTypes[key].String())
	}
	res.WriteString("]")
	return res.String()
}

func (o Object) private() {}

func (o Object) supportedGoTypes() []string {
	return []string{"map[string]tftypes.Value"}
}

func valueCanBeObject(val interface{}) bool {
	switch val.(type) {
	case map[string]Value:
		return true
	default:
		return false
	}
}

func valueFromObject(types map[string]Type, in interface{}) (Value, error) {
	switch value := in.(type) {
	case map[string]Value:
		// types should only be null if the "Object" is actually a
		// DynamicPseudoType being created from a map[string]Value. In
		// which case, we don't know what types it should have, or even
		// how many there will be, so let's not validate that at all
		if types != nil {
			for k := range types {
				if _, ok := value[k]; !ok {
					return Value{}, fmt.Errorf("can't create a tftypes.Value of type %s, required attribute %q not set", Object{AttributeTypes: types}, k)
				}
			}
			for k, v := range value {
				typ, ok := types[k]
				if !ok {
					return Value{}, fmt.Errorf("can't set a value on %q in tftypes.NewValue, key not part of the object type %s.", k, Object{AttributeTypes: types})
				}
				if !v.Type().Is(types[k]) && !types[k].Is(DynamicPseudoType) {
					return Value{}, fmt.Errorf("tftypes.NewValue can't use type %s as a value for %q in %s. Expected type is %s.", v.Type(), k, Object{AttributeTypes: types}, typ)
				}
			}
		}
		return Value{
			typ:   Object{AttributeTypes: types},
			value: value,
		}, nil
	default:
		return Value{}, fmt.Errorf("tftypes.NewValue can't use %T as a tftypes.Object. Expected types are: %s", in, formattedSupportedGoTypes(Object{}))
	}
}

// MarshalJSON returns a JSON representation of the full type signature of `o`,
// including the AttributeTypes.
//
// Deprecated: this is not meant to be called by third-party code.
func (o Object) MarshalJSON() ([]byte, error) {
	attrs, err := json.Marshal(o.AttributeTypes)
	if err != nil {
		return nil, err
	}
	return []byte(`["object",` + string(attrs) + `]`), nil
}
