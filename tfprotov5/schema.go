package tfprotov5

import "github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"

const (
	// SchemaNestedBlockNestingModeInvalid indicates that the nesting mode
	// for a nested block in the schema is invalid. This generally
	// indicates a nested block that was creating incorrectly.
	SchemaNestedBlockNestingModeInvalid SchemaNestedBlockNestingMode = 0

	// SchemaNestedBlockNestingModeSingle indicates that the nested block
	// should be treated as a single block, and there should not be more
	// than one of these blocks in the resource.
	SchemaNestedBlockNestingModeSingle SchemaNestedBlockNestingMode = 1

	// SchemaNestedBlockNestingModeList indicates that the nested block
	// should allow multiple blocks of that type in the config, and their
	// order should be considered significant, and identical values should
	// be permitted.
	SchemaNestedBlockNestingModeList SchemaNestedBlockNestingMode = 2

	// SchemaNestedBlockNestingModeSet indicates that the nested block
	// should allow multiple blocks of that type in the config, and their
	// order should not be significant, and identical values should be
	// considered only once.
	SchemaNestedBlockNestingModeSet SchemaNestedBlockNestingMode = 3

	// SchemaNestedBlockNestingModeMap indicates that the nested block should
	// TODO: what does this actually mean
	SchemaNestedBlockNestingModeMap SchemaNestedBlockNestingMode = 4

	// SchemaNestedBlockNestingModeGroup indicates that the nested block should
	// TODO: what does this actually mean
	SchemaNestedBlockNestingModeGroup SchemaNestedBlockNestingMode = 5
)

// Schema is how Terraform defines the shape of data. It can be thought of as
// the type information for resources, data sources, provider configuration,
// and all the other data that Terraform sends to providers. It is how
// providers express their requirements for that data.
type Schema struct {
	// Version indicates which version of the schema this is. Versions
	// should be monotonically incrementing numbers. When Terraform
	// encounters a resource stored in state with a schema version lower
	// that the schema version the provider advertises for that resource,
	// Terraform requests the provider upgrade the resource's state.
	Version int64

	// Block is the root level of the schema.
	// TODO: we can probably do a better explanation of the semantics of
	// this. If only I knew them. "It's that way because it is." isn't
	// ideal documentation, tho.
	Block *SchemaBlock
}

// SchemaBlock represents a block in a schema. Blocks are how Terraform creates
// groupings of attributes. In configurations, they don't use the equals sign
// and use dynamic instead of list comprehensions.
type SchemaBlock struct {
	// TODO: why do we have version in the block, too?
	Version int64

	// Attributes are the attributes defined within the block. These are
	// the fields that users can set using the equals sign or reference in
	// interpolations.
	Attributes []*SchemaAttribute

	// BlockTypes are the nested blocks within the block. These are used to
	// have blocks within blocks.
	BlockTypes []*SchemaNestedBlock

	// Description offers an end-user friendly description of what the
	// block is for. This will be surfaced to users through editor
	// integrations, documentation generation, and other settings.
	Description string

	// DescriptionKind indicates the formatting and encoding that the
	// Description field is using.
	DescriptionKind StringKind

	// Deprecated, when set to true, will indicate to the user that this
	// block is deprecated and the user should migrate away from it.
	//
	// TODO: does it? I feel like there should be a place for a
	// provider-specified message here... unless this is used for something
	// else.
	Deprecated bool
}

// SchemaAttribute represents a single attribute within a schema block.
// Attributes are the fields users can set in configuration using the equals
// sign, can assign to variables, can interpolate, and can use list
// comprehensions on.
type SchemaAttribute struct {
	// Name is the name of the attribute. This is what the user will put
	// before the equals sign to assign a value to this attribute.
	Name string

	// Type indicates the type of data the attribute expects. See the
	// documentation for the tftypes package for information on what types
	// are supported and their behaviors.
	Type tftypes.Type

	// Description offers an end-user friendly description of what the
	// attribute is for. This will be surfaced to users through editor
	// integrations, documentation generation, and other settings.
	Description string

	// Required, when set to true, indicates that this attribute must have
	// a value assigned to it by the user or Terraform will throw an error.
	Required bool

	// Optional, when set to true, indicates that the user does not need to
	// supply a value for this attribute, but may.
	Optional bool

	// Computed, when set to true, indicates the the provider will supply a
	// value for this field. If Optional and Required are false and
	// Computed is true, the user will not be able to specify a value for
	// this field without Terraform throwing an error. If Optional is true
	// and Computed is true, the user can specify a value for this field,
	// but the provider may supply a value if the user does not. It is
	// always a violation of Terraform's protocol to substitute a value for
	// what the user entered, even if Computed is true.
	Computed bool

	// Sensitive, when set to true, indicates that the contents of this
	// attribute should be considered sensitive and not included in output.
	// This does not encrypt or otherwise protect these values in state, it
	// only offers protection from them showing up in plans or other
	// output.
	Sensitive bool

	// DescriptionKind indicates the formatting and encoding that the
	// Description field is using.
	DescriptionKind StringKind

	// Deprecated, when set to true, will indicate to the user that this
	// attribute is deprecated and the user should migrate away from it.
	//
	// TODO: does it? I feel like there should be a place for a
	// provider-specified message here... unless this is used for something
	// else.
	Deprecated bool
}

// SchemaNestedBlock is a nested block within another block. See SchemaBlock
// for more information on blocks.
type SchemaNestedBlock struct {
	// TypeName is the name of the block. It is what the user will specify
	// when using the block in configuration.
	TypeName string

	// Block is the block being nested inside another block. See the
	// SchemaBlock documentation for more information on blocks.
	Block *SchemaBlock

	// Nesting is the kind of nesting the block is using. Different nesting
	// modes have different behaviors and imply different kinds of data.
	Nesting SchemaNestedBlockNestingMode

	// MinItems is the minimum number of instances of this block that a
	// user must specify or Terraform will return an error.
	MinItems int64

	// MaxItems is the maximum number of instances of this block that a
	// user may specify before Terraform returns an error.
	MaxItems int64
}

// SchemaNestedBlockNestingMode indicates the behavior a nested block should
// have and the type of data it implies. Blocks may have the same syntax and be
// the same construct, but they can be configured to imply different types of
// data and behave differently. See the documentation for each constant for the
// type of data it implies and the behavior it exhibits.
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
