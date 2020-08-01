package tfprotov5

const (
	StringKindPlain    StringKind = 0
	StringKindMarkdown StringKind = 1
)

type StringKind int32

func (s StringKind) String() string {
	switch s {
	case 0:
		return "PLAIN"
	case 1:
		return "MARKDOWN"
	}
	return "UNKNOWN"
}
