package fromproto

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

var ErrUnknownAttributePathStepType = errors.New("unknown type of AttributePath_Step")

func AttributePath(in *tfplugin5.AttributePath) (*tfprotov5.AttributePath, error) {
	steps, err := AttributePathSteps(in.Steps)
	if err != nil {
		return nil, err
	}
	return &tfprotov5.AttributePath{
		Steps: steps,
	}, nil
}

func AttributePaths(in []*tfplugin5.AttributePath) ([]*tfprotov5.AttributePath, error) {
	resp := make([]*tfprotov5.AttributePath, 0, len(in))
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

func AttributePathStep(step *tfplugin5.AttributePath_Step) (tfprotov5.AttributePathStep, error) {
	selector := step.GetSelector()
	if v, ok := selector.(*tfplugin5.AttributePath_Step_AttributeName); ok {
		return tfprotov5.AttributeName(v.AttributeName), nil
	}
	if v, ok := selector.(*tfplugin5.AttributePath_Step_ElementKeyString); ok {
		return tfprotov5.ElementKeyString(v.ElementKeyString), nil
	}
	if v, ok := selector.(*tfplugin5.AttributePath_Step_ElementKeyInt); ok {
		return tfprotov5.ElementKeyInt(v.ElementKeyInt), nil
	}
	return nil, ErrUnknownAttributePathStepType
}

func AttributePathSteps(in []*tfplugin5.AttributePath_Step) ([]tfprotov5.AttributePathStep, error) {
	resp := make([]tfprotov5.AttributePathStep, 0, len(in))
	for _, step := range in {
		if step == nil {
			continue
		}
		s, err := AttributePathStep(step)
		if err != nil {
			return resp, err
		}
		resp = append(resp, s)
	}
	return resp, nil
}
