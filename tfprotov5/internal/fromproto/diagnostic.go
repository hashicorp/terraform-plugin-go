package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func Diagnostic(in tfplugin5.Diagnostic) tfprotov5.Diagnostic {
	diag := tfprotov5.Diagnostic{
		Severity: DiagnosticSeverity(in.Severity),
		Summary:  in.Summary,
		Detail:   in.Detail,
	}
	if in.Attribute != nil {
		attr := AttributePath(*in.Attribute)
		diag.Attribute = &attr
	}
	return diag
}

func DiagnosticSeverity(in tfplugin5.Diagnostic_Severity) tfprotov5.DiagnosticSeverity {
	return tfprotov5.DiagnosticSeverity(in)
}

func Diagnostics(in []*tfplugin5.Diagnostic) []*tfprotov5.Diagnostic {
	diagnostics := make([]*tfprotov5.Diagnostic, 0, len(in))
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
