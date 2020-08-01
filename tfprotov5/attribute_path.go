package tfprotov5

type AttributePath struct {
	Steps []*AttributePathStep
}

type AttributePathStep struct {
	// Only one of these should be set
	AttributeName    string
	ElementKeyString string
	ElementKeyInt    int64
}
