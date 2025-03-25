package json

import (
	"json/decoder"
	"json/encoder"
)

type StringLike interface {
	string | []byte
}

func Decode[T StringLike](input T, out any) error {
	return decoder.Decode([]byte(input), out)
}

func Encode(v any) (string, error) {
	return encoder.Encode(v)
}
