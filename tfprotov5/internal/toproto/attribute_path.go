// Copyright (c) HashiCorp, Inc.

package toproto

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var ErrUnknownAttributePathStepType = errors.New("unknown type of AttributePath_Step")

func AttributePath(in *tftypes.AttributePath) (*tfplugin5.AttributePath, error) {
	if in == nil {
		return nil, nil
	}

	steps, err := AttributePath_Steps(in.Steps())

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.AttributePath{
		Steps: steps,
	}

	return resp, nil
}

func AttributePaths(in []*tftypes.AttributePath) ([]*tfplugin5.AttributePath, error) {
	resp := make([]*tfplugin5.AttributePath, 0, len(in))

	for _, a := range in {
		attr, err := AttributePath(a)

		if err != nil {
			return resp, err
		}

		resp = append(resp, attr)
	}

	return resp, nil
}

func AttributePath_Step(step tftypes.AttributePathStep) (*tfplugin5.AttributePath_Step, error) {
	var resp tfplugin5.AttributePath_Step
	if name, ok := step.(tftypes.AttributeName); ok {
		resp.Selector = &tfplugin5.AttributePath_Step_AttributeName{
			AttributeName: string(name),
		}
		return &resp, nil
	}
	if key, ok := step.(tftypes.ElementKeyString); ok {
		resp.Selector = &tfplugin5.AttributePath_Step_ElementKeyString{
			ElementKeyString: string(key),
		}
		return &resp, nil
	}
	if key, ok := step.(tftypes.ElementKeyInt); ok {
		resp.Selector = &tfplugin5.AttributePath_Step_ElementKeyInt{
			ElementKeyInt: int64(key),
		}
		return &resp, nil
	}
	if _, ok := step.(tftypes.ElementKeyValue); ok {
		// the protocol has no equivalent of an ElementKeyValue, so we
		// return nil for both the step and the error here, to signal
		// that we've hit a step we can't convey back to Terraform
		return nil, nil
	}
	return nil, ErrUnknownAttributePathStepType
}

func AttributePath_Steps(in []tftypes.AttributePathStep) ([]*tfplugin5.AttributePath_Step, error) {
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
		// in the face of a set, the protocol has no way to represent
		// the index, so we just bail and return the prefix we can
		// return.
		if s == nil {
			return resp, nil
		}
		resp = append(resp, s)
	}
	return resp, nil
}
