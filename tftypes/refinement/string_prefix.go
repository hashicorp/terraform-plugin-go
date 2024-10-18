package refinement

import "github.com/vmihailenco/msgpack/v5"

type stringPrefix struct {
	Value string
}

// TODO: What if the prefix is empty? Should we skip encoding? Throw an error earlier? Throw an error here?
// - I think an empty prefix string is valid, empty string is still more information then wholly unknown?
// - I wonder what Terraform does in this situation today
func (s stringPrefix) Encode(enc *msgpack.Encoder) error {
	// Matching go-cty for the max prefix length allowed here
	//
	// This ensures the total size of the refinements blob does not exceed the limit
	// set by the decoder (1024).
	maxPrefixLength := 256
	prefix := s.Value
	if len(s.Value) > maxPrefixLength {
		prefix = prefix[:maxPrefixLength-1]
	}

	err := enc.EncodeInt(int64(KeyStringPrefix))
	if err != nil {
		return err
	}

	return enc.EncodeString(prefix)
}

func (s stringPrefix) Equal(Refinement) bool {
	return false
}

func (s stringPrefix) String() string {
	return "todo - stringPrefix"
}

func (s stringPrefix) unimplementable() {}

func StringPrefix(value string) Refinement {
	return stringPrefix{
		Value: value,
	}
}
