package tftypes

import "fmt"

type Path []interface{}

func (p Path) AddIntStep(i int64) Path {
	return append(p, i)
}

func (p Path) AddStringStep(s string) Path {
	return append(p, s)
}

func (p Path) AddValueStep(v Value) Path {
	return append(p, v)
}

func (p Path) RemoveLastStep() Path {
	return p[:len(p)-1]
}

func (p Path) NewError(err error) error {
	return pathError{
		error: err,
		path:  p,
	}
}

func (p Path) NewErrorf(f string, args ...interface{}) error {
	return pathError{
		error: fmt.Errorf(f, args...),
		path:  p,
	}
}

type pathError struct {
	error
	path Path
}
