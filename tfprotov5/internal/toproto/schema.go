// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func Schema(in *tfprotov5.Schema) (*tfplugin5.Schema, error) {
	if in == nil {
		return nil, nil
	}

	block, err := Schema_Block(in.Block)

	if err != nil {
		return nil, fmt.Errorf("error marshalling block: %w", err)
	}

	resp := &tfplugin5.Schema{
		Block:   block,
		Version: in.Version,
	}

	return resp, nil
}

func Schema_Block(in *tfprotov5.SchemaBlock) (*tfplugin5.Schema_Block, error) {
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

	resp := &tfplugin5.Schema_Block{
		Attributes:      attributes,
		BlockTypes:      blockTypes,
		Deprecated:      in.Deprecated,
		Description:     in.Description,
		DescriptionKind: StringKind(in.DescriptionKind),
		Version:         in.Version,
	}

	return resp, nil
}

func Schema_Attribute(in *tfprotov5.SchemaAttribute) (*tfplugin5.Schema_Attribute, error) {
	if in == nil {
		return nil, nil
	}

	typ, err := CtyType(in.Type)

	if err != nil {
		return nil, fmt.Errorf("error marshaling %q type to JSON: %w", in.Name, err)
	}

	resp := &tfplugin5.Schema_Attribute{
		Computed:        in.Computed,
		Deprecated:      in.Deprecated,
		Description:     in.Description,
		DescriptionKind: StringKind(in.DescriptionKind),
		Name:            in.Name,
		Optional:        in.Optional,
		Required:        in.Required,
		Sensitive:       in.Sensitive,
		Type:            typ,
	}

	return resp, nil
}

func Schema_Attributes(in []*tfprotov5.SchemaAttribute) ([]*tfplugin5.Schema_Attribute, error) {
	resp := make([]*tfplugin5.Schema_Attribute, 0, len(in))

	for _, a := range in {
		attr, err := Schema_Attribute(a)

		if err != nil {
			return nil, err
		}

		resp = append(resp, attr)
	}

	return resp, nil
}

func Schema_NestedBlock(in *tfprotov5.SchemaNestedBlock) (*tfplugin5.Schema_NestedBlock, error) {
	if in == nil {
		return nil, nil
	}

	block, err := Schema_Block(in.Block)

	if err != nil {
		return nil, fmt.Errorf("error marshaling %q nested block: %w", in.TypeName, err)
	}

	resp := &tfplugin5.Schema_NestedBlock{
		Block:    block,
		MaxItems: in.MaxItems,
		MinItems: in.MinItems,
		Nesting:  Schema_NestedBlock_NestingMode(in.Nesting),
		TypeName: in.TypeName,
	}

	return resp, nil
}

func Schema_NestedBlocks(in []*tfprotov5.SchemaNestedBlock) ([]*tfplugin5.Schema_NestedBlock, error) {
	resp := make([]*tfplugin5.Schema_NestedBlock, 0, len(in))

	for _, b := range in {
		block, err := Schema_NestedBlock(b)

		if err != nil {
			return nil, err
		}

		resp = append(resp, block)
	}

	return resp, nil
}

func Schema_NestedBlock_NestingMode(in tfprotov5.SchemaNestedBlockNestingMode) tfplugin5.Schema_NestedBlock_NestingMode {
	return tfplugin5.Schema_NestedBlock_NestingMode(in)
}
