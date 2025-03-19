package tfprotov6

const (
	ActionCancelTypeSoft CancelType = 0
	ActionCancelTypeHard CancelType = 1
)

type ActionCancelType struct {
	CancelType CancelType
}

type CancelType int32

func (c CancelType) String() string {
	switch c {
	case 0:
		return "SOFT"
	case 1:
		return "HARD"
	}
	return "UNKNOWN"
}
