package decoder

import "testing"

func TestDecode_BooleanTrue(t *testing.T) {
	want := true
	var got bool
	err := Decode([]byte(`true`), &got)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if got != want {
		t.Errorf("Decode == %v, want %v", got, want)
	}
}

func TestDecode_BooleanFalse(t *testing.T) {
	want := false
	var got bool
	err := Decode([]byte(`false`), &got)
	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if got != want {
		t.Errorf("Decode == %v, want %v", got, want)
	}
}

func TestDecode_String(t *testing.T) {
	want := "hello"
	var got string
	err := Decode([]byte(`"hello"`), &got)

	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if got != want {
		t.Errorf("Decode == %q, want %q", got, want)
	}
}

func TestDecode_Int(t *testing.T) {
	want := 123
	var num int
	err := Decode([]byte(`123`), &num)

	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if num != want {
		t.Errorf("Decode == %v, want %v", num, want)
	}
}

func TestDecode_Float64(t *testing.T) {
	want := 0.25
	var num float64
	err := Decode([]byte(`0.25`), &num)

	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if num != want {
		t.Errorf("Decode == %v, want %v", num, want)
	}
}
