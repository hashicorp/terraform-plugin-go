package metadata

// A JSON document that represents the static metadata of a Terraform provider.
type ProviderMetadata struct {
	// The schemas for data sources in this provider.
	DataSourceSchemas map[string]SchemaBlock `json:"data_source_schemas,omitempty"`
	// The definitions for functions in this provider.
	Functions map[string]Function `json:"functions,omitempty"`
	// The schema for the provider's configuration.
	Provider *SchemaBlock `json:"provider,omitempty"`
	// Metadata for the provider server.
	ProviderServer *ProviderServer `json:"provider_server,omitempty"`
	// The schemas for managed resources in this provider.
	ResourceSchemas map[string]SchemaBlock `json:"resource_schemas,omitempty"`
	// The Provider Metadata Specification version, based on a shortened form (MAJOR.MINOR) of
	// semantic versioning. Major versions denote a potential backwards incompatibility for
	// consumers. Minor versions should be considered additive.
	Version string `json:"version"`
}

// The representation of a particular schema (provider config, managed resource, data
// source)
//
// The schema for the provider's configuration.
type SchemaBlock struct {
	// The root-level block of configuration values.
	Block *Block `json:"block,omitempty"`
	// If true, this resource can be imported. This property is only valid for managed resource
	// schemas.
	SupportsImportState *bool `json:"supports_import_state,omitempty"`
	// If true, this resource can be moved across resource types. This property is only valid
	// for managed resource schemas.
	SupportsMoveState *bool `json:"supports_move_state,omitempty"`
	// A list of validations that are applied to the entire configuration at the schema-level.
	Validations []SchemaValidation `json:"validations,omitempty"`
	// The version of the schema.
	Version int64 `json:"version"`
}

// The top-level metadata for a nested block within a schema.
type BlockType struct {
	// The block data for this block type, including attributes and subsequent nested blocks.
	Block *Block `json:"block,omitempty"`
	// The upper limit on items that can be declared of this block type
	MaxItems *int64 `json:"max_items,omitempty"`
	// The lower limit on items that can be declared of this block type
	MinItems *int64 `json:"min_items,omitempty"`
	// The nesting mode for this block.
	NestingMode *BlockTypeNestingMode `json:"nesting_mode,omitempty"`
}

// The root-level block of configuration values.
//
// A nested block within a particular schema.
//
// The block data for this block type, including attributes and subsequent nested blocks.
type Block struct {
	// The attributes defined within this block.
	Attributes map[string]Attribute `json:"attributes,omitempty"`
	// The nested blocks defined within this block.
	BlockTypes map[string]BlockType `json:"block_types,omitempty"`
	// If true, this block is deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`
	// A message that indicates what actions should be performed by the practitioner to handle
	// the block deprecation.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`
	// The description for this block
	Description *string `json:"description,omitempty"`
	// The format of the description. If no kind is provided, it can be assumed to be plain text.
	DescriptionKind *DescriptionKind `json:"description_kind,omitempty"`
	// A list of plan modification operations that are applied to the block. This property is
	// only valid for blocks in managed resource schemas.
	PlanModifications []BlockPlanModification `json:"plan_modifications,omitempty"`
	// A list of validations that are applied to the block.
	Validations []BlockValidation `json:"validations,omitempty"`
}

// Details about a nested attribute type. Either type or nested_type is set, never both.
//
// The top-level metadata for a nested attribute within a schema.
type NestedAttributeType struct {
	// A map of nested attributes.
	Attributes map[string]Attribute `json:"attributes,omitempty"`
	// The upper limit on number of items that can be declared of this attribute type (not
	// applicable to single nesting mode).
	MaxItems *int64 `json:"max_items,omitempty"`
	// The lower limit on the number of items that can be declared of this attribute type (not
	// applicable to single nesting mode).
	MinItems *int64 `json:"min_items,omitempty"`
	// The nesting mode for this attribute.
	NestingMode *NestedTypeNestingMode `json:"nesting_mode,omitempty"`
}

// An attribute in a schema block, nested attribute, or nested block.
type Attribute struct {
	// If true, this attribute is computed - it can be set by the provider. It may also be set
	// by configuration if optional is true.
	Computed *bool `json:"computed,omitempty"`
	// If true, this attribute is deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`
	// A message that indicates what actions should be performed by the practitioner to handle
	// the attribute deprecation.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`
	// The description for this attribute.
	Description *string `json:"description,omitempty"`
	// The format of the description. If no kind is provided, it can be assumed to be plain text.
	DescriptionKind *DescriptionKind `json:"description_kind,omitempty"`
	// Details about a nested attribute type. Either type or nested_type is set, never both.
	NestedType *NestedAttributeType `json:"nested_type,omitempty"`
	// If true, this attribute is optional - it does not need to be entered in configuration.
	Optional *bool `json:"optional,omitempty"`
	// A list of plan modification operations that are applied to the attribute. This property
	// is only valid for attributes in managed resource schemas.
	PlanModifications []AttributePlanModification `json:"plan_modifications,omitempty"`
	// If true, this attribute is required - it has to be entered in configuration.
	Required *bool `json:"required,omitempty"`
	// Optional information that provides detail on how the provider will interpret the value of
	// this attribute.
	SDKType *SDKType `json:"sdk_type,omitempty"`
	// If true, this attribute is sensitive and will not be displayed in logs.
	Sensitive *bool `json:"sensitive,omitempty"`
	// The attribute type. Either type or nested_type is set, never both.
	Type *Type `json:"type"`
	// A list of validations that are applied to the attribute.
	Validations []AttributeValidation `json:"validations,omitempty"`
}

// A plan modification operation that is applied to an attribute.
type AttributePlanModification struct {
	// The description for this plan modification operation.
	Description *string `json:"description,omitempty"`
	// The format of the description. If no kind is provided, it can be assumed to be plain text.
	DescriptionKind *DescriptionKind `json:"description_kind,omitempty"`
}

// A validation that is applied to an attribute.
type AttributeValidation struct {
	// The description for this validation.
	Description *string `json:"description,omitempty"`
	// The format of the description. If no kind is provided, it can be assumed to be plain text.
	DescriptionKind *DescriptionKind `json:"description_kind,omitempty"`
}

// A plan modification operation that is applied to a block.
type BlockPlanModification struct {
	// The description for this plan modification operation.
	Description *string `json:"description,omitempty"`
	// The format of the description. If no kind is provided, it can be assumed to be plain text.
	DescriptionKind *DescriptionKind `json:"description_kind,omitempty"`
}

// A validation that is applied to a block.
type BlockValidation struct {
	// The description for this validation.
	Description *string `json:"description,omitempty"`
	// The format of the description. If no kind is provided, it can be assumed to be plain text.
	DescriptionKind *DescriptionKind `json:"description_kind,omitempty"`
}

// A validation that is applied to the entire configuration for a schema.
type SchemaValidation struct {
	// The description for this validation.
	Description *string `json:"description,omitempty"`
	// The format of the description. If no kind is provided, it can be assumed to be plain text.
	DescriptionKind *DescriptionKind `json:"description_kind,omitempty"`
}

// A function signature for a provider-defined function.
type Function struct {
	// A message that indicates that the function should be considered deprecated and what
	// actions should be performed by the practitioner to handle the deprecation.
	DeprecationMessage *string `json:"deprecation_message,omitempty"`
	// The human-readable description.
	Description *string `json:"description,omitempty"`
	// The function's fixed positional parameters.
	Parameters []FunctionParameter `json:"parameters,omitempty"`
	// Optional information that provides detail on the return value of this function.
	ReturnSDKType *SDKType `json:"return_sdk_type,omitempty"`
	// The function's return type.
	ReturnType *Type `json:"return_type"`
	// A shortened description of the function.
	Summary *string `json:"summary,omitempty"`
	// The function's variadic parameter if it is supported.
	VariadicParameter *FunctionParameter `json:"variadic_parameter,omitempty"`
}

// A parameter to a function.
//
// The function's variadic parameter if it is supported.
type FunctionParameter struct {
	// The human-readable description of the argument.
	Description *string `json:"description,omitempty"`
	// If true, null will be an acceptable value for the argument.
	IsNullable *bool `json:"is_nullable,omitempty"`
	// A name for the argument.
	Name *string `json:"name,omitempty"`
	// Optional information that provides detail on how the provider will interpret the value of
	// this parameter.
	SDKType *SDKType `json:"sdk_type,omitempty"`
	// A type that an argument for this parameter must conform to.
	Type *Type `json:"type"`
}

// Metadata for the provider server.
type ProviderServer struct {
	// Address is the full address of the provider. Full address form has three parts separated
	// by forward slashes (/): Hostname, namespace, and provider type ('name').
	Address *string `json:"address,omitempty"`
	// The protocol version that should be used when serving the provider. Either protocol
	// version 5 or protocol version 6 can be used. Protocol version 6 introduces support for
	// nested attributes.
	ProtocolVersion *int64 `json:"protocol_version,omitempty"`
}

// The nesting mode for this block.
//
// The nesting mode for a particular nested schema block.
type BlockTypeNestingMode string

const (
	FluffyList   BlockTypeNestingMode = "list"
	FluffyMap    BlockTypeNestingMode = "map"
	FluffySet    BlockTypeNestingMode = "set"
	FluffySingle BlockTypeNestingMode = "single"
	Group        BlockTypeNestingMode = "group"
)

// The nesting mode for this attribute.
//
// The nesting mode for a particular nested schema attribute.
type NestedTypeNestingMode string

const (
	PurpleList   NestedTypeNestingMode = "list"
	PurpleMap    NestedTypeNestingMode = "map"
	PurpleSet    NestedTypeNestingMode = "set"
	PurpleSingle NestedTypeNestingMode = "single"
)

// The format of the description. If no kind is provided, it can be assumed to be plain
// text.
//
// Describes the format type for a particular description's field. If no kind is provided,
// it can be assumed to be plain text.
type DescriptionKind string

const (
	Markdown DescriptionKind = "markdown"
	Plain    DescriptionKind = "plain"
)

// Optional information that provides detail on how the provider will interpret the value of
// this attribute.
//
// Additional type information that provides detail on how the provider will interpret the
// value of this attribute.
//
// Optional information that provides detail on how the provider will interpret the value of
// this parameter.
//
// Optional information that provides detail on the return value of this function.
type SDKType string

const (
	Float32 SDKType = "float32"
	Float64 SDKType = "float64"
	Int32   SDKType = "int32"
	Int64   SDKType = "int64"
)

type ReturnTypeEnum string

// ** TODO: add extra types here later
const (
	Bool    ReturnTypeEnum = "bool"
	Dynamic ReturnTypeEnum = "dynamic"
	Number  ReturnTypeEnum = "number"
	String  ReturnTypeEnum = "string"
)

// The attribute type. Either type or nested_type is set, never both.
//
// A type that is used by Terraform core for early validation and conversions.
//
// A type that an argument for this parameter must conform to.
//
// The function's return type.
type Type struct {
	Enum       *ReturnTypeEnum
	UnionArray []Te
}

// Complex types like maps, sets, lists, objects.
type Te struct {
	AnythingArray []interface{}
	AnythingMap   map[string]interface{}
	String        *string
}
