package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

func Schema(in tfplugin5.Schema) tfprotov5.Schema {
	var resp tfprotov5.Schema
	resp.Version = in.Version
	if in.Block != nil {
		block := SchemaBlock(*in.Block)
		resp.Block = &block
	}
	return resp
}

func SchemaBlock(in tfplugin5.Schema_Block) tfprotov5.SchemaBlock {
	return tfprotov5.SchemaBlock{
		Version:         in.Version,
		Attributes:      SchemaAttributes(in.Attributes),
		BlockTypes:      SchemaNestedBlocks(in.BlockTypes),
		Description:     in.Description,
		DescriptionKind: StringKind(in.DescriptionKind),
		Deprecated:      in.Deprecated,
	}
}

func SchemaAttribute(in tfplugin5.Schema_Attribute) tfprotov5.SchemaAttribute {
	return tfprotov5.SchemaAttribute{
		Name:            in.Name,
		Type:            in.Type,
		Description:     in.Description,
		Required:        in.Required,
		Optional:        in.Optional,
		Computed:        in.Computed,
		Sensitive:       in.Sensitive,
		DescriptionKind: StringKind(in.DescriptionKind),
		Deprecated:      in.Deprecated,
	}
}

func SchemaAttributes(in []*tfplugin5.Schema_Attribute) []*tfprotov5.SchemaAttribute {
	resp := make([]*tfprotov5.SchemaAttribute, 0, len(in))
	for _, a := range in {
		if a == nil {
			resp = append(resp, nil)
			continue
		}
		attr := SchemaAttribute(*a)
		resp = append(resp, &attr)
	}
	return resp
}

func SchemaNestedBlock(in tfplugin5.Schema_NestedBlock) tfprotov5.SchemaNestedBlock {
	resp := tfprotov5.SchemaNestedBlock{
		TypeName: in.TypeName,
		Nesting:  SchemaNestedBlockNestingMode(in.Nesting),
		MinItems: in.MinItems,
		MaxItems: in.MaxItems,
	}
	if in.Block != nil {
		block := SchemaBlock(*in.Block)
		resp.Block = &block
	}
	return resp
}

func SchemaNestedBlocks(in []*tfplugin5.Schema_NestedBlock) []*tfprotov5.SchemaNestedBlock {
	resp := make([]*tfprotov5.SchemaNestedBlock, 0, len(in))
	for _, b := range in {
		if b == nil {
			resp = append(resp, nil)
			continue
		}
		block := SchemaNestedBlock(*b)
		resp = append(resp, &block)
	}
	return resp
}

func SchemaNestedBlockNestingMode(in tfplugin5.Schema_NestedBlock_NestingMode) tfprotov5.SchemaNestedBlockNestingMode {
	return tfprotov5.SchemaNestedBlockNestingMode(in)
}
