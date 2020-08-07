package tfprotov5

const (
	SchemaNestedBlockNestingModeInvalid SchemaNestedBlockNestingMode = 0
	SchemaNestedBlockNestingModeSingle  SchemaNestedBlockNestingMode = 1
	SchemaNestedBlockNestingModeList    SchemaNestedBlockNestingMode = 2
	SchemaNestedBlockNestingModeSet     SchemaNestedBlockNestingMode = 3
	SchemaNestedBlockNestingModeMap     SchemaNestedBlockNestingMode = 4
	SchemaNestedBlockNestingModeGroup   SchemaNestedBlockNestingMode = 5
)

type Schema struct {
	Version int64
	Block   *SchemaBlock
}

type SchemaBlock struct {
	Version         int64
	Attributes      []*SchemaAttribute
	BlockTypes      []*SchemaNestedBlock
	Description     string
	DescriptionKind StringKind
	Deprecated      bool
}

type SchemaAttribute struct {
	Name            string
	Type            []byte // TODO: this is the JSON encoding of a cty type
	Description     string
	Required        bool
	Optional        bool
	Computed        bool
	Sensitive       bool
	DescriptionKind StringKind
	Deprecated      bool
}

type SchemaNestedBlock struct {
	TypeName string
	Block    *SchemaBlock
	Nesting  SchemaNestedBlockNestingMode
	MinItems int64
	MaxItems int64
}

type SchemaNestedBlockNestingMode int32

func (s SchemaNestedBlockNestingMode) String() string {
	switch s {
	case 0:
		return "INVALID"
	case 1:
		return "SINGLE"
	case 2:
		return "LIST"
	case 3:
		return "SET"
	case 4:
		return "MAP"
	case 5:
		return "GROUP"
	}
	return "UNKNOWN"
}
