package strings

import "testing"

func TestEscapeString_NormalString(t *testing.T) {
	input := []byte("Hello, World!")
	got := EscapeString(input)
	want := []byte("Hello, World!")

	if string(got) != string(want) {
		t.Errorf("escapeString(%v) == %v, want %v", input, got, want)
	}
}

func TestEscapeString_Escaped(t *testing.T) {
	input := []byte("\"\\/\b\f\n\r\t")
	got := EscapeString(input)
	want := []byte("\\\"\\\\\\/\\\b\\\f\\\n\\\r\\\t")

	if string(got) != string(want) {
		t.Errorf("escapeString(%v) == %v, want %v", input, got, want)
	}
}

func TestUnescapeString_NormalString(t *testing.T) {
	input := []byte("Hello, World!")
	got, err := UnescapeString(input)
	want := []byte("Hello, World!")

	if err != nil {
		t.Fatalf("unescapeString(%v) returned an error: %v", input, err)
	}
	if string(got) != string(want) {
		t.Errorf("unescapeString(%v) == %v, want %v", input, got, want)
	}
}

func TestUnescapeString_Escaped(t *testing.T) {
	input := []byte("\\\"\\\\\\/\\b\\f\\n\\r\\t")
	got, err := UnescapeString(input)
	want := []byte("\"\\/\b\f\n\r\t")

	if err != nil {
		t.Fatalf("unescapeString(%v) returned an error: %v", input, err)
	}
	if string(got) != string(want) {
		t.Errorf("unescapeString(%v) == %v, want %v", input, got, want)
	}
}
