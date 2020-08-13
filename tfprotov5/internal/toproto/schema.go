package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func Schema(in tfprotov5.Schema) tfplugin5.Schema {
	var resp tfplugin5.Schema
	resp.Version = in.Version
	if in.Block != nil {
		block := Schema_Block(*in.Block)
		resp.Block = &block
	}
	return resp
}

func Schema_Block(in tfprotov5.SchemaBlock) tfplugin5.Schema_Block {
	return tfplugin5.Schema_Block{
		Version:         in.Version,
		Attributes:      Schema_Attributes(in.Attributes),
		BlockTypes:      Schema_NestedBlocks(in.BlockTypes),
		Description:     in.Description,
		DescriptionKind: StringKind(in.DescriptionKind),
		Deprecated:      in.Deprecated,
	}
}

func Schema_Attribute(in tfprotov5.SchemaAttribute) tfplugin5.Schema_Attribute {
	return tfplugin5.Schema_Attribute{
		Name:            in.Name,
		Type:            CtyType(in.Type),
		Description:     in.Description,
		Required:        in.Required,
		Optional:        in.Optional,
		Computed:        in.Computed,
		Sensitive:       in.Sensitive,
		DescriptionKind: StringKind(in.DescriptionKind),
		Deprecated:      in.Deprecated,
	}
}

func Schema_Attributes(in []*tfprotov5.SchemaAttribute) []*tfplugin5.Schema_Attribute {
	resp := make([]*tfplugin5.Schema_Attribute, 0, len(in))
	for _, a := range in {
		if a == nil {
			resp = append(resp, nil)
			continue
		}
		attr := Schema_Attribute(*a)
		resp = append(resp, &attr)
	}
	return resp
}

func Schema_NestedBlock(in tfprotov5.SchemaNestedBlock) tfplugin5.Schema_NestedBlock {
	resp := tfplugin5.Schema_NestedBlock{
		TypeName: in.TypeName,
		Nesting:  Schema_NestedBlock_NestingMode(in.Nesting),
		MinItems: in.MinItems,
		MaxItems: in.MaxItems,
	}
	if in.Block != nil {
		block := Schema_Block(*in.Block)
		resp.Block = &block
	}
	return resp
}

func Schema_NestedBlocks(in []*tfprotov5.SchemaNestedBlock) []*tfplugin5.Schema_NestedBlock {
	resp := make([]*tfplugin5.Schema_NestedBlock, 0, len(in))
	for _, b := range in {
		if b == nil {
			resp = append(resp, nil)
			continue
		}
		block := Schema_NestedBlock(*b)
		resp = append(resp, &block)
	}
	return resp
}

func Schema_NestedBlock_NestingMode(in tfprotov5.SchemaNestedBlockNestingMode) tfplugin5.Schema_NestedBlock_NestingMode {
	return tfplugin5.Schema_NestedBlock_NestingMode(in)
}
