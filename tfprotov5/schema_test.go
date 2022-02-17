package tfprotov5_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestSchemaAttributeValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schemaAttribute *tfprotov5.SchemaAttribute
		expected        tftypes.Type
	}{
		"nil": {
			schemaAttribute: nil,
			expected:        nil,
		},
		"Bool": {
			schemaAttribute: &tfprotov5.SchemaAttribute{
				Type: tftypes.Bool,
			},
			expected: tftypes.Bool,
		},
		"DynamicPseudoType": {
			schemaAttribute: &tfprotov5.SchemaAttribute{
				Type: tftypes.DynamicPseudoType,
			},
			expected: tftypes.DynamicPseudoType,
		},
		"List-String": {
			schemaAttribute: &tfprotov5.SchemaAttribute{
				Type: tftypes.List{
					ElementType: tftypes.String,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.String,
			},
		},
		"Map-String": {
			schemaAttribute: &tfprotov5.SchemaAttribute{
				Type: tftypes.Map{
					ElementType: tftypes.String,
				},
			},
			expected: tftypes.Map{
				ElementType: tftypes.String,
			},
		},
		"Number": {
			schemaAttribute: &tfprotov5.SchemaAttribute{
				Type: tftypes.Number,
			},
			expected: tftypes.Number,
		},
		"Object-String": {
			schemaAttribute: &tfprotov5.SchemaAttribute{
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
			schemaAttribute: &tfprotov5.SchemaAttribute{
				Type: tftypes.Set{
					ElementType: tftypes.String,
				},
			},
			expected: tftypes.Set{
				ElementType: tftypes.String,
			},
		},
		"String": {
			schemaAttribute: &tfprotov5.SchemaAttribute{
				Type: tftypes.String,
			},
			expected: tftypes.String,
		},
		"Tuple-String": {
			schemaAttribute: &tfprotov5.SchemaAttribute{
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
		schemaBlock *tfprotov5.SchemaBlock
		expected    tftypes.Type
	}{
		"nil": {
			schemaBlock: nil,
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"Attribute-String": {
			schemaBlock: &tfprotov5.SchemaBlock{
				Attributes: []*tfprotov5.SchemaAttribute{
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
			schemaBlock: &tfprotov5.SchemaBlock{
				BlockTypes: []*tfprotov5.SchemaNestedBlock{
					{
						Block: &tfprotov5.SchemaBlock{
							Attributes: []*tfprotov5.SchemaAttribute{
								{
									Name: "test_string_attribute",
									Type: tftypes.String,
								},
							},
						},
						Nesting:  tfprotov5.SchemaNestedBlockNestingModeList,
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
		schemaNestedBlock *tfprotov5.SchemaNestedBlock
		expected          tftypes.Type
	}{
		"nil": {
			schemaNestedBlock: nil,
			expected:          nil,
		},
		"NestingGroup-Attribute-String": {
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeGroup,
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_string_attribute": tftypes.String,
				},
			},
		},
		"NestingGroup-Block-NestingList-String": {
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					BlockTypes: []*tfprotov5.SchemaNestedBlock{
						{
							Block: &tfprotov5.SchemaBlock{
								Attributes: []*tfprotov5.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov5.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeGroup,
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
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Nesting: tfprotov5.SchemaNestedBlockNestingModeInvalid,
			},
			expected: nil,
		},
		"NestingList-missing-Block": {
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Nesting: tfprotov5.SchemaNestedBlockNestingModeList,
			},
			expected: tftypes.List{
				ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{},
				},
			},
		},
		"NestingList-Attribute-String": {
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeList,
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
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					BlockTypes: []*tfprotov5.SchemaNestedBlock{
						{
							Block: &tfprotov5.SchemaBlock{
								Attributes: []*tfprotov5.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov5.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeList,
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
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeMap,
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
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					BlockTypes: []*tfprotov5.SchemaNestedBlock{
						{
							Block: &tfprotov5.SchemaBlock{
								Attributes: []*tfprotov5.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov5.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeMap,
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
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeSet,
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
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					BlockTypes: []*tfprotov5.SchemaNestedBlock{
						{
							Block: &tfprotov5.SchemaBlock{
								Attributes: []*tfprotov5.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov5.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeSet,
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
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name: "test_string_attribute",
							Type: tftypes.String,
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeSingle,
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_string_attribute": tftypes.String,
				},
			},
		},
		"NestingSingle-Block-NestingList-String": {
			schemaNestedBlock: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					BlockTypes: []*tfprotov5.SchemaNestedBlock{
						{
							Block: &tfprotov5.SchemaBlock{
								Attributes: []*tfprotov5.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov5.SchemaNestedBlockNestingModeList,
							TypeName: "test_list_block",
						},
					},
				},
				Nesting: tfprotov5.SchemaNestedBlockNestingModeSingle,
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

func TestSchemaValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schema   *tfprotov5.Schema
		expected tftypes.Type
	}{
		"nil": {
			schema: nil,
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"missing-Block": {
			schema: &tfprotov5.Schema{},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"Attribute-String": {
			schema: &tfprotov5.Schema{
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
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
			schema: &tfprotov5.Schema{
				Block: &tfprotov5.SchemaBlock{
					BlockTypes: []*tfprotov5.SchemaNestedBlock{
						{
							Block: &tfprotov5.SchemaBlock{
								Attributes: []*tfprotov5.SchemaAttribute{
									{
										Name: "test_string_attribute",
										Type: tftypes.String,
									},
								},
							},
							Nesting:  tfprotov5.SchemaNestedBlockNestingModeList,
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
