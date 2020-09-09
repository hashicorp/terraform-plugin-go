package tftypes

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
)

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
			return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
		}
	case r.Type.Is(Set{}):
		// this can't be a value with the cty information included; we
		// assume, at parsing time, that those are Tuples, not Sets.
		// So this _has_ to be a Set, no ambiguity exists here.
		switch dst.(type) {
		case *[]interface{}:
			*(dst.(*[]interface{})) = r.Value.([]interface{})
		default:
			return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
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
			val, ok := r.Value.([]interface{})
			if !ok {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; invalid type information. Intermediary type is %T, not %T", r.Value, []interface{}{})
			}
			if len(val) != 2 {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; invalid type information. Expected %d items, got %d.", 2, len(val))
			}
			var typ interface{}
			str, ok := val[0].(string)
			if !ok {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; invalid type information. Expected %T for first value, got %T", str, val[0])
			}
			err := json.Unmarshal([]byte(str), &typ)
			if err != nil {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; invalid type information. First value %q isn't valid JSON: %w", str, err)
			}
			parsedType, err := parseType(typ)
			if err != nil {
				return fmt.Errorf("error parsing type: %w", err)
			}
			rv, err := rawValueFromComplexType(parsedType, val[1])
			if err != nil {
				return err
			}
			*(dst.(*RawValue)) = rv
		default:
			return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
		}
	case r.Type.Is(Map{}):
		// this can't be a value with the cty information included; we
		// assume, at parsing time, that those are Objects, not Maps.
		// So this _has_ to be a Map, no ambiguity exists here.
		switch dst.(type) {
		case *map[string]interface{}:
			*(dst.(*map[string]interface{})) = r.Value.(map[string]interface{})
		default:
			return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
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
			val, ok := r.Value.(map[string]interface{})
			if !ok {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; invalid type information. Intermediary type is %T, not %T", r.Value, map[string]interface{}{})
			}
			var typ interface{}
			typeInterface, ok := val["type"]
			if !ok {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; invalid type information. No \"type\" key found.")
			}
			typeStr, ok := typeInterface.(string)
			if !ok {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; invalid type information. Expected %T for \"type\" key, got %T", typeStr, typeInterface)
			}
			err := json.Unmarshal([]byte(typeStr), &typ)
			if err != nil {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; invalid type information. \"type\" key's value %q isn't valid JSON: %w", typeStr, err)
			}
			parsedType, err := parseType(typ)
			if err != nil {
				return fmt.Errorf("error parsing type: %w", err)
			}
			valueInterface, ok := val["value"]
			if !ok {
				return fmt.Errorf("can't unmarshal into tftypes.RawValue; no \"value\" key found")
			}
			rv, err := rawValueFromComplexType(parsedType, valueInterface)
			if err != nil {
				return fmt.Errorf("error parsing value: %w", err)
			}
			*(dst.(*RawValue)) = rv
		default:
			return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
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
			return RawValue{}, fmt.Errorf("type is %T, not %T or %T. This shouldn't be possible", typ, Set{}, List{})
		}
		interfaceValues, ok := val.([]interface{})
		if !ok {
			return RawValue{}, fmt.Errorf("unexpected intermediary type %T, expected %T", val, interfaceValues)
		}
		values := make([]RawValue, len(interfaceValues))
		for pos, v := range interfaceValues {
			value, err := rawValueFromComplexType(elementType, v)
			if err != nil {
				return RawValue{}, fmt.Errorf("error converting element %d of %T to a tftypes.RawValue: %w", pos, typ, err)
			}
			values[pos] = value
		}
		return RawValue{
			Type:  typ,
			Value: values,
		}, nil
	case typ.Is(Tuple{}):
		elementTypes := typ.(Tuple).ElementTypes
		interfaceValues, ok := val.([]interface{})
		if !ok {
			return RawValue{}, fmt.Errorf("unexpected intermediary type %T, expected %T", val, interfaceValues)
		}
		if len(elementTypes) != len(interfaceValues) {
			return RawValue{}, fmt.Errorf("expected %d element types in %T, got %d", len(interfaceValues), typ, len(elementTypes))
		}
		elements := make([]RawValue, len(interfaceValues))
		for pos, v := range interfaceValues {
			value, err := rawValueFromComplexType(elementTypes[pos], v)
			if err != nil {
				return RawValue{}, fmt.Errorf("error converting element %d of %T to a tftypes.RawValue: %w", pos, typ, err)
			}
			elements[pos] = value
		}
		return RawValue{
			Type:  typ,
			Value: elements,
		}, nil
	case typ.Is(Map{}):
		attributeType := typ.(Map).AttributeType
		msiValues, ok := val.(map[string]interface{})
		if !ok {
			return RawValue{}, fmt.Errorf("unexpected intermediary type %T, expected %T", val, msiValues)
		}
		values := make(map[string]RawValue, len(msiValues))
		for key, v := range msiValues {
			value, err := rawValueFromComplexType(attributeType, v)
			if err != nil {
				return RawValue{}, fmt.Errorf("error converting attribute %q of %T to a tftypes.RawValue: %w", key, typ, err)
			}
			values[key] = value
		}
		return RawValue{
			Type:  typ,
			Value: values,
		}, nil
	case typ.Is(Object{}):
		attributeTypes := typ.(Object).AttributeTypes
		msiValues, ok := val.(map[string]interface{})
		if !ok {
			return RawValue{}, fmt.Errorf("unexpected intermediary type %T, expected %T", val, msiValues)
		}
		values := make(map[string]RawValue, len(msiValues))
		if len(attributeTypes) != len(msiValues) {
			return RawValue{}, fmt.Errorf("expected %d attribute types in %T, got %d", len(msiValues), typ, len(attributeTypes))
		}
		for key, v := range msiValues {
			attributeType, ok := attributeTypes[key]
			if !ok {
				return RawValue{}, fmt.Errorf("expected type for attribute %q of %T, but didn't get one", key, typ)
			}
			value, err := rawValueFromComplexType(attributeType, v)
			if err != nil {
				return RawValue{}, fmt.Errorf("error converting attribute %q of %T to a tftypes.RawValue: %w", key, typ, err)
			}
			values[key] = value
		}
		return RawValue{
			Type:  typ,
			Value: values,
		}, nil
	}
	return RawValue{}, fmt.Errorf("unrecognized type %T can't be converted to a tftypes.RawValue", typ)
}
