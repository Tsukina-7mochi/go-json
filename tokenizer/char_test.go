package tokenizer

import "testing"

func TestIsWhitespace_HT(t *testing.T) {
	if !isWhitespace('\t') {
		t.Errorf("isWhitespace('\t') == false, want true")
	}
}

func TestIsWhitespace_SP(t *testing.T) {
	if !isWhitespace(' ') {
		t.Errorf("isWhitespace(' ') == false, want true")
	}
}

func TestIsWhitespace_LF(t *testing.T) {
	if !isWhitespace('\n') {
		t.Errorf("isWhitespace('\n') == false, want true")
	}
}

func TestIsWhitespace_CR(t *testing.T) {
	if !isWhitespace('\r') {
		t.Errorf("isWhitespace('\r') == false, want true")
	}
}

func TestIsAlpha_a(t *testing.T) {
	if !isAlpha('a') {
		t.Errorf("isAlpha('a') == false, want true")
	}
}

func TestIsAlpha_z(t *testing.T) {
	if !isAlpha('z') {
		t.Errorf("isAlpha('z') == false, want true")
	}
}

func TestIsAlpha_A(t *testing.T) {
	if !isAlpha('A') {
		t.Errorf("isAlpha('A') == false, want true")
	}
}

func TestIsAlpha_Z(t *testing.T) {
	if !isAlpha('Z') {
		t.Errorf("isAlpha('Z') == false, want true")
	}
}

func TestIsAlphaNumeric_0(t *testing.T) {
	if !isAlphaNumeric('0') {
		t.Errorf("isAlphaNumeric('0') == false, want true")
	}
}

func TestIsAlphaNumeric_9(t *testing.T) {
	if !isAlphaNumeric('9') {
		t.Errorf("isAlphaNumeric('9') == false, want true")
	}
}

func TestIsAlphaNumeric_a(t *testing.T) {
	if !isAlphaNumeric('a') {
		t.Errorf("isAlphaNumeric('a') == false, want true")
	}
}

func TestIsAlphaNumeric_z(t *testing.T) {
	if !isAlphaNumeric('z') {
		t.Errorf("isAlphaNumeric('z') == false, want true")
	}
}

func TestIsAlphaNumeric_A(t *testing.T) {
	if !isAlphaNumeric('A') {
		t.Errorf("isAlphaNumeric('A') == false, want true")
	}
}

func TestIsAlphaNumeric_Z(t *testing.T) {
	if !isAlphaNumeric('Z') {
		t.Errorf("isAlphaNumeric('Z') == false, want true")
	}
}
