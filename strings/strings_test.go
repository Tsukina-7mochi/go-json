package strings

import "testing"

func assertSnakeCaseToUpperCamelCase(t *testing.T, input string, want string) {
	got := SnakeCaseToUpperCamelCase(input)
	if got != want {
		t.Errorf("SnakeCaseToCamelCase == %q, want %q", got, want)
	}
}

func assertCamelCaseToSnakeCase(t *testing.T, input string, want string) {
	got := CamelCaseToSnakeCase(input)
	if got != want {
		t.Errorf("CamelCaseToSnakeCase == %q, want %q", got, want)
	}
}

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

func TestSnakeCaseToUpperCamelCase_1(t *testing.T) {
	assertSnakeCaseToUpperCamelCase(t, "hello_world", "HelloWorld")
}

func TestSnakeCaseToUpperCamelCase_2(t *testing.T) {
	assertSnakeCaseToUpperCamelCase(t, "HelloWorld", "HelloWorld")
}

func TestSnakeCaseToUpperCamelCase_3(t *testing.T) {
	assertSnakeCaseToUpperCamelCase(t, "__hello_world", "HelloWorld")
}

func TestCamelCaseToSnakeCase_1(t *testing.T) {
	assertCamelCaseToSnakeCase(t, "HelloWorld", "hello_world")
}

func TestCamelCaseToSnakeCase_2(t *testing.T) {
	assertCamelCaseToSnakeCase(t, "helloWorld", "hello_world")
}

func TestCamelCaseToSnakeCase_3(t *testing.T) {
	assertCamelCaseToSnakeCase(t, "hello_world", "hello_world")
}
