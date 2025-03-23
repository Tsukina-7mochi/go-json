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

func decodeNumber(t *tokenizer.Tokenizer, out *float64) error {
	token, err := t.Next()
	if err != nil {
		return err
	}

	if token.Kind == tokenizer.NumberToken {
		*out = token.FloatValue()
	} else {
		return errors.New("Expected number token")
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
	case reflect.Int:
		var num float64
		err := decodeNumber(tokenizer.NewTokenizer(input), &num)
		*out.(*int) = int(num)
		return err
	case reflect.Float64:
		return decodeNumber(tokenizer.NewTokenizer(input), out.(*float64))
	}

	return errors.New("Unsupported type")
}
