package decoder

import (
	"errors"
	"json/tokenizer"
	"reflect"
)

func decodeString(t *tokenizer.Tokenizer, out *string) error {
	token, err := t.Next()
	if err != nil {
		return err
	}

	if token.Kind == tokenizer.StringToken {
		*out = token.StringValue()
	} else {
		return errors.New("Expected string token")
	}

	return nil
}

func Decode(input []byte, out any) error {
	outTypePtr := reflect.TypeOf(out)
	if outTypePtr.Kind() != reflect.Ptr {
		panic("out must be a pointer")
	}

	outType := outTypePtr.Elem()

	switch outType.Kind() {
	case reflect.String:
		return decodeString(tokenizer.NewTokenizer(input), out.(*string))
	}

	return errors.New("Unsupported type")
}
