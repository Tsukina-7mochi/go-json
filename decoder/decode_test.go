package decoder

import "testing"

func TestDecode_String(t *testing.T) {
	var str string
	err := Decode([]byte(`"hello"`), &str)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if str != "hello" {
		t.Errorf("Decode == %q, want %q", str, "hello")
	}
}
