package tfprotov5

const (
	DiagnosticSeverityInvalid DiagnosticSeverity = 0
	DiagnosticSeverityError   DiagnosticSeverity = 1
	DiagnosticSeverityWarning DiagnosticSeverity = 2
)

type Diagnostic struct {
	Severity  DiagnosticSeverity
	Summary   string
	Detail    string
	Attribute *AttributePath
}

type DiagnosticSeverity int32

func (d DiagnosticSeverity) String() string {
	switch d {
	case 0:
		return "INVALID"
	case 1:
		return "ERROR"
	case 2:
		return "WARNING"
	}
	return "UNKNOWN"
}
