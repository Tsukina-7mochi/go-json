package tokenizer

import "testing"

func assertGetToken(t *testing.T, input string, want Token) {
	tokenizer := NewTokenizer([]byte(input))
	got, err := tokenizer.Next()

	if err != nil {
		t.Fatalf("NewTokenizer(\"%s\").Next() returned error: %v", input, err)
	}
	if got == nil || *got != want {
		t.Errorf("NewTokenizer(\"%s\").Next() == %v, want &%v", input, got, want)
	}
}

func assertGetTokenSlice(t *testing.T, input string, want []Token) {
	tokenizer := NewTokenizer([]byte(input))

	tokens := make([]Token, 0)
	for {
		token, err := tokenizer.Next()
		if err != nil {
			t.Fatalf("NewTokenizer(\"%s\").Next() returned error: %v", input, err)
		}
		print(token.Kind)
		tokens = append(tokens, *token)

		if token.Kind == EOFToken {
			break
		}
	}

	if len(tokens) != len(want) {
		t.Fatalf("{...NewTokenizer(\"%s\").Next()} == %v, want %v", input, tokens, want)
		return
	}
	for i := range tokens {
		if tokens[i] != want[i] {
			t.Errorf("{...NewTokenizer(\"%s\").Next()} == %v, want %v", input, tokens, want)
			return
		}
	}
}

func TestTokenize_EOF(t *testing.T) {
	assertGetToken(t, " ", Token{Kind: EOFToken})
}

func TestTokenize_beginArray(t *testing.T) {
	assertGetToken(t, "[", Token{Kind: BeginArrayToken})
}

func TestTokenize_endArray(t *testing.T) {
	assertGetToken(t, "]", Token{Kind: EndArrayToken})
}

func TestTokenize_beginObject(t *testing.T) {
	assertGetToken(t, "{", Token{Kind: BeginObjectToken})
}

func TestTokenize_endObject(t *testing.T) {
	assertGetToken(t, "}", Token{Kind: EndObjectToken})
}

func TestTokenize_nameSeparator(t *testing.T) {
	assertGetToken(t, ":", Token{Kind: NameSeparatorToken})
}

func TestTokenize_valueSeparator(t *testing.T) {
	assertGetToken(t, ",", Token{Kind: ValueSeparatorToken})
}

func TestTokenize_null(t *testing.T) {
	assertGetToken(t, "null", Token{Kind: NullToken})
}

func TestTokenize_boolean_true(t *testing.T) {
	assertGetToken(t, "true", Token{Kind: BooleanToken, Value: true})
}

func TestTokenize_boolean_false(t *testing.T) {
	assertGetToken(t, "false", Token{Kind: BooleanToken, Value: false})
}

func TestTokenize_number_int(t *testing.T) {
	assertGetToken(t, "0", Token{Kind: NumberToken, Value: float64(0)})
	assertGetToken(t, "123", Token{Kind: NumberToken, Value: float64(123)})
	assertGetToken(t, "-123", Token{Kind: NumberToken, Value: float64(-123)})
}

func TestTokenize_number_float(t *testing.T) {
	assertGetToken(t, "0.1", Token{Kind: NumberToken, Value: float64(0.1)})
	assertGetToken(t, "-0.1", Token{Kind: NumberToken, Value: float64(-0.1)})
	assertGetToken(t, "-123.456e78", Token{Kind: NumberToken, Value: float64(-123.456e78)})
	assertGetToken(t, "0.123e+45", Token{Kind: NumberToken, Value: float64(0.123e+45)})
	assertGetToken(t, "0.123E-45", Token{Kind: NumberToken, Value: float64(0.123e-45)})
}

func TestTokenize_string(t *testing.T) {
	assertGetToken(t, "\"\"", Token{Kind: StringToken, Value: ""})
	assertGetToken(t, "\"foo\"", Token{Kind: StringToken, Value: "foo"})
	assertGetToken(t, "\"foo\\\"bar\"", Token{Kind: StringToken, Value: "foo\"bar"})
}

func TestTokenizer_ArrayStructure(t *testing.T) {
	assertGetTokenSlice(t, `[null, false, "a", 0, 0.5, [], {}]`, []Token{
		{Kind: BeginArrayToken},
		{Kind: NullToken},
		{Kind: ValueSeparatorToken},
		{Kind: BooleanToken, Value: false},
		{Kind: ValueSeparatorToken},
		{Kind: StringToken, Value: "a"},
		{Kind: ValueSeparatorToken},
		{Kind: NumberToken, Value: float64(0)},
		{Kind: ValueSeparatorToken},
		{Kind: NumberToken, Value: float64(0.5)},
		{Kind: ValueSeparatorToken},
		{Kind: BeginArrayToken},
		{Kind: EndArrayToken},
		{Kind: ValueSeparatorToken},
		{Kind: BeginObjectToken},
		{Kind: EndObjectToken},
		{Kind: EndArrayToken},
		{Kind: EOFToken},
	})
}

func TestTokenizer_ObjectStructure(t *testing.T) {
	assertGetTokenSlice(t, `{"a": null, "b": false, "c": "a", "d": 0, "e": 0.5, "f": [], "g": {}}`, []Token{
		{Kind: BeginObjectToken},
		{Kind: StringToken, Value: "a"},
		{Kind: NameSeparatorToken},
		{Kind: NullToken},
		{Kind: ValueSeparatorToken},
		{Kind: StringToken, Value: "b"},
		{Kind: NameSeparatorToken},
		{Kind: BooleanToken, Value: false},
		{Kind: ValueSeparatorToken},
		{Kind: StringToken, Value: "c"},
		{Kind: NameSeparatorToken},
		{Kind: StringToken, Value: "a"},
		{Kind: ValueSeparatorToken},
		{Kind: StringToken, Value: "d"},
		{Kind: NameSeparatorToken},
		{Kind: NumberToken, Value: float64(0)},
		{Kind: ValueSeparatorToken},
		{Kind: StringToken, Value: "e"},
		{Kind: NameSeparatorToken},
		{Kind: NumberToken, Value: float64(0.5)},
		{Kind: ValueSeparatorToken},
		{Kind: StringToken, Value: "f"},
		{Kind: NameSeparatorToken},
		{Kind: BeginArrayToken},
		{Kind: EndArrayToken},
		{Kind: ValueSeparatorToken},
		{Kind: StringToken, Value: "g"},
		{Kind: NameSeparatorToken},
		{Kind: BeginObjectToken},
		{Kind: EndObjectToken},
		{Kind: EndObjectToken},
		{Kind: EOFToken},
	})
}
