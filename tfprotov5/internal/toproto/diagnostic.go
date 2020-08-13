package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func Diagnostic(in tfprotov5.Diagnostic) tfplugin5.Diagnostic {
	diag := tfplugin5.Diagnostic{
		Severity: Diagnostic_Severity(in.Severity),
		Summary:  in.Summary,
		Detail:   in.Detail,
	}
	if in.Attribute != nil {
		attr := AttributePath(*in.Attribute)
		diag.Attribute = &attr
	}
	return diag
}

func Diagnostic_Severity(in tfprotov5.DiagnosticSeverity) tfplugin5.Diagnostic_Severity {
	return tfplugin5.Diagnostic_Severity(in)
}

func Diagnostics(in []*tfprotov5.Diagnostic) []*tfplugin5.Diagnostic {
	diagnostics := make([]*tfplugin5.Diagnostic, 0, len(in))
	for _, diag := range in {
		if diag == nil {
			diagnostics = append(diagnostics, nil)
			continue
		}
		d := Diagnostic(*diag)
		diagnostics = append(diagnostics, &d)
	}
	return diagnostics
}
