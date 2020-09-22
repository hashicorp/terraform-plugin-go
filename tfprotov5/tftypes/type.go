package tftypes

import "fmt"

type Type interface {
	Is(Type) bool
	String() string
	MarshalJSON() ([]byte, error)
	private()
}

func ParseType(in interface{}) (Type, error) {
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
			return nil, fmt.Errorf("unknown type %q", v)
		}
	case []interface{}:
		// sets, lists, tuples, maps, and objects are
		// represented as slices, recursive iterations of this
		// type/value syntax
		if len(v) < 1 {
			return nil, fmt.Errorf("improperly formatted type information; need at least %d elements, got %d", 1, len(v))
		}
		switch v[0] {
		case "set":
			if len(v) != 2 {
				return nil, fmt.Errorf("improperly formatted type information; need %d elements, got %d", 2, len(v))
			}
			subType, err := ParseType(v[1])
			if err != nil {
				return nil, fmt.Errorf("error parsing element type for tftypes.Set: %w", err)
			}
			return Set{
				ElementType: subType,
			}, nil
		case "list":
			if len(v) != 2 {
				return nil, fmt.Errorf("improperly formatted type information; need %d elements, got %d", 2, len(v))
			}
			subType, err := ParseType(v[1])
			if err != nil {
				return nil, fmt.Errorf("error parsing element type for tftypes.List: %w", err)
			}
			return List{
				ElementType: subType,
			}, nil
		case "tuple":
			if len(v) < 2 {
				return nil, fmt.Errorf("improperly formatted type information; need at least %d elements, got %d", 2, len(v))
			}
			var types []Type
			for pos, typ := range v {
				subType, err := ParseType(typ)
				if err != nil {
					return nil, fmt.Errorf("error parsing type of element %d for tftypes.Tuple: %w", pos, err)
				}
				types = append(types, subType)
			}
			return Tuple{
				ElementTypes: types,
			}, nil
		case "map":
			if len(v) != 2 {
				return nil, fmt.Errorf("improperly formatted type information; need %d elements, got %d", 2, len(v))
			}
			subType, err := ParseType(v[1])
			if err != nil {
				return nil, fmt.Errorf("error parsing attribute type for tftypes.Map: %w", err)
			}
			return Map{
				AttributeType: subType,
			}, nil
		case "object":
			if len(v) != 2 {
				return nil, fmt.Errorf("improperly formatted type information; need %d elements, got %d", 2, len(v))
			}
			types := map[string]Type{}
			valTypes := v[1].(map[string]interface{})
			for key, typ := range valTypes {
				subType, err := ParseType(typ)
				if err != nil {
					return nil, fmt.Errorf("error parsing type of attribute %q for tftypes.Object: %w", key, err)
				}
				types[key] = subType
			}
			return Object{
				AttributeTypes: types,
			}, nil
		default:
			return nil, fmt.Errorf("unknown type %q", v[0])
		}
	}
	return nil, fmt.Errorf("unhandled intermediary type %T", in)
}
