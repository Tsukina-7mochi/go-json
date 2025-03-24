package decoder

import (
	"errors"
	"json/tokenizer"
	"reflect"
)

var ErrUnexpectedTokenType = errors.New("Unexpected token type")
var ErrUnexpectedTargetType = errors.New("Unexpected target type")

func expectToken(t *tokenizer.Tokenizer, kind tokenizer.TokenKind) (*tokenizer.Token, error) {
	token, err := t.Next()
	if err != nil {
		return nil, err
	}
	if token.Kind != kind {
		return nil, ErrUnexpectedTokenType
	}

	return token, nil
}

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

func decodeArray(t *tokenizer.Tokenizer, out any) error {
	// Assume we are decoding []T.
	// The type of `out` is *[]T (not necessarily a pointer, but it is treated as
	// such for consistency with other types).

	// outTypePtr should be *[]t
	outTypePtr := reflect.TypeOf(out)
	if outTypePtr.Kind() != reflect.Ptr {
		panic("out must be a pointer")
	}

	// outType and outValue should be []T
	outType := outTypePtr.Elem()
	if outType.Kind() != reflect.Slice {
		panic("out must be a slice")
	}
	outValue := reflect.ValueOf(out).Elem()

	// itemType should be T
	itemType := outType.Elem()

	_, err := expectToken(t, tokenizer.BeginArrayToken)
	if err != nil {
		return err
	}

	for {
		// item should be *T
		item := reflect.New(itemType).Interface()
		err := decode(t, item)
		if err != nil {
			return err
		}

		// it is equal to *out = append(*out, *item)
		outValue.Set(reflect.Append(outValue, reflect.ValueOf(item).Elem()))

		token, err := t.Next()
		if err != nil {
			return err
		}
		if token.Kind == tokenizer.EndArrayToken {
			break
		} else if token.Kind != tokenizer.ValueSeparatorToken {
			return ErrUnexpectedTokenType
		}
	}

	return nil
}

func decode(t *tokenizer.Tokenizer, out any) error {
	outTypePtr := reflect.TypeOf(out)
	if outTypePtr.Kind() != reflect.Ptr {
		panic("out must be a pointer")
	}

	outType := outTypePtr.Elem()

	switch outType.Kind() {
	case reflect.Bool:
		return decodeBoolean(t, out.(*bool))
	case reflect.String:
		return decodeString(t, out.(*string))
	case reflect.Int:
		var num float64
		err := decodeNumber(t, &num)
		*out.(*int) = int(num)
		return err
	case reflect.Float64:
		return decodeNumber(t, out.(*float64))
	case reflect.Slice:
		return decodeArray(t, out)
	}

	return ErrUnexpectedTargetType
}

func Decode(input []byte, out any) error {
	t := tokenizer.NewTokenizer(input)
	return decode(t, out)
}
