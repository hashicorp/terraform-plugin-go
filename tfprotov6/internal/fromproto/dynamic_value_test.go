package fromproto_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// nolint:unparam // typ may require other values in the future
func testTfplugin6DynamicValue(t *testing.T, typ tftypes.Type, value tftypes.Value) *tfplugin6.DynamicValue {
	t.Helper()

	return toproto.DynamicValue(testTfprotov6DynamicValue(t, typ, value))
}

func testTfprotov6DynamicValue(t *testing.T, typ tftypes.Type, value tftypes.Value) *tfprotov6.DynamicValue {
	t.Helper()

	dynamicValue, err := tfprotov6.NewDynamicValue(typ, value)

	if err != nil {
		t.Fatalf("unable to create DynamicValue: %s", err)
	}

	return &dynamicValue
}
