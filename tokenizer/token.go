package tokenizer

type TokenKind string

const (
	EOFToken            = TokenKind("EOF")
	BeginArrayToken     = TokenKind("BeginArray")
	EndArrayToken       = TokenKind("EndArray")
	BeginObjectToken    = TokenKind("BeginObject")
	EndObjectToken      = TokenKind("EndObject")
	NameSeparatorToken  = TokenKind("NameSeparator")
	ValueSeparatorToken = TokenKind("ValueSeparator")
	NullToken           = TokenKind("Null")
	BooleanToken        = TokenKind("Boolean")
	NumberToken         = TokenKind("Number")
	StringToken         = TokenKind("String")
)

type Token struct {
	Kind  TokenKind
	Value interface{}
}

func (t *Token) BoolValue() bool {
	return t.Value.(bool)
}

func (t *Token) FloatValue() float64 {
	return t.Value.(float64)
}

func (t *Token) StringValue() string {
	return t.Value.(string)
}
