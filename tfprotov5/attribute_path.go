package tfprotov5

import (
	"errors"
)

var (
	ErrNotAttributePathStepper = errors.New("doesn't fill tfprotov5.AttributePathStepper interface")
	ErrInvalidStep             = errors.New("step cannot be applied to this value")
)

type AttributePath struct {
	Steps []AttributePathStep
}

type AttributePathStep interface {
	unfulfillable() // make this interface fillable only by this package
}

var (
	_ AttributePathStep = AttributeName("")
	_ AttributePathStep = ElementKeyString("")
	_ AttributePathStep = ElementKeyInt(0)
)

type AttributeName string

func (a AttributeName) unfulfillable() {}

type ElementKeyString string

func (e ElementKeyString) unfulfillable() {}

type ElementKeyInt int64

func (e ElementKeyInt) unfulfillable() {}

type AttributePathStepper interface {
	// Return the attribute or element the AttributePathStep is referring
	// to, or an error if the AttributePathStep is referring to an
	// attribute or element that doesn't exist.
	ApplyTerraform5AttributePathStep(AttributePathStep) (interface{}, error)
}

func WalkAttributePath(in interface{}, path AttributePath) (interface{}, AttributePath, error) {
	if len(path.Steps) < 1 {
		return in, path, nil
	}
	stepper, ok := in.(AttributePathStepper)
	if !ok {
		stepper, ok = builtinAttributePathStepper(in)
		if !ok {
			return in, path, ErrNotAttributePathStepper
		}
	}
	next, err := stepper.ApplyTerraform5AttributePathStep(path.Steps[0])
	if err != nil {
		return in, path, err
	}
	path.Steps = path.Steps[1:]
	return WalkAttributePath(next, path)
}

func builtinAttributePathStepper(in interface{}) (AttributePathStepper, bool) {
	switch v := in.(type) {
	case map[string]interface{}:
		return mapStringInterfaceAttributePathStepper(v), true
	case []interface{}:
		return interfaceSliceAttributePathStepper(v), true
	default:
		return nil, false
	}
}

type mapStringInterfaceAttributePathStepper map[string]interface{}

func (m mapStringInterfaceAttributePathStepper) ApplyTerraform5AttributePathStep(step AttributePathStep) (interface{}, error) {
	_, isAttributeName := step.(AttributeName)
	_, isElementKeyString := step.(ElementKeyString)
	if !isAttributeName && !isElementKeyString {
		return nil, ErrInvalidStep
	}
	var stepValue string
	if isAttributeName {
		stepValue = string(step.(AttributeName))
	}
	if isElementKeyString {
		stepValue = string(step.(ElementKeyString))
	}
	v, ok := m[stepValue]
	if !ok {
		return nil, ErrInvalidStep
	}
	return v, nil
}

type interfaceSliceAttributePathStepper []interface{}

func (i interfaceSliceAttributePathStepper) ApplyTerraform5AttributePathStep(step AttributePathStep) (interface{}, error) {
	eki, isElementKeyInt := step.(ElementKeyInt)
	if !isElementKeyInt {
		return nil, ErrInvalidStep
	}
	// slices can only have items up to the max value of int
	// but we get ElementKeyInt as an int64
	// we keep ElementKeyInt as an int64 and cast the length of the slice
	// to int64 here because if ElementKeyInt is greater than the max value
	// of int, we will always (correctly) error out here. This lets us
	// confidently cast ElementKeyInt to an int below, knowing we're not
	// truncating data
	if int64(eki) >= int64(len(i)) {
		return nil, ErrInvalidStep
	}
	return i[int(eki)], nil
}
