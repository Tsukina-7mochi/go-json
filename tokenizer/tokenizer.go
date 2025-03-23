package tokenizer

import (
	"errors"
	"regexp"
	"strconv"

	. "json/strings"
)

type TokenKind string

const (
	EOF            = TokenKind("EOF")
	BeginArray     = TokenKind("BeginArray")
	EndArray       = TokenKind("EndArray")
	BeginObject    = TokenKind("BeginObject")
	EndObject      = TokenKind("EndObject")
	NameSeparator  = TokenKind("NameSeparator")
	ValueSeparator = TokenKind("ValueSeparator")
	Null           = TokenKind("Null")
	Boolean        = TokenKind("Boolean")
	Number         = TokenKind("Number")
	String         = TokenKind("String")
)

type Token struct {
	kind  TokenKind
	value interface{}
}

func (t *Token) BoolValue() bool {
	return t.value.(bool)
}

func (t *Token) FloatValue() float64 {
	return t.value.(float64)
}

func (t *Token) StringValue() string {
	return t.value.(string)
}

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

			return &Token{kind: String, value: string(value)}, nil
		}
	}

	return nil, errors.New("Unexpected EOF")
}

func (t *Tokenizer) takeNumber() *Token {
	re := regexp.MustCompile(`^-?(0|[1-9]\d*)(\.\d+)?([Ee][+\-]?\d+)?`)
	match := re.Find(t.input[t.index:])
	if match == nil {
		return nil
	}

	matchBytes := t.input[t.index : t.index+len(match)]
	value, err := strconv.ParseFloat(string(matchBytes), 64)
	if err != nil {
		panic(err)
	}
	return &Token{kind: Number, value: value}
}

func (t *Tokenizer) Next() (*Token, error) {
	t.skipWhitespace()

	if t.index >= len(t.input) {
		return &Token{kind: EOF}, nil
	}

	switch t.input[t.index] {
	case '[':
		t.index += 1
		return &Token{kind: BeginArray}, nil
	case ']':
		t.index += 1
		return &Token{kind: EndArray}, nil
	case '{':
		t.index += 1
		return &Token{kind: BeginObject}, nil
	case '}':
		t.index += 1
		return &Token{kind: EndObject}, nil
	case ':':
		t.index += 1
		return &Token{kind: NameSeparator}, nil
	case ',':
		t.index += 1
		return &Token{kind: ValueSeparator}, nil
	}

	if t.matchesHeadingIdentifier([]byte("null")) {
		t.index += 4
		return &Token{kind: Null}, nil
	}
	if t.matchesHeadingIdentifier([]byte("true")) {
		t.index += 4
		return &Token{kind: Boolean, value: true}, nil
	}
	if t.matchesHeadingIdentifier([]byte("false")) {
		t.index += 5
		return &Token{kind: Boolean, value: false}, nil
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

	return nil, errors.New("Unexpected character")
}
