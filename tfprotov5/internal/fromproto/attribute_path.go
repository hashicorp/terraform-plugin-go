package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

func AttributePath(in tfplugin5.AttributePath) tfprotov5.AttributePath {
	return tfprotov5.AttributePath{
		Steps: AttributePathSteps(in.Steps),
	}
}

func AttributePaths(in []*tfplugin5.AttributePath) []*tfprotov5.AttributePath {
	resp := make([]*tfprotov5.AttributePath, 0, len(in))
	for _, a := range in {
		if a == nil {
			resp = append(resp, nil)
			continue
		}
		attr := AttributePath(*a)
		resp = append(resp, &attr)
	}
	return resp
}

func AttributePathStep(step tfplugin5.AttributePath_Step) tfprotov5.AttributePathStep {
	return tfprotov5.AttributePathStep{
		AttributeName:    step.GetAttributeName(),
		ElementKeyString: step.GetElementKeyString(),
		ElementKeyInt:    step.GetElementKeyInt(),
	}
}

func AttributePathSteps(in []*tfplugin5.AttributePath_Step) []*tfprotov5.AttributePathStep {
	resp := make([]*tfprotov5.AttributePathStep, 0, len(in))
	for _, step := range in {
		if step == nil {
			resp = append(resp, nil)
			continue
		}
		s := AttributePathStep(*step)
		resp = append(resp, &s)
	}
	return resp
}
