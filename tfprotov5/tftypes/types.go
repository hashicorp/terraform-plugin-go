package tftypes

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
)

type unknown byte

const (
	UnknownValue      = unknown(0)
	UnknownType       = Type("Unknown")
	DynamicPseudoType = Type("DynamicPseudoType")
	String            = Type("String")
	Number            = Type("Number")
	Bool              = Type("Bool")
	List              = Type("List")
	Set               = Type("Set")
	Map               = Type("Map")
	Tuple             = Type("Tuple")
	Object            = Type("Object")
)

type Unmarshaler interface {
	UnmarshalTerraform5Type(RawValue) error
}

type ErrUnhandledType Type

func (e ErrUnhandledType) Error() string {
	return fmt.Sprintf("unhandled Terraform type %s", Type(e).String())
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
	switch r.Type {
	case String:
		if _, ok := dst.(*string); !ok {
			return fmt.Errorf("Can't unmarshal %s into %T", r.Type, dst)
		}
		*(dst.(*string)) = r.Value.(string)
	case Number:
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
	case Bool:
		if _, ok := dst.(*bool); !ok {
			return fmt.Errorf("Can't unmarshal %s into %T", r.Type, dst)
		}
		*(dst.(*bool)) = r.Value.(bool)
	case List:
		// this could be an actual, honest-to-goodness list, or it
		// could be a primitive with its cty information included. We
		// only know based on what the caller tells us they're
		// expecting; an []interface{} is  expecting a list, a RawValue
		// is expecting a *new* RawValue, with the Type property
		// derived from the cty information given to us.
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
			rv, err := rawValueFromComplexType(typ, val[1])
			if err != nil {
				return err
			}
			*(dst.(*RawValue)) = rv
		default:
			// TODO: return error
		}
	case Set:
		// this can only be something the user decided explicitly was a
		// Set; we can't get this from default behavior, so there's no
		// ambiguity here.
		switch dst.(type) {
		case *[]interface{}:
			*(dst.(*[]interface{})) = r.Value.([]interface{})
		default:
			// TODO: return error
		}
	case Tuple:
		// this can only be something the user decided explicitly was a
		// Tuple; we can't get this from default behavior, so there's
		// no ambiguity here.
		switch dst.(type) {
		case *[]interface{}:
			*(dst.(*[]interface{})) = r.Value.([]interface{})
		default:
			// TODO: return error
		}
	case Map:
		// this could be an actual, honest-to-goodness map, or it could
		// be a value with its cty information included. We only know
		// based on what the caller tells us they're expecting; a
		// map[string]interface{} is expecting a map, a RawValue is
		// expecting a *new* RawValue, with the TYpe property derived
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
			rv, err := rawValueFromComplexType(typ, val["value"])
			if err != nil {
				return err
			}
			*(dst.(*RawValue)) = rv
		default:
			// TODO: return error
		}
	case Object:
		// this can only be something the user decided explicitly was
		// an Object; we can't get this from default behavior, so
		// there's no ambiguity here.
		switch dst.(type) {
		case *map[string]interface{}:
			*(dst.(*map[string]interface{})) = r.Value.(map[string]interface{})
		default:
			// TODO: return error
		}
	default:
		return ErrUnhandledType(r.Type)
	}
	return nil
}

func rawValueFromComplexType(typ, val interface{}) (RawValue, error) {
	switch v := typ.(type) {
	// primitive types are represented just as strings,
	// with the type name in the string itself
	case string:
		switch v {
		case "string":
			return RawValue{
				Type:  String,
				Value: val,
			}, nil
		case "number":
			return RawValue{
				Type:  Number,
				Value: 0, // TODO: how to represent numbers?
			}, nil
		case "bool":
			return RawValue{
				Type:  Bool,
				Value: val,
			}, nil
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
			if len(v) < 2 {
				// TODO: return an error
			}
			var vals []RawValue
			for _, value := range val.([]interface{}) {
				rv, err := rawValueFromComplexType(v[1], value)
				if err != nil {
					// TODO: return an error
				}
				vals = append(vals, rv)
			}
			return RawValue{
				Type:  Set,
				Value: vals,
			}, nil
		case "list":
			if len(v) < 2 {
				// TODO: return an error
			}
			var vals []RawValue
			for _, value := range val.([]interface{}) {
				rv, err := rawValueFromComplexType(v[1], value)
				if err != nil {
					// TODO: return an error
				}
				vals = append(vals, rv)
			}
			return RawValue{
				Type:  List,
				Value: vals,
			}, nil
		case "tuple":
			if len(v) < len(val.([]interface{}))+1 {
				// TODO: return an error
			}
			var vals []RawValue
			for pos, value := range val.([]interface{}) {
				rv, err := rawValueFromComplexType(v[pos+1], value)
				if err != nil {
					// TODO: return an error
				}
				vals = append(vals, rv)
			}
			return RawValue{
				Type:  Tuple,
				Value: vals,
			}, nil
		case "map":
			if len(v) < 2 {
				// TODO: return an error
			}
			vals := map[string]RawValue{}
			for key, value := range val.(map[string]interface{}) {
				rv, err := rawValueFromComplexType(v[1], value)
				if err != nil {
					// TODO: return an error
				}
				vals[key] = rv
			}
			return RawValue{
				Type:  Map,
				Value: vals,
			}, nil
		case "object":
			if len(v) < 2 {
				// TODO: return an error
			}
			vals := map[string]RawValue{}
			valTypes := v[1].(map[string]interface{})
			if len(valTypes) != len(val.(map[string]interface{})) {
				// TODO: return an error
			}
			for key, value := range val.(map[string]interface{}) {
				typ, ok := valTypes[key]
				if !ok {
					// TODO: return error
				}
				rv, err := rawValueFromComplexType(typ, value)
				if err != nil {
					// TODO: return an error
				}
				vals[key] = rv
			}
			return RawValue{
				Type:  Object,
				Value: vals,
			}, nil
		default:
		}
	}
	// TODO: return error?
	return RawValue{}, nil
}

// Type is used to identify Terraform types.
type Type string

func (t Type) String() string {
	return "tftypes." + string(t)
}
