package refinement

import "github.com/vmihailenco/msgpack/v5"

type nullness struct {
	Value bool
}

func (n nullness) Encode(enc *msgpack.Encoder) error {
	err := enc.EncodeInt(int64(KeyNullness))
	if err != nil {
		return err
	}

	// A value that is definitely null cannot be unknown
	return enc.EncodeBool(false)
}

func (n nullness) Equal(Refinement) bool {
	return false
}

func (n nullness) String() string {
	return "todo - nullness"
}

func (n nullness) unimplementable() {}

// TODO: Should this accept a value? If a value is unknown and the it's refined to be null
// then the value should be a known value of null instead.
func Nullness(value bool) Refinement {
	return nullness{
		Value: value,
	}
}
