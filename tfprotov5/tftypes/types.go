package tftypes

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
)

type unknown byte

type Type interface {
	Is(Type) bool
	String() string
	private()
}

type primitive string

func (p primitive) Is(t Type) bool {
	v, ok := t.(primitive)
	if !ok {
		return false
	}
	return p == v
}

func (p primitive) String() string {
	return "tftypes." + string(p)
}

func (p primitive) private() {}

const (
	UnknownValue      = unknown(0)
	UnknownType       = primitive("Unknown")
	DynamicPseudoType = primitive("DynamicPseudoType")
	String            = primitive("String")
	Number            = primitive("Number")
	Bool              = primitive("Bool")
)

var (
	_ Type = primitive("test")
)

type List struct {
	ElementType Type
}

func (l List) Is(t Type) bool {
	_, ok := t.(List)
	return ok
}

func (l List) String() string {
	return "tftypes.List"
}

func (l List) private() {}

type Set struct {
	ElementType Type
}

func (s Set) Is(t Type) bool {
	_, ok := t.(Set)
	return ok
}

func (s Set) String() string {
	return "tftypes.Set"
}

func (s Set) private() {}

type Map struct {
	AttributeType Type
}

func (m Map) Is(t Type) bool {
	_, ok := t.(Map)
	return ok
}

func (m Map) String() string {
	return "tftypes.Map"
}

func (m Map) private() {}

type Tuple struct {
	ElementTypes []Type
}

func (tu Tuple) Is(t Type) bool {
	_, ok := t.(Tuple)
	return ok
}

func (t Tuple) String() string {
	return "tftypes.Tuple"
}

func (t Tuple) private() {}

type Object struct {
	AttributeTypes map[string]Type
}

func (o Object) Is(t Type) bool {
	_, ok := t.(Object)
	return ok
}

func (o Object) String() string {
	return "tftypes.Object"
}

func (o Object) private() {}

type Unmarshaler interface {
	UnmarshalTerraform5Type(RawValue) error
}

type ErrUnhandledType string

func (e ErrUnhandledType) Error() string {
	return fmt.Sprintf("unhandled Terraform type %s", string(e))
}

// RawValue represents a form of a Terraform type that can be parsed into a Go
// type.
type RawValue struct {
	Type  Type
	Value interface{}
}

func (r RawValue) Unmarshal(dst interface{}) error {
	if unmarshaler, ok := dst.(Unmarshaler); ok {
		return unmarshaler.UnmarshalTerraform5Type(r)
	}
	switch {
	case r.Type.Is(String):
		if _, ok := dst.(*string); !ok {
			return fmt.Errorf("Can't unmarshal %s into %T", r.Type, dst)
		}
		*(dst.(*string)) = r.Value.(string)
	case r.Type.Is(Number):
		switch r.Value.(type) {
		case int64:
			switch dst.(type) {
			case *int64:
				*(dst.(*int64)) = r.Value.(int64)
			case *int32:
				if r.Value.(int64) > math.MaxInt32 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				if r.Value.(int64) < math.MinInt32 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*int32)) = int32(r.Value.(int64))
			case *int16:
				if r.Value.(int64) > math.MaxInt16 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				if r.Value.(int64) < math.MinInt16 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*int16)) = int16(r.Value.(int64))
			case *int8:
				if r.Value.(int64) > math.MaxInt8 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				if r.Value.(int64) < math.MinInt8 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*int8)) = int8(r.Value.(int64))
			case *int:
				// int types are only guaranteed to be able to
				// hold 32 bits; anything more is dependent on
				// the system. Because providers need to work
				// across architectures, we're going to ensure
				// that only the minimum is used here. Anyone
				// that needs more can use int64
				if r.Value.(int64) > math.MaxInt32 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't always fit in %T", r.Type, dst, r.Value, dst)
				}
				if r.Value.(int64) < math.MinInt32 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't always fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*int)) = int(r.Value.(int64))
			default:
				return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
			}
		case uint64:
			switch dst.(type) {
			case *uint64:
				*(dst.(*uint64)) = r.Value.(uint64)
			case *uint32:
				if r.Value.(uint64) > math.MaxUint32 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*uint32)) = uint32(r.Value.(uint64))
			case *uint16:
				if r.Value.(uint64) > math.MaxUint16 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*uint16)) = uint16(r.Value.(uint64))
			case *uint8:
				if r.Value.(uint64) > math.MaxUint8 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*uint8)) = uint8(r.Value.(uint64))
			case *uint:
				// uint types are only guaranteed to be able to
				// hold 32 bits; anything more is dependent on
				// the system. Because providers need to work
				// across architectures, we're going to ensure
				// that only the minimum is used here. Anyone
				// that needs more can use uint64
				if r.Value.(uint64) > math.MaxUint32 {
					return fmt.Errorf("can't unmarshal %s into %T: %d doesn't always fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*uint)) = uint(r.Value.(uint64))
			default:
				return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
			}
		case float64:
			switch dst.(type) {
			case *float64:
				*(dst.(*float64)) = r.Value.(float64)
			case *float32:
				if r.Value.(float64) > math.MaxFloat32 {
					return fmt.Errorf("can't unmarshal %s into %T: %f doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				if r.Value.(float64) < math.SmallestNonzeroFloat32 {
					return fmt.Errorf("can't unmarshal %s into %T: %f doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				*(dst.(*float32)) = float32(r.Value.(float64))
			default:
				return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
			}
		case big.Float:
			// TODO: handle big.Float values, like infinity or larger than 64 bit numbers
		}
	case r.Type.Is(Bool):
		if _, ok := dst.(*bool); !ok {
			return fmt.Errorf("Can't unmarshal %s into %T", r.Type, dst)
		}
		*(dst.(*bool)) = r.Value.(bool)
	case r.Type.Is(List{}):
		// this can't be a value with the cty information included; we
		// assume, at parsing time, that those are Tuples, not Lists.
		// So this _has_ to be a List, no ambiguity exists here.
		switch dst.(type) {
		case *[]interface{}:
			*(dst.(*[]interface{})) = r.Value.([]interface{})
		default:
			// TODO: return error
		}
	case r.Type.Is(Set{}):
		// this can't be a value with the cty information included; we
		// assume, at parsing time, that those are Tuples, not Sets.
		// So this _has_ to be a Set, no ambiguity exists here.
		switch dst.(type) {
		case *[]interface{}:
			*(dst.(*[]interface{})) = r.Value.([]interface{})
		default:
			// TODO: return error
		}
	case r.Type.Is(Tuple{}):
		// this could be an actual, honest-to-goodness Tuple, or it
		// could be a value with its cty information included. We only
		// know based on what the caller tells us they're expecting; an
		// []interface{} is expecting a Tuple, a RawValue is expecting
		// a *new* RawValue, with the Type property derived from the
		// cty information given to us.
		switch dst.(type) {
		case *[]interface{}:
			*(dst.(*[]interface{})) = r.Value.([]interface{})
		case *RawValue:
			val := r.Value.([]interface{})
			if len(val) != 2 {
				// TODO: return error
			}
			var typ interface{}
			err := json.Unmarshal([]byte(val[0].(string)), &typ)
			if err != nil {
				// TODO: return error
			}
			parsedType, err := parseType(typ)
			rv, err := rawValueFromComplexType(parsedType, val[1])
			if err != nil {
				return err
			}
			*(dst.(*RawValue)) = rv
		default:
			// TODO: return error
		}
	case r.Type.Is(Map{}):
		// this can't be a value with the cty information included; we
		// assume, at parsing time, that those are Objects, not Maps.
		// So this _has_ to be a Map, no ambiguity exists here.
		switch dst.(type) {
		case *map[string]interface{}:
			*(dst.(*map[string]interface{})) = r.Value.(map[string]interface{})
		default:
			// TODO: return error
		}
	case r.Type.Is(Object{}):
		// this could be an actual, honest-to-goodness Object, or it
		// could be a value with its cty information included. We only
		// know based on what the caller tells us they're expecting; a
		// map[string]interface{} is expecting an object, a RawValue is
		// expecting a *new* RawValue, with the Type property derived
		// from the cty information given to us.
		switch dst.(type) {
		case *map[string]interface{}:
			*(dst.(*map[string]interface{})) = r.Value.(map[string]interface{})
		case *RawValue:
			val := r.Value.(map[string]interface{})
			var typ interface{}
			err := json.Unmarshal([]byte(val["type"].(string)), &typ)
			if err != nil {
				// TODO: return error
			}
			parsedType, err := parseType(typ)
			if err != nil {
				// TODO: return error
			}
			rv, err := rawValueFromComplexType(parsedType, val["value"])
			if err != nil {
				return err
			}
			*(dst.(*RawValue)) = rv
		default:
			// TODO: return error
		}
	default:
		return ErrUnhandledType(r.Type.String())
	}
	return nil
}

func rawValueFromComplexType(typ Type, val interface{}) (RawValue, error) {
	if _, ok := typ.(primitive); ok {
		return RawValue{
			Type:  typ,
			Value: val,
		}, nil
	}
	switch {
	case typ.Is(List{}) || typ.Is(Set{}):
		var elementType Type
		if l, ok := typ.(List); ok {
			elementType = l.ElementType
		} else if s, ok := typ.(Set); ok {
			elementType = s.ElementType
		} else {
			// TODO: throw an error, this should never happen
		}
		values := make([]RawValue, len(val.([]interface{})))
		for pos, v := range val.([]interface{}) {
			value, err := rawValueFromComplexType(elementType, v)
			if err != nil {
				// TODO: return error
			}
			values[pos] = value
		}
		return RawValue{
			Type:  typ,
			Value: values,
		}, nil
	case typ.Is(Tuple{}):
		elementTypes := typ.(Tuple).ElementTypes
		if len(elementTypes) != len(val.([]interface{})) {
			// TODO: return error
		}
		elements := make([]RawValue, len(val.([]interface{})))
		for pos, v := range val.([]interface{}) {
			value, err := rawValueFromComplexType(elementTypes[pos], v)
			if err != nil {
				// TODO: return error
			}
			elements[pos] = value
		}
		return RawValue{
			Type:  typ,
			Value: elements,
		}, nil
	case typ.Is(Map{}):
		attributeType := typ.(Map).AttributeType
		values := make(map[string]RawValue, len(val.(map[string]interface{})))
		for key, v := range val.(map[string]interface{}) {
			value, err := rawValueFromComplexType(attributeType, v)
			if err != nil {
				// TODO: return error
			}
			values[key] = value
		}
		return RawValue{
			Type:  typ,
			Value: values,
		}, nil
	case typ.Is(Object{}):
		attributeTypes := typ.(Object).AttributeTypes
		values := make(map[string]RawValue, len(val.(map[string]interface{})))
		if len(attributeTypes) != len(val.(map[string]interface{})) {
			// TODO: return error
		}
		for key, v := range val.(map[string]interface{}) {
			attributeType, ok := attributeTypes[key]
			if !ok {
				// TODO: return error
			}
			value, err := rawValueFromComplexType(attributeType, v)
			if err != nil {
				// TODO: return error
			}
			values[key] = value
		}
		return RawValue{
			Type:  typ,
			Value: values,
		}, nil
	}
	// TODO: return error
	return RawValue{}, nil
}

func parseType(in interface{}) (Type, error) {
	switch v := in.(type) {
	// primitive types are represented just as strings,
	// with the type name in the string itself
	case string:
		switch v {
		case "string":
			return String, nil
		case "number":
			return Number, nil
		case "bool":
			return Bool, nil
		default:
			// TODO: return an error
		}
	case []interface{}:
		// sets, lists, tuples, maps, and objects are
		// represented as slices, recursive iterations of this
		// type/value syntax
		if len(v) < 1 {
			// TODO: return an error
		}
		switch v[0] {
		case "set":
			if len(v) != 2 {
				// TODO: return an error
			}
			subType, err := parseType(v[1])
			if err != nil {
				// TODO: return an error
			}
			return Set{
				ElementType: subType,
			}, nil
		case "list":
			if len(v) != 2 {
				// TODO: return an error
			}
			subType, err := parseType(v[1])
			if err != nil {
				// TODO: return an error
			}
			return List{
				ElementType: subType,
			}, nil
		case "tuple":
			if len(v) < 2 {
				// TODO: return an error
			}
			var types []Type
			for _, typ := range v {
				subType, err := parseType(typ)
				if err != nil {
					// TODO: return an error
				}
				types = append(types, subType)
			}
			return Tuple{
				ElementTypes: types,
			}, nil
		case "map":
			if len(v) != 2 {
				// TODO: return an error
			}
			subType, err := parseType(v[1])
			if err != nil {
				// TODO: return an error
			}
			return Map{
				AttributeType: subType,
			}, nil
		case "object":
			if len(v) < 2 {
				// TODO: return an error
			}
			types := map[string]Type{}
			valTypes := v[1].(map[string]interface{})
			for key, typ := range valTypes {
				subType, err := parseType(typ)
				if err != nil {
					// TODO: return an error
				}
				types[key] = subType
			}
			return Object{
				AttributeTypes: types,
			}, nil
		default:
			// TODO: return an error
			return nil, nil
		}
	}
	// TODO: return an error
	return nil, nil
}
