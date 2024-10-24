package refinement

import "github.com/vmihailenco/msgpack/v5"

type StringPrefix struct {
	value string
}

// TODO: What if the prefix is empty? Should we skip encoding? Throw an error earlier? Throw an error here?
// - I think an empty prefix string is valid, empty string is still more information then wholly unknown?
// - I wonder what Terraform does in this situation today
func (s StringPrefix) Encode(enc *msgpack.Encoder) error {
	// Matching go-cty for the max prefix length allowed here
	//
	// This ensures the total size of the refinements blob does not exceed the limit
	// set by the decoder (1024).
	maxPrefixLength := 256
	prefix := s.value
	if len(s.value) > maxPrefixLength {
		prefix = prefix[:maxPrefixLength-1]
	}

	err := enc.EncodeInt(int64(KeyStringPrefix))
	if err != nil {
		return err
	}

	return enc.EncodeString(prefix)
}

func (s StringPrefix) Equal(Refinement) bool {
	return false
}

func (s StringPrefix) String() string {
	return "todo - stringPrefix"
}

func (s StringPrefix) PrefixValue() string {
	return s.value
}

func (s StringPrefix) unimplementable() {}

func NewStringPrefix(value string) Refinement {
	return StringPrefix{
		value: value,
	}
}
