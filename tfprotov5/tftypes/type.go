package tftypes

type Type interface {
	Is(Type) bool
	String() string
	private()
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
