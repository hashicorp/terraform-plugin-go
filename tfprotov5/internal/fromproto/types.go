package fromproto

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/vmihailenco/msgpack"
	msgpackCodes "github.com/vmihailenco/msgpack/codes"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

var ErrUnknownDynamicValueType = errors.New("DynamicValue had no JSON or msgpack data set")

func TerraformTypesRawValue(in tfplugin5.DynamicValue) (tftypes.RawValue, error) {
	if len(in.Msgpack) > 0 {
		return msgpackToRawValue(in.Msgpack)
	}
	if len(in.Json) > 0 {
		return jsonToRawValue(in.Json)
	}
	return tftypes.RawValue{}, ErrUnknownDynamicValueType
}

func msgpackToRawValue(in []byte) (tftypes.RawValue, error) {
	r := bytes.NewReader(in)
	dec := msgpack.NewDecoder(r)

	peek, err := dec.PeekCode()
	if err != nil {
		return tftypes.RawValue{}, err
	}
	if msgpackCodes.IsExt(peek) {
		// We just assume _all_ extensions are unknown values,
		// since we don't have any other extensions.
		err = dec.Skip() // skip what we've peeked
		if err != nil {
			return tftypes.RawValue{}, err
		}
		return tftypes.RawValue{
			// we don't know what type this is yet, the caller
			// decides
			Type:  tftypes.UnknownType,
			Value: tftypes.UnknownValue,
		}, nil
	}
	// if the caller wants this to be dynamic, we unmarshalDynamic

	if peek == msgpackCodes.Nil {
		err = dec.Skip() // skip what we've peeked
		if err != nil {
			return tftypes.RawValue{}, err
		}
		return tftypes.RawValue{
			// we don't know what type this is yet, the caller
			// decides
			Type:  tftypes.UnknownType,
			Value: nil,
		}, nil
	}

	val, err := dec.DecodeInterface()
	if err != nil {
		return tftypes.RawValue{}, err
	}

	t, err := typeFromValue(val)
	if err != nil {
		return tftypes.RawValue{}, err
	}

	return tftypes.RawValue{
		Type:  t,
		Value: val,
	}, nil
}

func jsonToRawValue(in []byte) (tftypes.RawValue, error) {
	r := bytes.NewReader(in)
	dec := json.NewDecoder(r)
	dec.UseNumber()

	var result interface{}
	if err := dec.Decode(&result); err != nil {
		return tftypes.RawValue{}, err
	}
	if dec.More() {
		return tftypes.RawValue{}, errors.New("more than one JSON element to decode")
	}

	t, err := typeFromValue(result)
	if err != nil {
		return tftypes.RawValue{}, err
	}
	return tftypes.RawValue{
		Type:  t,
		Value: result,
	}, nil
}

func typeFromValue(in interface{}) (tftypes.Type, error) {
	if in == nil {
		return tftypes.UnknownType, nil
	}
	switch v := in.(type) {
	case string:
		return tftypes.String, nil
	case json.Number, int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		return tftypes.Number, nil
	case bool:
		return tftypes.Bool, nil
	case []interface{}:
		var types []tftypes.Type
		for pos, val := range v {
			typ, err := typeFromValue(val)
			if err != nil {
				return nil, fmt.Errorf("error choosing default tftypes.Type for value %d of type %T: %w", pos, in, err)
			}
			types = append(types, typ)
		}
		return tftypes.Tuple{
			ElementTypes: types,
		}, nil
	case map[string]interface{}:
		types := map[string]tftypes.Type{}
		for k, val := range v {
			typ, err := typeFromValue(val)
			if err != nil {
				return nil, fmt.Errorf("error choosing default tftypes.Type for attribute %q of type %T: %w", k, in, err)
			}
			types[k] = typ
		}
		return tftypes.Object{
			AttributeTypes: types,
		}, nil
	}
	return tftypes.UnknownType, fmt.Errorf("Go type %T has no default tftypes.Type", in)
}

func TerraformTypesType(in []byte) (tftypes.Type, error) {
	var raw interface{}
	err := json.Unmarshal(in, &raw)
	if err != nil {
		return nil, fmt.Errorf("error parsing type, not valid JSON: %w", err)
	}
	return tftypes.ParseType(raw)
}
