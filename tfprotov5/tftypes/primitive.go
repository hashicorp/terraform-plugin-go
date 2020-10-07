package tftypes

import "fmt"

const (
	UnknownType       = primitive("Unknown")
	DynamicPseudoType = primitive("DynamicPseudoType")
	String            = primitive("String")
	Number            = primitive("Number")
	Bool              = primitive("Bool")
)

var (
	_ Type = primitive("test")
)

type primitive string

func (p primitive) Is(t Type) bool {
	v, ok := t.(primitive)
	if !ok {
		return false
	}
	return p == v
}

func (p primitive) String() string {
	return "tftypes." + string(p)
}

func (p primitive) private() {}

func (p primitive) MarshalJSON() ([]byte, error) {
	switch p {
	case String:
		return []byte(`"string"`), nil
	case Number:
		return []byte(`"number"`), nil
	case Bool:
		return []byte(`"bool"`), nil
	case DynamicPseudoType:
		return []byte(`"dynamic"`), nil
	}
	return nil, fmt.Errorf("unknown primitive type %q", p)
}
