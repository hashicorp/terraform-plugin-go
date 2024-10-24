package refinement

import "github.com/vmihailenco/msgpack/v5"

type Nullness struct {
	value bool
}

func (n Nullness) Encode(enc *msgpack.Encoder) error {
	err := enc.EncodeInt(int64(KeyNullness))
	if err != nil {
		return err
	}

	// It shouldn't be possible for an unknown value to be definitely null (i.e. nullness.value = true),
	// as that should be represented by a known null value instead. This encoding is in place to be compliant
	// with Terraform's encoding which uses a definitely null refinement to collapse into a known null value.
	return enc.EncodeBool(n.value)
}

func (n Nullness) Equal(Refinement) bool {
	return false
}

func (n Nullness) String() string {
	return "todo - Nullness"
}

func (n Nullness) NotNull() bool {
	return !n.value
}

func (n Nullness) unimplementable() {}

func NewNullness(value bool) Refinement {
	return Nullness{
		value: value,
	}
}
