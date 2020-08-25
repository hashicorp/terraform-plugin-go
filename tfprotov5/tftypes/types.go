package tftypes

import (
	"fmt"
	"math"
	"math/big"
)

type unknown byte

const (
	UnknownValue      = unknown(0)
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
					return fmt.Errorf("can't unmarshal %s floato %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
				}
				if r.Value.(float64) < math.SmallestNonzeroFloat32 {
					return fmt.Errorf("can't unmarshal %s floato %T: %d doesn't fit in %T", r.Type, dst, r.Value, dst)
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
		// TODO: handle ambiguity of msgpack types
	case Set:
		// TODO: handle ambiguity of msgpack types
	case Tuple:
		// TODO: handle ambiguity of msgpack types
	case Map:
		// TODO: handle ambiguity of json types
	case Object:
		// TODO: handle ambiguity of json types
	}
	return ErrUnhandledType(r.Type)
}

// Type is used to identify Terraform types.
type Type string

func (t Type) String() string {
	return "tftypes." + string(t)
}
