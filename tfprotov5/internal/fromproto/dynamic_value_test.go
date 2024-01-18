package fromproto_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// nolint:unparam // typ may require other values in the future
func testTfplugin5DynamicValue(t *testing.T, typ tftypes.Type, value tftypes.Value) *tfplugin5.DynamicValue {
	t.Helper()

	return toproto.DynamicValue(testTfprotov5DynamicValue(t, typ, value))
}

func testTfprotov5DynamicValue(t *testing.T, typ tftypes.Type, value tftypes.Value) *tfprotov5.DynamicValue {
	t.Helper()

	dynamicValue, err := tfprotov5.NewDynamicValue(typ, value)

	if err != nil {
		t.Fatalf("unable to create DynamicValue: %s", err)
	}

	return &dynamicValue
}
