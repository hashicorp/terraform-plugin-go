package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func Diagnostic(in *tfprotov5.Diagnostic) (*tfplugin5.Diagnostic, error) {
	diag := &tfplugin5.Diagnostic{
		Severity: Diagnostic_Severity(in.Severity),
		Summary:  in.Summary,
		Detail:   in.Detail,
	}
	if in.Attribute != nil {
		attr, err := AttributePath(in.Attribute)
		if err != nil {
			return diag, err
		}
		diag.Attribute = attr
	}
	return diag, nil
}

func Diagnostic_Severity(in tfprotov5.DiagnosticSeverity) tfplugin5.Diagnostic_Severity {
	return tfplugin5.Diagnostic_Severity(in)
}

func Diagnostics(in []*tfprotov5.Diagnostic) ([]*tfplugin5.Diagnostic, error) {
	diagnostics := make([]*tfplugin5.Diagnostic, 0, len(in))
	for _, diag := range in {
		if diag == nil {
			diagnostics = append(diagnostics, nil)
			continue
		}
		d, err := Diagnostic(diag)
		if err != nil {
			return diagnostics, err
		}
		diagnostics = append(diagnostics, d)
	}
	return diagnostics, nil
}
