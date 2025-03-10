package tokenizer

func isWhitespace(c byte) bool {
	return c == 0x09 || c == 0x20 || c == 0x0A || c == 0x0d
}

func isAlpha(c byte) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z')
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || ('0' <= c && c <= '9')
}
