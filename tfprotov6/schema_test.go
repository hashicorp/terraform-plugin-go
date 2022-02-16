package tfprotov6_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestSchemaAttributeValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schemaAttribute *tfprotov6.SchemaAttribute
		expected        tftypes.Type
	}{
		"nil": {
			schemaAttribute: nil,
			expected:        nil,
		},
		"Bool": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.Bool,
			},
			expected: tftypes.Bool,
		},
		"DynamicPseudoType": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.DynamicPseudoType,
			},
			expected: tftypes.DynamicPseudoType,
		},
		"List-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.List{
					ElementType: tftypes.String,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.String,
			},
		},
		"Map-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.Map{
					ElementType: tftypes.String,
				},
			},
			expected: tftypes.Map{
				ElementType: tftypes.String,
			},
		},
		"NestedList-Bool": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "bool",
							Type: tftypes.Bool,
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"bool": tftypes.Bool,
					},
				},
			},
		},
		"NestedList-DynamicPseudoType": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "dynamic",
							Type: tftypes.DynamicPseudoType,
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"dynamic": tftypes.DynamicPseudoType,
					},
				},
			},
		},
		"NestedList-List-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "list",
							Type: tftypes.List{
								ElementType: tftypes.String,
							},
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"list": tftypes.List{
							ElementType: tftypes.String,
						},
					},
				},
			},
		},
		"NestedList-Map-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "map",
							Type: tftypes.Map{
								ElementType: tftypes.String,
							},
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"map": tftypes.Map{
							ElementType: tftypes.String,
						},
					},
				},
			},
		},
		"NestedList-Number": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "number",
							Type: tftypes.Number,
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"number": tftypes.Number,
					},
				},
			},
		},
		"NestedList-Object-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "object",
							Type: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"string": tftypes.String,
								},
							},
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"object": tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"string": tftypes.String,
							},
						},
					},
				},
			},
		},
		"NestedList-Set-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "set",
							Type: tftypes.Set{
								ElementType: tftypes.String,
							},
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"set": tftypes.Set{
							ElementType: tftypes.String,
						},
					},
				},
			},
		},
		"NestedList-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "string",
							Type: tftypes.String,
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"string": tftypes.String,
					},
				},
			},
		},
		"NestedList-Tuple-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "tuple",
							Type: tftypes.Tuple{
								ElementTypes: []tftypes.Type{tftypes.String},
							},
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"tuple": tftypes.Tuple{
							ElementTypes: []tftypes.Type{tftypes.String},
						},
					},
				},
			},
		},
		"NestedMap-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "string",
							Type: tftypes.String,
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeMap,
				},
			},
			expected: tftypes.Map{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"string": tftypes.String,
					},
				},
			},
		},
		"NestedSet-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "string",
							Type: tftypes.String,
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeSet,
				},
			},
			expected: tftypes.Set{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"string": tftypes.String,
					},
				},
			},
		},
		"NestedSingle-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "string",
							Type: tftypes.String,
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeSingle,
				},
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"string": tftypes.String,
				},
			},
		},
		"Number": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.Number,
			},
			expected: tftypes.Number,
		},
		"Object-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"string": tftypes.String,
					},
				},
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"string": tftypes.String,
				},
			},
		},
		"Set-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.Set{
					ElementType: tftypes.String,
				},
			},
			expected: tftypes.Set{
				ElementType: tftypes.String,
			},
		},
		"String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.String,
			},
			expected: tftypes.String,
		},
		"Tuple-String": {
			schemaAttribute: &tfprotov6.SchemaAttribute{
				Type: tftypes.Tuple{
					ElementTypes: []tftypes.Type{tftypes.String},
				},
			},
			expected: tftypes.Tuple{
				ElementTypes: []tftypes.Type{tftypes.String},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.schemaAttribute.ValueType()

			if testCase.expected == nil {
				if got == nil {
					return
				}

				t.Fatalf("expected nil, got: %s", got)
			}

			if !testCase.expected.Equal(got) {
				t.Errorf("expected %s, got: %s", testCase.expected, got)
			}
		})
	}
}

func TestSchemaBlockValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schemaBlock *tfprotov6.SchemaBlock
		expected    tftypes.Type
	}{
		"nil": {
			schemaBlock: nil,
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"Attribute-String": {
			schemaBlock: &tfprotov6.SchemaBlock{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "test_string_attribute",
						Type: tftypes.String,
					},
				},
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_string_attribute": tftypes.String,
				},
			},
		},
		"Block-NestingList-String": {
			schemaBlock: &tfprotov6.SchemaBlock{
				BlockTypes: []*tfprotov6.SchemaNestedBlock{
					{
						Block: &tfprotov6.SchemaBlock{
							Attributes: []*tfprotov6.SchemaAttribute{
								{
									Name: "test_string_attribute",
									Type: tftypes.String,
								},
							},
						},
						Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
						TypeName: "test_list_block",
					},
				},
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_list_block": tftypes.List{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"test_string_attribute": tftypes.String,
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.schemaBlock.ValueType()

			if testCase.expected == nil {
				if got == nil {
					return
				}

				t.Fatalf("expected nil, got: %s", got)
			}

			if !testCase.expected.Equal(got) {
				t.Errorf("expected %s, got: %s", testCase.expected, got)
			}
		})
	}
}

func TestSchemaNestedBlockValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schemaNestedBlock *tfprotov6.SchemaNestedBlock
		expected          tftypes.Type
	}{
		"nil": {
			schemaNestedBlock: nil,
			expected:          nil,
		},
		"NestingGroup-Attribute-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeGroup,
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_string_attribute": tftypes.String,
				},
			},
		},
		"NestingGroup-Block-NestingList-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					BlockTypes: []*tfprotov6.SchemaNestedBlock{
						{
							Block: &tfprotov6.SchemaBlock{
								Attributes: []*tfprotov6.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeGroup,
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_list_block": tftypes.List{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"test_string_attribute": tftypes.String,
							},
						},
					},
				},
			},
		},
		"NestingInvalid": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Nesting: tfprotov6.SchemaNestedBlockNestingModeInvalid,
			},
			expected: nil,
		},
		"NestingList-missing-Block": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Nesting: tfprotov6.SchemaNestedBlockNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{},
				},
			},
		},
		"NestingList-Attribute-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"test_string_attribute": tftypes.String,
					},
				},
			},
		},
		"NestingList-Block-NestingList-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					BlockTypes: []*tfprotov6.SchemaNestedBlock{
						{
							Block: &tfprotov6.SchemaBlock{
								Attributes: []*tfprotov6.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"test_list_block": tftypes.List{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"test_string_attribute": tftypes.String,
								},
							},
						},
					},
				},
			},
		},
		"NestingMap-Attribute-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeMap,
			},
			expected: tftypes.Map{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"test_string_attribute": tftypes.String,
					},
				},
			},
		},
		"NestingMap-Block-NestingList-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					BlockTypes: []*tfprotov6.SchemaNestedBlock{
						{
							Block: &tfprotov6.SchemaBlock{
								Attributes: []*tfprotov6.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeMap,
			},
			expected: tftypes.Map{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"test_list_block": tftypes.List{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"test_string_attribute": tftypes.String,
								},
							},
						},
					},
				},
			},
		},
		"NestingSet-Attribute-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeSet,
			},
			expected: tftypes.Set{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"test_string_attribute": tftypes.String,
					},
				},
			},
		},
		"NestingSet-Block-NestingList-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					BlockTypes: []*tfprotov6.SchemaNestedBlock{
						{
							Block: &tfprotov6.SchemaBlock{
								Attributes: []*tfprotov6.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeSet,
			},
			expected: tftypes.Set{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"test_list_block": tftypes.List{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"test_string_attribute": tftypes.String,
								},
							},
						},
					},
				},
			},
		},
		"NestingSingle-Attribute-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeSingle,
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_string_attribute": tftypes.String,
				},
			},
		},
		"NestingSingle-Block-NestingList-String": {
			schemaNestedBlock: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					BlockTypes: []*tfprotov6.SchemaNestedBlock{
						{
							Block: &tfprotov6.SchemaBlock{
								Attributes: []*tfprotov6.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov6.SchemaNestedBlockNestingModeSingle,
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_list_block": tftypes.List{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"test_string_attribute": tftypes.String,
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.schemaNestedBlock.ValueType()

			if testCase.expected == nil {
				if got == nil {
					return
				}

				t.Fatalf("expected nil, got: %s", got)
			}

			if !testCase.expected.Equal(got) {
				t.Errorf("expected %s, got: %s", testCase.expected, got)
			}
		})
	}
}

func TestSchemaObjectValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schemaObject *tfprotov6.SchemaObject
		expected     tftypes.Type
	}{
		"nil": {
			schemaObject: nil,
			expected:     nil,
		},
		"NestedList-Bool": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "bool",
						Type: tftypes.Bool,
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"bool": tftypes.Bool,
					},
				},
			},
		},
		"NestedList-DynamicPseudoType": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "dynamic",
						Type: tftypes.DynamicPseudoType,
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"dynamic": tftypes.DynamicPseudoType,
					},
				},
			},
		},
		"NestedList-List-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "list",
						Type: tftypes.List{
							ElementType: tftypes.String,
						},
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"list": tftypes.List{
							ElementType: tftypes.String,
						},
					},
				},
			},
		},
		"NestedList-Map-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "map",
						Type: tftypes.Map{
							ElementType: tftypes.String,
						},
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"map": tftypes.Map{
							ElementType: tftypes.String,
						},
					},
				},
			},
		},
		"NestedList-Number": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "number",
						Type: tftypes.Number,
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"number": tftypes.Number,
					},
				},
			},
		},
		"NestedList-Object-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "object",
						Type: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"string": tftypes.String,
							},
						},
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"object": tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"string": tftypes.String,
							},
						},
					},
				},
			},
		},
		"NestedList-Set-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "set",
						Type: tftypes.Set{
							ElementType: tftypes.String,
						},
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"set": tftypes.Set{
							ElementType: tftypes.String,
						},
					},
				},
			},
		},
		"NestedList-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "string",
						Type: tftypes.String,
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"string": tftypes.String,
					},
				},
			},
		},
		"NestedList-Tuple-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "tuple",
						Type: tftypes.Tuple{
							ElementTypes: []tftypes.Type{tftypes.String},
						},
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"tuple": tftypes.Tuple{
							ElementTypes: []tftypes.Type{tftypes.String},
						},
					},
				},
			},
		},
		"NestedMap-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "string",
						Type: tftypes.String,
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeMap,
			},
			expected: tftypes.Map{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"string": tftypes.String,
					},
				},
			},
		},
		"NestedSet-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "string",
						Type: tftypes.String,
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeSet,
			},
			expected: tftypes.Set{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"string": tftypes.String,
					},
				},
			},
		},
		"NestedSingle-String": {
			schemaObject: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "string",
						Type: tftypes.String,
					},
				},
				Nesting: tfprotov6.SchemaObjectNestingModeSingle,
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"string": tftypes.String,
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.schemaObject.ValueType()

			if testCase.expected == nil {
				if got == nil {
					return
				}

				t.Fatalf("expected nil, got: %s", got)
			}

			if !testCase.expected.Equal(got) {
				t.Errorf("expected %s, got: %s", testCase.expected, got)
			}
		})
	}
}

func TestSchemaValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schema   *tfprotov6.Schema
		expected tftypes.Type
	}{
		"nil": {
			schema: nil,
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"missing-Block": {
			schema: &tfprotov6.Schema{},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"Attribute-String": {
			schema: &tfprotov6.Schema{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_string_attribute": tftypes.String,
				},
			},
		},
		"Block-NestingList-String": {
			schema: &tfprotov6.Schema{
				Block: &tfprotov6.SchemaBlock{
					BlockTypes: []*tfprotov6.SchemaNestedBlock{
						{
							Block: &tfprotov6.SchemaBlock{
								Attributes: []*tfprotov6.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_list_block": tftypes.List{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"test_string_attribute": tftypes.String,
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.schema.ValueType()

			if testCase.expected == nil {
				if got == nil {
					return
				}

				t.Fatalf("expected nil, got: %s", got)
			}

			if !testCase.expected.Equal(got) {
				t.Errorf("expected %s, got: %s", testCase.expected, got)
			}
		})
	}
}
