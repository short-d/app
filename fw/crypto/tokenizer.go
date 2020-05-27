package crypto

type TokenPayload = map[string]interface{}

type Tokenizer interface {
	Encode(payload TokenPayload) (string, error)
	Decode(tokenStr string) (TokenPayload, error)
}
