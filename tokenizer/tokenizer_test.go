package tokenizer

import "testing"

func assertGetToken(t *testing.T, input string, want Token) {
	tokenizer := NewTokenizer([]byte(input))
	got, err := tokenizer.Next()

	if err != nil {
		t.Fatalf("NextTokenOf(\"%s\") returned error: %v", input, err)
	}
	if got == nil || *got != want {
		t.Errorf("NextTokenOf(\"%s\") == %v, want &%v", input, got, want)
	}
}

func TestTokenize_EOF(t *testing.T) {
	assertGetToken(t, " ", Token{kind: EOF})
}

func TestTokenize_beginArray(t *testing.T) {
	assertGetToken(t, "[", Token{kind: BeginArray})
}

func TestTokenize_endArray(t *testing.T) {
	assertGetToken(t, "]", Token{kind: EndArray})
}

func TestTokenize_beginObject(t *testing.T) {
	assertGetToken(t, "{", Token{kind: BeginObject})
}

func TestTokenize_endObject(t *testing.T) {
	assertGetToken(t, "}", Token{kind: EndObject})
}

func TestTokenize_nameSeparator(t *testing.T) {
	assertGetToken(t, ":", Token{kind: NameSeparator})
}

func TestTokenize_valueSeparator(t *testing.T) {
	assertGetToken(t, ",", Token{kind: ValueSeparator})
}

func TestTokenize_null(t *testing.T) {
	assertGetToken(t, "null", Token{kind: Null})
}

func TestTokenize_boolean_true(t *testing.T) {
	assertGetToken(t, "true", Token{kind: Boolean, value: true})
}

func TestTokenize_boolean_false(t *testing.T) {
	assertGetToken(t, "false", Token{kind: Boolean, value: false})
}

func TestTokenize_number_int(t *testing.T) {
	assertGetToken(t, "0", Token{kind: Number, value: float64(0)})
	assertGetToken(t, "123", Token{kind: Number, value: float64(123)})
	assertGetToken(t, "-123", Token{kind: Number, value: float64(-123)})
}

func TestTokenize_number_float(t *testing.T) {
	assertGetToken(t, "0.1", Token{kind: Number, value: float64(0.1)})
	assertGetToken(t, "-0.1", Token{kind: Number, value: float64(-0.1)})
	assertGetToken(t, "-123.456e78", Token{kind: Number, value: float64(-123.456e78)})
	assertGetToken(t, "0.123e+45", Token{kind: Number, value: float64(0.123e+45)})
	assertGetToken(t, "0.123E-45", Token{kind: Number, value: float64(0.123e-45)})
}

func TestTokenize_string(t *testing.T) {
	assertGetToken(t, "\"\"", Token{kind: String, value: ""})
	assertGetToken(t, "\"foo\"", Token{kind: String, value: "foo"})
	assertGetToken(t, "\"foo\\\"bar\"", Token{kind: String, value: "foo\"bar"})
}
