package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func Schema(in *tfprotov6.Schema) (*tfplugin6.Schema, error) {
	if in == nil {
		return nil, nil
	}

	block, err := Schema_Block(in.Block)

	if err != nil {
		return nil, fmt.Errorf("error marshalling block: %w", err)
	}

	resp := &tfplugin6.Schema{
		Block:   block,
		Version: in.Version,
	}

	return resp, nil
}

func Schema_Block(in *tfprotov6.SchemaBlock) (*tfplugin6.Schema_Block, error) {
	if in == nil {
		return nil, nil
	}

	attributes, err := Schema_Attributes(in.Attributes)

	if err != nil {
		return nil, err
	}

	blockTypes, err := Schema_NestedBlocks(in.BlockTypes)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.Schema_Block{
		Attributes:      attributes,
		BlockTypes:      blockTypes,
		Deprecated:      in.Deprecated,
		Description:     in.Description,
		DescriptionKind: StringKind(in.DescriptionKind),
		Version:         in.Version,
	}

	return resp, nil
}

func Schema_Attribute(in *tfprotov6.SchemaAttribute) (*tfplugin6.Schema_Attribute, error) {
	if in == nil {
		return nil, nil
	}

	typ, err := CtyType(in.Type)

	if err != nil {
		return nil, fmt.Errorf("error marshaling %q type to JSON: %w", in.Name, err)
	}

	nestedType, err := Schema_Object(in.NestedType)

	if err != nil {
		return nil, fmt.Errorf("error marshaling %q nested type to JSON: %w", in.Name, err)
	}

	resp := &tfplugin6.Schema_Attribute{
		Computed:        in.Computed,
		Deprecated:      in.Deprecated,
		Description:     in.Description,
		DescriptionKind: StringKind(in.DescriptionKind),
		Name:            in.Name,
		NestedType:      nestedType,
		Optional:        in.Optional,
		Required:        in.Required,
		Sensitive:       in.Sensitive,
		Type:            typ,
	}

	return resp, nil
}

func Schema_Attributes(in []*tfprotov6.SchemaAttribute) ([]*tfplugin6.Schema_Attribute, error) {
	resp := make([]*tfplugin6.Schema_Attribute, 0, len(in))

	for _, a := range in {
		attr, err := Schema_Attribute(a)

		if err != nil {
			return nil, err
		}

		resp = append(resp, attr)
	}

	return resp, nil
}

func Schema_NestedBlock(in *tfprotov6.SchemaNestedBlock) (*tfplugin6.Schema_NestedBlock, error) {
	if in == nil {
		return nil, nil
	}

	block, err := Schema_Block(in.Block)

	if err != nil {
		return nil, fmt.Errorf("error marshaling %q nested block: %w", in.TypeName, err)
	}

	resp := &tfplugin6.Schema_NestedBlock{
		Block:    block,
		MaxItems: in.MaxItems,
		MinItems: in.MinItems,
		Nesting:  Schema_NestedBlock_NestingMode(in.Nesting),
		TypeName: in.TypeName,
	}

	return resp, nil
}

func Schema_NestedBlocks(in []*tfprotov6.SchemaNestedBlock) ([]*tfplugin6.Schema_NestedBlock, error) {
	resp := make([]*tfplugin6.Schema_NestedBlock, 0, len(in))

	for _, b := range in {
		block, err := Schema_NestedBlock(b)

		if err != nil {
			return nil, err
		}

		resp = append(resp, block)
	}

	return resp, nil
}

func Schema_NestedBlock_NestingMode(in tfprotov6.SchemaNestedBlockNestingMode) tfplugin6.Schema_NestedBlock_NestingMode {
	return tfplugin6.Schema_NestedBlock_NestingMode(in)
}

func Schema_Object_NestingMode(in tfprotov6.SchemaObjectNestingMode) tfplugin6.Schema_Object_NestingMode {
	return tfplugin6.Schema_Object_NestingMode(in)
}

func Schema_Object(in *tfprotov6.SchemaObject) (*tfplugin6.Schema_Object, error) {
	if in == nil {
		return nil, nil
	}

	attributes, err := Schema_Attributes(in.Attributes)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.Schema_Object{
		Attributes: attributes,
		Nesting:    Schema_Object_NestingMode(in.Nesting),
	}

	return resp, nil
}
