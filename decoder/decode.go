package decoder

import (
	"errors"
	"json/tokenizer"
	"reflect"
)

var ErrUnexpectedTokenType = errors.New("Unexpected token type")
var ErrUnexpectedTargetType = errors.New("Unexpected target type")

func decodeBoolean(t *tokenizer.Tokenizer, out *bool) error {
	token, err := t.Next()
	if err != nil {
		return err
	}
	if token.Kind == tokenizer.BooleanToken {
		*out = token.BoolValue()
	} else {
		return ErrUnexpectedTokenType
	}
	return nil
}

func decodeString(t *tokenizer.Tokenizer, out *string) error {
	token, err := t.Next()
	if err != nil {
		return err
	}

	if token.Kind == tokenizer.StringToken {
		*out = token.StringValue()
	} else {
		return ErrUnexpectedTokenType
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
		return ErrUnexpectedTokenType
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
	case reflect.Bool:
		return decodeBoolean(tokenizer.NewTokenizer(input), out.(*bool))
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

	return ErrUnexpectedTargetType
}
