package crypto

import (
	"encoding/json"
)

var _ Tokenizer = (*TokenizerFake)(nil)

type TokenizerFake struct {
}

func (t TokenizerFake) Encode(payload TokenPayload) (string, error) {
	buf, err := json.Marshal(payload)
	return string(buf), err
}

func (t TokenizerFake) Decode(tokenStr string) (TokenPayload, error) {
	payload := map[string]interface{}{}
	err := json.Unmarshal([]byte(tokenStr), &payload)
	return payload, err
}

func NewTokenizerFake() TokenizerFake {
	return TokenizerFake{}
}
