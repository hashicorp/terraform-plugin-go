package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func AttributePath(in tfprotov5.AttributePath) tfplugin5.AttributePath {
	return tfplugin5.AttributePath{
		Steps: AttributePath_Steps(in.Steps),
	}
}

func AttributePaths(in []*tfprotov5.AttributePath) []*tfplugin5.AttributePath {
	resp := make([]*tfplugin5.AttributePath, 0, len(in))
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

func AttributePath_Step(step tfprotov5.AttributePathStep) tfplugin5.AttributePath_Step {
	var resp tfplugin5.AttributePath_Step
	if step.AttributeName != "" {
		resp.Selector = &tfplugin5.AttributePath_Step_AttributeName{
			AttributeName: step.AttributeName,
		}
	}
	if step.ElementKeyString != "" {
		resp.Selector = &tfplugin5.AttributePath_Step_ElementKeyString{
			ElementKeyString: step.ElementKeyString,
		}
	}
	if step.ElementKeyInt != 0 {
		resp.Selector = &tfplugin5.AttributePath_Step_ElementKeyInt{
			ElementKeyInt: step.ElementKeyInt,
		}
	}
	return resp
}

func AttributePath_Steps(in []*tfprotov5.AttributePathStep) []*tfplugin5.AttributePath_Step {
	resp := make([]*tfplugin5.AttributePath_Step, 0, len(in))
	for _, step := range in {
		if step == nil {
			resp = append(resp, nil)
			continue
		}
		s := AttributePath_Step(*step)
		resp = append(resp, &s)
	}
	return resp
}
