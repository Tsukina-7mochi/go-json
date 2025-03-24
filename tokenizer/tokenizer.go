package tokenizer

import (
	"errors"
	"regexp"
	"strconv"

	. "json/strings"
)

var ErrUnexpectedEOF = errors.New("Unexpected EOF")
var ErrUnexpectedCharacter = errors.New("Unexpected character")

type Tokenizer struct {
	input []byte
	index int
}

func NewTokenizer(input []byte) *Tokenizer {
	return &Tokenizer{input: input, index: 0}
}

func (t *Tokenizer) skipWhitespace() {
	for ; t.index < len(t.input); t.index += 1 {
		if !isWhitespace(t.input[t.index]) {
			break
		}
	}
}

func (t *Tokenizer) matchesHeadingIdentifier(identifier []byte) bool {
	if t.index+len(identifier) > len(t.input) {
		return false
	}

	for i := 0; i < len(identifier); i += 1 {
		if t.input[t.index+i] != identifier[i] {
			return false
		}
	}

	if t.index+len(identifier) < len(t.input) {
		if isAlphaNumeric(t.input[t.index+len(identifier)]) || t.input[t.index+len(identifier)] == '_' {
			return false
		}
	}

	return true
}

func (t *Tokenizer) takeString() (*Token, error) {
	if t.index >= len(t.input) {
		return nil, nil
	}
	if t.input[t.index] != '"' {
		return nil, nil
	}

	start := t.index + 1

	for t.index += 1; t.index < len(t.input); t.index += 1 {
		if t.input[t.index] == '\\' {
			t.index += 1
		} else if t.input[t.index] == '"' {
			t.index += 1

			value, err := UnescapeString(t.input[start : t.index-1])
			if err != nil {
				return nil, err
			}

			return &Token{Kind: StringToken, Value: string(value)}, nil
		}
	}

	return nil, ErrUnexpectedEOF
}

func (t *Tokenizer) takeNumber() *Token {
	re := regexp.MustCompile(`^-?(0|[1-9]\d*)(\.\d+)?([Ee][+\-]?\d+)?`)
	match := re.Find(t.input[t.index:])
	if match == nil {
		return nil
	}

	matchBytes := t.input[t.index : t.index+len(match)]
	t.index += len(match)

	value, err := strconv.ParseFloat(string(matchBytes), 64)
	if err != nil {
		panic(err)
	}
	return &Token{Kind: NumberToken, Value: value}
}

func (t *Tokenizer) Next() (*Token, error) {
	t.skipWhitespace()

	if t.index >= len(t.input) {
		return &Token{Kind: EOFToken}, nil
	}

	switch t.input[t.index] {
	case '[':
		t.index += 1
		return &Token{Kind: BeginArrayToken}, nil
	case ']':
		t.index += 1
		return &Token{Kind: EndArrayToken}, nil
	case '{':
		t.index += 1
		return &Token{Kind: BeginObjectToken}, nil
	case '}':
		t.index += 1
		return &Token{Kind: EndObjectToken}, nil
	case ':':
		t.index += 1
		return &Token{Kind: NameSeparatorToken}, nil
	case ',':
		t.index += 1
		return &Token{Kind: ValueSeparatorToken}, nil
	}

	if t.matchesHeadingIdentifier([]byte("null")) {
		t.index += 4
		return &Token{Kind: NullToken}, nil
	}
	if t.matchesHeadingIdentifier([]byte("true")) {
		t.index += 4
		return &Token{Kind: BooleanToken, Value: true}, nil
	}
	if t.matchesHeadingIdentifier([]byte("false")) {
		t.index += 5
		return &Token{Kind: BooleanToken, Value: false}, nil
	}

	stringToken, err := t.takeString()
	if err != nil {
		return nil, err
	}
	if stringToken != nil {
		return stringToken, nil
	}

	numberToken := t.takeNumber()
	if numberToken != nil {
		return numberToken, nil
	}

	return nil, ErrUnexpectedCharacter
}
