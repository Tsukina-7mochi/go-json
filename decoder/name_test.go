package decoder

import "testing"

func assertSnakeCaseToUpperCamelCase(t *testing.T, input string, want string) {
	got := snakeCaseToUpperCamelCase(input)
	if got != want {
		t.Errorf("snakeCaseToCamelCase == %q, want %q", got, want)
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
