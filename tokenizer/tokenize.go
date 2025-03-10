package tokenizer

import (
	"errors"
	"regexp"
	"strconv"
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

func headingIdentifierOf(input []byte) []byte {
	if len(input) == 0 {
		return nil
	}
	if !(isAlpha(input[0]) || input[0] == '_') {
		return nil
	}

	end := 1
	for ; end < len(input); end += 1 {
		if !(isAlphaNumeric(input[end]) || input[end] == '_') {
			break
		}
	}

	return input[:end]
}

func headingStringOf(input []byte) []byte {
	if len(input) == 0 {
		return nil
	}
	if input[0] != '"' {
		return nil
	}

	for i := 1; i < len(input); i += 1 {
		if input[i] == '\\' {
			i += 1
		} else if input[i] == '"' {
			return input[0 : i+1]
		}
	}

	return nil
}

func headingNumberOf(input []byte) *float64 {
	re := regexp.MustCompile(`^-?(0|[1-9]\d*)(\.\d+)?([Ee][+\-]?\d+)?`)
	match := re.Find(input)
	if match == nil {
		return nil
	}

	res, err := strconv.ParseFloat(string(input[:len(match)]), 64)
	if err != nil {
		panic(err)
	}
	return &res
}

func unescapeString(str []byte) []byte {
	result := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			if i+1 >= len(str) {
				return nil
			}
			switch str[i+1] {
			case '"':
				result = append(result, '"')
			case '\\':
				result = append(result, '\\')
			case '/':
				result = append(result, '/')
			case 'b':
				result = append(result, '\b')
			case 'f':
				result = append(result, '\f')
			case 'n':
				result = append(result, '\n')
			case 'r':
				result = append(result, '\r')
			case 't':
				result = append(result, '\t')
			case 'u':
				panic("Unicode literal is not supported")
			}
			i += 1
		} else {
			result = append(result, str[i])
		}
	}

	return result
}

func NextTokenOf(input []byte) (*Token, error) {
	start := 0
	for ; start < len(input); start += 1 {
		if !isWhitespace(input[start]) {
			break
		}
	}

	if start >= len(input) {
		return &Token{kind: EOF}, nil
	}

	input = input[start:]

	switch input[0] {
	case '[':
		return &Token{kind: BeginArray}, nil
	case ']':
		return &Token{kind: EndArray}, nil
	case '{':
		return &Token{kind: BeginObject}, nil
	case '}':
		return &Token{kind: EndObject}, nil
	case ':':
		return &Token{kind: NameSeparator}, nil
	case ',':
		return &Token{kind: ValueSeparator}, nil
	case '"':
		bytes := headingStringOf(input)
		if bytes == nil {
			return nil, errors.New("Unexpected EOF")
		}
		str := string(unescapeString((bytes[1 : len(bytes)-1])))
		return &Token{kind: String, value: str}, nil
	}

	if identifier := headingIdentifierOf(input); identifier != nil {
		switch string(identifier) {
		case "null":
			return &Token{kind: Null}, nil
		case "true":
			return &Token{kind: Boolean, value: true}, nil
		case "false":
			return &Token{kind: Boolean, value: false}, nil
		}
	}

	number := headingNumberOf(input)
	if number != nil {
		return &Token{kind: Number, value: *number}, nil
	}

	return nil, errors.New("Unexpected character")
}
