package encoder

import "testing"

func assertEncode(t *testing.T, input any, want string) {
	got, err := Encode(input)

	if err != nil {
		t.Fatalf("Encode returned error: %v", err)
	}
	if got != want {
		t.Errorf("Encode(%v) == %q, want %q", input, got, want)
	}
}

func TestEncode_Bool(t *testing.T) {
	assertEncode(t, true, "true")
}

func TestEncode_String(t *testing.T) {
	assertEncode(t, "hello", `"hello"`)
}

func TestEncode_StringWithEscape(t *testing.T) {
	assertEncode(t, `hello "world"`, `"hello \"world\""`)
}

func TestEncode_Int(t *testing.T) {
	assertEncode(t, 42, "42")
}

func TestEncode_Float(t *testing.T) {
	assertEncode(t, 3.14, "3.14")
}

func TestEncode_StringArray(t *testing.T) {
	assertEncode(t, []string{"hello", "world"}, `["hello","world"]`)
}

func TestEncode_Struct(t *testing.T) {
	type model struct {
		Name  string
		Value int
	}
	assertEncode(t, model{"foo", 42}, `{"name":"foo","value":42}`)
}

func TestEncode_TaggedStruct(t *testing.T) {
	type model struct {
		Name  string
		Value int `json:"bar"`
	}
	assertEncode(t, model{"foo", 42}, `{"name":"foo","bar":42}`)
}
