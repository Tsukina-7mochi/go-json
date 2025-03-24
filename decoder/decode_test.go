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

func TestDecode_BooleanArray(t *testing.T) {
	want := []bool{true, false, true}
	var got []bool
	err := Decode([]byte(`[true, false, true]`), &got)

	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("Decode == %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Decode == %v, want %v", got, want)
		}
	}
}

func TestDecode_StringArray(t *testing.T) {
	want := []string{"hello", "world"}
	var got []string
	err := Decode([]byte(`["hello", "world"]`), &got)

	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("Decode == %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Decode == %v, want %v", got, want)
		}
	}
}

func TestDecode_IntArray(t *testing.T) {
	want := []int{1, 2, 3}
	var got []int
	err := Decode([]byte(`[1, 2, 3]`), &got)

	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("Decode == %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Decode == %v, want %v", got, want)
		}
	}
}

func TestDecode_Float64Array(t *testing.T) {
	want := []float64{0.1, 0.2, 0.3}
	var got []float64
	err := Decode([]byte(`[0.1, 0.2, 0.3]`), &got)

	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("Decode == %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Decode == %v, want %v", got, want)
		}
	}
}

func TestDecode_IntArrayArray(t *testing.T) {
	want := [][]int{{1, 2}, {3, 4}}
	var got [][]int
	err := Decode([]byte(`[[1, 2], [3, 4]]`), &got)

	if err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("Decode == %v, want %v", got, want)
	}
	for i := range got {
		if len(got[i]) != len(want[i]) {
			t.Fatalf("Decode == %v, want %v", got, want)
		}
		for j := range got[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("Decode == %v, want %v", got, want)
			}
		}
	}
}
