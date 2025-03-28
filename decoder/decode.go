package decoder

import (
	"errors"
	"reflect"

	. "json/strings"
	"json/tokenizer"
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

	if _, err := expectToken(t, tokenizer.BeginArrayToken); err != nil {
		return err
	}

	for {
		// item should be *T
		item := reflect.New(itemType).Interface()
		if err := decode(t, item); err != nil {
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

func decodeObject(t *tokenizer.Tokenizer, out any) error {
	outTypePtr := reflect.TypeOf(out)
	if outTypePtr.Kind() != reflect.Ptr {
		panic("out must be a pointer")
	}
	outValue := reflect.ValueOf(out).Elem()

	outType := outTypePtr.Elem()
	if outType.Kind() != reflect.Struct {
		panic("out must be a struct")
	}

	tagFieldNames := map[string]string{}
	numFields := outType.NumField()
	for i := 0; i < numFields; i++ {
		field := outType.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" {
			tagFieldNames[tag] = field.Name
		}
	}

	if _, err := expectToken(t, tokenizer.BeginObjectToken); err != nil {
		return err
	}

	for {
		nameToken, err := expectToken(t, tokenizer.StringToken)
		if err != nil {
			return err
		}

		name := nameToken.StringValue()
		fieldName := tagFieldNames[name]
		if fieldName == "" {
			fieldName = SnakeCaseToUpperCamelCase(name)
		}

		if _, err := expectToken(t, tokenizer.NameSeparatorToken); err != nil {
			return err
		}

		field := outValue.FieldByName(fieldName)
		if !field.IsValid() {
			return ErrUnexpectedTargetType
		}
		if err = decode(t, field.Addr().Interface()); err != nil {
			return err
		}

		token, err := t.Next()
		if err != nil {
			return err
		}
		if token.Kind == tokenizer.EndObjectToken {
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
	case reflect.Struct:
		return decodeObject(t, out)
	}

	return ErrUnexpectedTargetType
}

func Decode(input []byte, out any) error {
	t := tokenizer.NewTokenizer(input)
	return decode(t, out)
}
