package tfprotov5

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func jsonByteDecoder(buf []byte) *json.Decoder {
	r := bytes.NewReader(buf)
	dec := json.NewDecoder(r)
	dec.UseNumber()
	return dec
}

func jsonUnmarshal(buf []byte, typ tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)

	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}

	if tok == nil {
		return tftypes.NewValue(typ, nil), nil
	}

	switch {
	case typ.Is(tftypes.String):
		return jsonUnmarshalString(buf, typ, p)
	case typ.Is(tftypes.Number):
		return jsonUnmarshalNumber(buf, typ, p)
	case typ.Is(tftypes.Bool):
		return jsonUnmarshalBool(buf, typ, p)
	case typ.Is(tftypes.DynamicPseudoType):
		return jsonUnmarshalDynamicPseudoType(buf, typ, p)
	case typ.Is(tftypes.List{}):
		return jsonUnmarshalList(buf, typ.(tftypes.List).ElementType, p)
	case typ.Is(tftypes.Set{}):
		return jsonUnmarshalSet(buf, typ.(tftypes.Set).ElementType, p)

	case typ.Is(tftypes.Map{}):
		return jsonUnmarshalMap(buf, typ.(tftypes.Map).AttributeType, p)
	case typ.Is(tftypes.Tuple{}):
		return jsonUnmarshalTuple(buf, typ.(tftypes.Tuple).ElementTypes, p)
	case typ.Is(tftypes.Object{}):
		return jsonUnmarshalObject(buf, typ.(tftypes.Object).AttributeTypes, p)
	}
	return tftypes.Value{}, p.NewErrorf("unknown type %s", typ)
}

func jsonUnmarshalString(buf []byte, typ tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)

	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	switch v := tok.(type) {
	case string:
		return tftypes.NewValue(tftypes.String, v), nil
	case json.Number:
		return tftypes.NewValue(tftypes.String, string(v)), nil
	case bool:
		// TODO: convert boolean to a string
		// not really sure why, but... compatibility!
	}
	return tftypes.Value{}, p.NewErrorf("unsupported type %T sent as %s", tok, tftypes.String)
}

func jsonUnmarshalNumber(buf []byte, typ tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)

	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	numTok, ok := tok.(json.Number)
	if !ok {
		return tftypes.Value{}, p.NewErrorf("unsupported type %T sent as %s", tok, tftypes.Number)
	}
	// TODO: convert numTok to big.Float
	return tftypes.NewValue(typ, numTok), nil
}

func jsonUnmarshalBool(buf []byte, typ tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)
	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	switch v := tok.(type) {
	case bool:
		return tftypes.NewValue(tftypes.Bool, v), nil
	case string:
		// TODO: convert string to boolean
	}
	return tftypes.Value{}, p.NewErrorf("unsupported type %T sent as %s", tok, tftypes.Bool)
}

func jsonUnmarshalDynamicPseudoType(buf []byte, typ tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)
	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('{') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('{'), tok)
	}
	var t tftypes.Type
	var valBody []byte
	for dec.More() {
		tok, err = dec.Token()
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
		}
		key, ok := tok.(string)
		if !ok {
			return tftypes.Value{}, p.NewErrorf("expected key to be a string, got %T", tok)
		}
		var rawVal json.RawMessage
		err = dec.Decode(&rawVal)
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error decoding value: %w", err)
		}
		switch key {
		case "type":
			t, err = parseJSONType(rawVal)
			if err != nil {
				return tftypes.Value{}, p.NewErrorf("error decoding type information: %w", err)
			}
		case "value":
			valBody = rawVal
		default:
			return tftypes.Value{}, p.NewErrorf("invalid key %q in dynamically-typed value", key)
		}
	}
	tok, err = dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('}') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('}'), tok)
	}
	if t == nil {
		return tftypes.Value{}, p.NewErrorf("missing type in dynamically-typed value")
	}
	if valBody == nil {
		return tftypes.Value{}, p.NewErrorf("missing value in dynamically-typed value")
	}
	return jsonUnmarshal(valBody, t, p)
}

func jsonUnmarshalList(buf []byte, elementType tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)

	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('[') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('['), tok)
	}

	// we want to have a value for this always, even if there are no
	// elements, because no elements is *technically* different than empty,
	// and we want to preserve that distinction
	//
	// var vals []tftypes.Value
	// would evaluate as nil if the list is empty
	//
	// while generally in Go it's undesirable to treat empty and nil slices
	// separately, in this case we're surfacing a non-Go-in-origin
	// distinction, so we'll allow it.
	vals := []tftypes.Value{}

	// add a placeholder at the end of the path
	// we'll fix this in each part of the loop to have the right value
	// we can't just append in the loop, we need to replace, or we'll
	// be adding, not modifying, the last part of the path
	p = append(p, nil)
	var idx int64
	for dec.More() {
		// correct the last value in the path
		p[len(p)-1] = idx
		// update the index
		idx++

		var rawVal json.RawMessage
		err = dec.Decode(&rawVal)
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error decoding value: %w", err)
		}
		val, err := jsonUnmarshal(rawVal, elementType, p)
		if err != nil {
			return tftypes.Value{}, err
		}
		vals = append(vals, val)
	}
	// drop the last value, we're out of the loop now
	p = p.RemoveLastStep()

	tok, err = dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim(']') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim(']'), tok)
	}
	return tftypes.NewValue(tftypes.List{
		ElementType: elementType,
	}, vals), nil
}

func jsonUnmarshalSet(buf []byte, elementType tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)

	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('[') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('['), tok)
	}

	// we want to have a value for this always, even if there are no
	// elements, because no elements is *technically* different than empty,
	// and we want to preserve that distinction
	//
	// var vals []tftypes.Value
	// would evaluate as nil if the set is empty
	//
	// while generally in Go it's undesirable to treat empty and nil slices
	// separately, in this case we're surfacing a non-Go-in-origin
	// distinction, so we'll allow it.
	vals := []tftypes.Value{}

	p = p.AddValueStep(tftypes.NewValue(elementType, tftypes.UnknownValue))
	for dec.More() {
		var rawVal json.RawMessage
		err = dec.Decode(&rawVal)
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error decoding value: %w", err)
		}
		val, err := jsonUnmarshal(rawVal, elementType, p)
		if err != nil {
			return tftypes.Value{}, err
		}
		vals = append(vals, val)
	}
	// drop the last value, we're out of the loop now
	p = p.RemoveLastStep()

	tok, err = dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim(']') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim(']'), tok)
	}
	return tftypes.NewValue(tftypes.Set{
		ElementType: elementType,
	}, vals), nil
}

func jsonUnmarshalMap(buf []byte, attrType tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)

	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('{') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('{'), tok)
	}

	vals := map[string]tftypes.Value{}
	p = p.AddValueStep(tftypes.NewValue(attrType, tftypes.UnknownValue))
	for dec.More() {
		tok, err := dec.Token()
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
		}
		key, ok := tok.(string)
		if !ok {
			return tftypes.Value{}, p.NewErrorf("expected map key to be a string, got %T", tok)
		}

		//fix the path value, we have an actual key now
		p[len(p)-1] = key

		var rawVal json.RawMessage
		err = dec.Decode(&rawVal)
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error decoding value: %w", err)
		}
		val, err := jsonUnmarshal(rawVal, attrType, p)
		if err != nil {
			return tftypes.Value{}, err
		}
		vals[key] = val
	}
	// drop the last value, we're out of the loop now
	p = p.RemoveLastStep()

	tok, err = dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('}') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('}'), tok)
	}

	return tftypes.NewValue(tftypes.Map{
		AttributeType: attrType,
	}, vals), nil
}

func jsonUnmarshalTuple(buf []byte, elementTypes []tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)

	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('[') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('['), tok)
	}

	// we want to have a value for this always, even if there are no
	// elements, because no elements is *technically* different than empty,
	// and we want to preserve that distinction
	//
	// var vals []tftypes.Value
	// would evaluate as nil if the tuple is empty
	//
	// while generally in Go it's undesirable to treat empty and nil slices
	// separately, in this case we're surfacing a non-Go-in-origin
	// distinction, so we'll allow it.
	vals := []tftypes.Value{}

	// add a placeholder at the end of the path
	// we'll fix this in each part of the loop to have the right value
	// we can't just append in the loop, we need to replace, or we'll
	// be adding, not modifying, the last part of the path
	p = append(p, nil)
	var idx int64
	for dec.More() {
		if idx >= int64(len(elementTypes)) {
			return tftypes.Value{}, p[:len(p)-1].NewErrorf("too many tuple elements (only have types for %d)", len(elementTypes))
		}

		p[len(p)-1] = idx
		elementType := elementTypes[idx]
		idx++

		var rawVal json.RawMessage
		err = dec.Decode(&rawVal)
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error decoding value: %w", err)
		}
		val, err := jsonUnmarshal(rawVal, elementType, p)
		if err != nil {
			return tftypes.Value{}, err
		}
		vals = append(vals, val)
	}
	//drop the last value, we're out of the loop now
	p = p.RemoveLastStep()

	tok, err = dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim(']') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim(']'), tok)
	}

	if len(vals) != len(elementTypes) {
		return tftypes.Value{}, p.NewErrorf("not enough tuple elements (only have %d, have types for %d)", len(vals), len(elementTypes))
	}

	return tftypes.NewValue(tftypes.Tuple{
		ElementTypes: elementTypes,
	}, vals), nil
}

func jsonUnmarshalObject(buf []byte, attrTypes map[string]tftypes.Type, p tftypes.Path) (tftypes.Value, error) {
	dec := jsonByteDecoder(buf)

	tok, err := dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('{') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('{'), tok)
	}

	vals := map[string]tftypes.Value{}
	// placeholder for the attributes we're about to loop through
	p = append(p, nil)

	for dec.More() {
		tok, err := dec.Token()
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
		}
		key, ok := tok.(string)
		if !ok {
			return tftypes.Value{}, p.NewErrorf("object attribute key was %T, not string", tok)
		}
		attrType, ok := attrTypes[key]
		if !ok {
			return tftypes.Value{}, p.NewErrorf("unsupported attribute %q", key)
		}

		p[len(p)-1] = key

		var rawVal json.RawMessage
		err = dec.Decode(&rawVal)
		if err != nil {
			return tftypes.Value{}, p.NewErrorf("error decoding value: %w", err)
		}
		val, err := jsonUnmarshal(rawVal, attrType, p)
		if err != nil {
			return tftypes.Value{}, err
		}
		vals[key] = val
	}
	// we're out of the loop, drop the key from our path
	p = p.RemoveLastStep()

	tok, err = dec.Token()
	if err != nil {
		return tftypes.Value{}, p.NewErrorf("error reading token: %w", err)
	}
	if tok != json.Delim('}') {
		return tftypes.Value{}, p.NewErrorf("invalid JSON, expected %q, got %q", json.Delim('}'), tok)
	}

	// make sure we have a value for every attribute
	for k, typ := range attrTypes {
		if _, ok := vals[k]; !ok {
			vals[k] = tftypes.NewValue(typ, nil)
		}
	}

	return tftypes.NewValue(tftypes.Object{
		AttributeTypes: attrTypes,
	}, vals), nil
}

type jsonType struct {
	t tftypes.Type
}

func parseJSONType(buf []byte) (tftypes.Type, error) {
	var t jsonType
	err := json.Unmarshal(buf, &t)
	return t.t, err
}

func (t *jsonType) UnmarshalJSON(buf []byte) error {
	r := bytes.NewReader(buf)
	dec := json.NewDecoder(r)

	tok, err := dec.Token()
	if err != nil {
		return err
	}

	switch v := tok.(type) {
	case string:
		switch v {
		case "bool":
			t.t = tftypes.Bool
		case "number":
			t.t = tftypes.Number
		case "string":
			t.t = tftypes.String
		case "dynamic":
			t.t = tftypes.DynamicPseudoType
		default:
			return fmt.Errorf("invalid primitive type name %q", v)
		}

		if dec.More() {
			return fmt.Errorf("extraneous data after type description")
		}
		return nil
	case json.Delim:
		if rune(v) != '[' {
			return fmt.Errorf("invalid complex type description")
		}

		tok, err = dec.Token()
		if err != nil {
			return err
		}

		kind, ok := tok.(string)
		if !ok {
			return fmt.Errorf("invalid complex type kind name")
		}

		switch kind {
		case "list":
			var ety jsonType
			err = dec.Decode(&ety)
			if err != nil {
				return err
			}
			t.t = tftypes.List{
				ElementType: ety.t,
			}
		case "map":
			var ety jsonType
			err = dec.Decode(&ety)
			if err != nil {
				return err
			}
			t.t = tftypes.Map{
				AttributeType: ety.t,
			}
		case "set":
			var ety jsonType
			err = dec.Decode(&ety)
			if err != nil {
				return err
			}
			t.t = tftypes.Set{
				ElementType: ety.t,
			}
		case "object":
			var atys map[string]jsonType
			err = dec.Decode(&atys)
			if err != nil {
				return err
			}
			types := make(map[string]tftypes.Type, len(atys))
			for k, v := range atys {
				types[k] = v.t
			}
			t.t = tftypes.Object{
				AttributeTypes: types,
			}
		case "tuple":
			var etys []jsonType
			err = dec.Decode(&etys)
			if err != nil {
				return err
			}
			types := make([]tftypes.Type, 0, len(etys))
			for _, ty := range etys {
				types = append(types, ty.t)
			}
			t.t = tftypes.Tuple{
				ElementTypes: types,
			}
		default:
			return fmt.Errorf("invalid complex type kind name")
		}

		tok, err = dec.Token()
		if err != nil {
			return err
		}
		if delim, ok := tok.(json.Delim); !ok || rune(delim) != ']' || dec.More() {
			return fmt.Errorf("unexpected extra data in type description")
		}

		return nil

	default:
		return fmt.Errorf("invalid type description")
	}
}
