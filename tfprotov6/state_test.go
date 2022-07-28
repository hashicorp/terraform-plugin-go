package tfprotov6_test

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestRawStateUnmarshalWithOpts(t *testing.T) {
	t.Parallel()
	type testCase struct {
		rawState tfprotov6.RawState
		value    tftypes.Value
		typ      tftypes.Type
		opts     tfprotov6.UnmarshalOpts
	}
	tests := map[string]testCase{
		"object-of-bool-number": {
			rawState: tfprotov6.RawState{
				JSON: []byte(`{"bool":true,"number":0}`),
			},
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"bool":   tftypes.Bool,
					"number": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"bool":   tftypes.NewValue(tftypes.Bool, true),
				"number": tftypes.NewValue(tftypes.Number, big.NewFloat(0)),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"bool":   tftypes.Bool,
					"number": tftypes.Number,
				},
			},
		},
		"object-with-missing-attribute": {
			rawState: tfprotov6.RawState{
				JSON: []byte(`{"bool":true,"number":0,"unknown":"whatever"}`),
			},
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"bool":   tftypes.Bool,
					"number": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"bool":   tftypes.NewValue(tftypes.Bool, true),
				"number": tftypes.NewValue(tftypes.Number, big.NewFloat(0)),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"bool":   tftypes.Bool,
					"number": tftypes.Number,
				},
			},
			opts: tfprotov6.UnmarshalOpts{
				ValueFromJSONOpts: tftypes.ValueFromJSONOpts{
					IgnoreUndefinedAttributes: true,
				},
			},
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			val, err := test.rawState.UnmarshalWithOpts(test.typ, test.opts)
			if err != nil {
				t.Fatalf("unexpected error unmarshaling: %s", err)
			}

			if diff := cmp.Diff(test.value, val); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}
