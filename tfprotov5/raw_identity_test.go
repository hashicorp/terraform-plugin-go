// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5_test

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestRawIdentityUnmarshalWithOpts(t *testing.T) {
	t.Parallel()
	type testCase struct {
		RawIdentity tfprotov5.RawIdentity
		value       tftypes.Value
		typ         tftypes.Type
		opts        tfprotov5.UnmarshalOpts
	}
	tests := map[string]testCase{
		"object-of-bool-number": {
			RawIdentity: tfprotov5.RawIdentity{
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
			RawIdentity: tfprotov5.RawIdentity{
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
			opts: tfprotov5.UnmarshalOpts{
				ValueFromJSONOpts: tftypes.ValueFromJSONOpts{
					IgnoreUndefinedAttributes: true,
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			val, err := test.RawIdentity.UnmarshalWithOpts(test.typ, test.opts)
			if err != nil {
				t.Fatalf("unexpected error unmarshaling: %s", err)
			}

			if diff := cmp.Diff(test.value, val); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}
