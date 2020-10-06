package toproto

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

var ErrUnknownAttributePathStepType = errors.New("unknown type of AttributePath_Step")

func AttributePath(in *tfprotov5.AttributePath) (*tfplugin5.AttributePath, error) {
	steps, err := AttributePath_Steps(in.Steps)
	if err != nil {
		return nil, err
	}
	return &tfplugin5.AttributePath{
		Steps: steps,
	}, nil
}

func AttributePaths(in []*tfprotov5.AttributePath) ([]*tfplugin5.AttributePath, error) {
	resp := make([]*tfplugin5.AttributePath, 0, len(in))
	for _, a := range in {
		if a == nil {
			resp = append(resp, nil)
			continue
		}
		attr, err := AttributePath(a)
		if err != nil {
			return resp, err
		}
		resp = append(resp, attr)
	}
	return resp, nil
}

func AttributePath_Step(step tfprotov5.AttributePathStep) (*tfplugin5.AttributePath_Step, error) {
	var resp tfplugin5.AttributePath_Step
	if name, ok := step.(tfprotov5.AttributeName); ok {
		resp.Selector = &tfplugin5.AttributePath_Step_AttributeName{
			AttributeName: string(name),
		}
		return &resp, nil
	}
	if key, ok := step.(tfprotov5.ElementKeyString); ok {
		resp.Selector = &tfplugin5.AttributePath_Step_ElementKeyString{
			ElementKeyString: string(key),
		}
		return &resp, nil
	}
	if key, ok := step.(tfprotov5.ElementKeyInt); ok {
		resp.Selector = &tfplugin5.AttributePath_Step_ElementKeyInt{
			ElementKeyInt: int64(key),
		}
		return &resp, nil
	}
	return nil, ErrUnknownAttributePathStepType
}

func AttributePath_Steps(in []tfprotov5.AttributePathStep) ([]*tfplugin5.AttributePath_Step, error) {
	resp := make([]*tfplugin5.AttributePath_Step, 0, len(in))
	for _, step := range in {
		if step == nil {
			resp = append(resp, nil)
			continue
		}
		s, err := AttributePath_Step(step)
		if err != nil {
			return resp, err
		}
		resp = append(resp, s)
	}
	return resp, nil
}
