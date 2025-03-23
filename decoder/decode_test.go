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

func TestDecode_Int(t *testing.T) {
	var num int
	err := Decode([]byte(`123`), &num)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if num != 123 {
		t.Errorf("Decode == %v, want %v", num, 123)
	}
}

func TestDecode_Float64(t *testing.T) {
	var num float64
	err := Decode([]byte(`0.25`), &num)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if num != 0.25 {
		t.Errorf("Decode == %v, want %v", num, 0.25)
	}
}
