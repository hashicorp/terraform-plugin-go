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
		//typ, val := standardizeRawValue(r)
		var typ Type
		var val interface{}
		return unmarshaler.UnmarshalTerraform5Type(RawValue{
			Type:  typ,
			Value: val,
		})
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
			err := unmarshalInt64(r.Value.(int64), dst)
			if err != nil {
				return fmt.Errorf("can't unmarshal %s into %T: %w", r.Type, dst, err)
			}
			return nil
		case uint64:
			err := unmarshalUint64(r.Value.(uint64), dst)
			if err != nil {
				return fmt.Errorf("can't unmarshal %s into %T: %w", r.Type, dst, err)
			}
			return nil
		case float64:
			err := unmarshalFloat64(r.Value.(float64), dst)
			if err != nil {
				return fmt.Errorf("can't unmarshal %s into %T: %w", r.Type, dst, err)
			}
		case *big.Float:
			switch dst.(type) {
			case *big.Float:
				dst.(*big.Float).Set(r.Value.(*big.Float))
			case *uint64, *uint32, *uint16, *uint8, *uint:
				if !r.Value.(*big.Float).IsInt() {
					return fmt.Errorf("can't unmarshal %s into %T: is not an integer", r.Type, dst)
				}
				if r.Value.(*big.Float).Cmp(big.NewFloat(math.MaxUint64)) == 1 {
					return fmt.Errorf("can't unmarshal %s into %T: value too large", r.Type, dst)
				}
				if r.Value.(*big.Float).Cmp(big.NewFloat(0)) == -1 {
					return fmt.Errorf("can't unmarshal %s into %T: value too small", r.Type, dst)
				}
				v, _ := r.Value.(*big.Float).Uint64()
				err := unmarshalUint64(v, dst)
				if err != nil {
					return fmt.Errorf("can't unmarshal %s into %T: %w", r.Type, dst, err)
				}
				return nil
			case *int64, *int32, *int16, *int8, int:
				if !r.Value.(*big.Float).IsInt() {
					return fmt.Errorf("can't unmarshal %s into %T: is not an integer", r.Type, dst)
				}
				if r.Value.(*big.Float).Cmp(big.NewFloat(math.MaxInt64)) == 1 {
					return fmt.Errorf("can't unmarshal %s into %T: value too large", r.Type, dst)
				}
				if r.Value.(*big.Float).Cmp(big.NewFloat(math.MinInt64)) == -1 {
					return fmt.Errorf("can't unmarshal %s into %T: value too small", r.Type, dst)
				}
				v, _ := r.Value.(*big.Float).Int64()
				err := unmarshalInt64(v, dst)
				if err != nil {
					return fmt.Errorf("can't unmarshal %s into %T: %w", r.Type, dst, err)
				}
				return nil
			case *float64, *float32:
				if r.Value.(*big.Float).Cmp(big.NewFloat(math.MaxFloat64)) == 1 {
					return fmt.Errorf("can't unmarshal %s into %T: value too large", r.Type, dst)
				}
				if r.Value.(*big.Float).Cmp(big.NewFloat(math.SmallestNonzeroFloat64)) == -1 {
					return fmt.Errorf("can't unmarshal %s into %T: value too small", r.Type, dst)
				}
				v, _ := r.Value.(*big.Float).Float64()
				err := unmarshalFloat64(v, dst)
				if err != nil {
					return fmt.Errorf("can't unmarshal %s into %T: %w", r.Type, dst, err)
				}
				return nil
			default:
				return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
			}
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
			var res []interface{}
			for _, i := range r.Value.([]interface{}) {
				res = append(res, RawValue{
					Type:  r.Type.(List).ElementType,
					Value: i,
				})
			}
			*(dst.(*[]interface{})) = res
		default:
			return fmt.Errorf("can't unmarshal %s into %T", r.Type, dst)
		}
	case r.Type.Is(Set{}):
		// this can't be a value with the cty information included; we
		// assume, at parsing time, that those are Tuples, not Sets.
		// So this _has_ to be a Set, no ambiguity exists here.
		switch dst.(type) {
		case *[]interface{}:
			var res []interface{}
			for _, i := range r.Value.([]interface{}) {
				res = append(res, RawValue{
					Type:  r.Type.(Set).ElementType,
					Value: i,
				})
			}
			*(dst.(*[]interface{})) = res
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
			parsedType, err := ParseType(typ)
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
			parsedType, err := ParseType(typ)
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

func unmarshalUint64(in uint64, dst interface{}) error {
	switch v := dst.(type) {
	case *uint64:
		*v = in
	case *uint32:
		if in > math.MaxUint32 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		*v = uint32(in)
	case *uint16:
		if in > math.MaxUint16 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		*v = uint16(in)
	case *uint8:
		if in > math.MaxUint8 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		*v = uint8(in)
	case *uint:
		// uint types are only guaranteed to be able to hold 32 bits;
		// anything more is dependent on the system. Because providers
		// need to work across architectures, we're going to ensure
		// that only the minimum is used here. Anyone that needs more
		// can use uint64
		if in > math.MaxUint32 {
			return fmt.Errorf("%d doesn't always fit in %T", in, dst)
		}
		*v = uint(in)
	default:
		return fmt.Errorf("can't unmarshal uint64 into %T", dst)
	}
	return nil
}

func unmarshalInt64(in int64, dst interface{}) error {
	switch v := dst.(type) {
	case *int64:
		*v = in
	case *int32:
		if in > math.MaxInt32 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		if in < math.MinInt32 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		*v = int32(in)
	case *int16:
		if in > math.MaxInt16 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		if in < math.MinInt16 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		*v = int16(in)
	case *int8:
		if in > math.MaxInt8 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		if in < math.MinInt8 {
			return fmt.Errorf("%d doesn't fit in %T", in, dst)
		}
		*v = int8(in)
	case *int:
		// int types are only guaranteed to be able to hold 32 bits;
		// anything more is dependent on the system. Because providers
		// need to work across architectures, we're going to ensure
		// that only the minimum is used here. Anyone that needs more
		// can use int64
		if in > math.MaxInt32 {
			return fmt.Errorf("%d doesn't always fit in %T", in, dst)
		}
		if in < math.MinInt32 {
			return fmt.Errorf("%d doesn't always fit in %T", in, dst)
		}
		*v = int(in)
	default:
		return fmt.Errorf("can't unmarshal int64 into %T", dst)
	}
	return nil
}

func unmarshalFloat64(in float64, dst interface{}) error {
	switch v := dst.(type) {
	case *float64:
		*v = in
	case *float32:
		if in > math.MaxFloat32 {
			return fmt.Errorf("%f doesn't fit in %T", in, dst)
		}
		if in < math.SmallestNonzeroFloat32 {
			return fmt.Errorf("%f doesn't fit in %T", in, dst)
		}
		*v = float32(in)
	default:
		return fmt.Errorf("can't unmarshal float64 into %T", dst)
	}
	return nil
}

func rawValueFromComplexType(typ Type, val interface{}) (RawValue, error) {
	if _, ok := typ.(primitive); ok {
		if typ.Is(Number) {
			result := RawValue{
				Type: Number,
			}
			switch v := val.(type) {
			case *big.Float:
				v.Set(v)
			case uint64:
				result.Value = big.NewFloat(float64(v))
			case uint32:
				result.Value = big.NewFloat(float64(v))
			case uint16:
				result.Value = big.NewFloat(float64(v))
			case uint8:
				result.Value = big.NewFloat(float64(v))
			case uint:
				result.Value = big.NewFloat(float64(v))
			case int64:
				result.Value = big.NewFloat(float64(v))
			case int32:
				result.Value = big.NewFloat(float64(v))
			case int16:
				result.Value = big.NewFloat(float64(v))
			case int8:
				result.Value = big.NewFloat(float64(v))
			case int:
				result.Value = big.NewFloat(float64(v))
			case float64:
				result.Value = big.NewFloat(v)
			case float32:
				result.Value = big.NewFloat(float64(v))
			default:
				return result, fmt.Errorf("can't use type %T as a %s", val, Number)
			}
			return result, nil
		}
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
