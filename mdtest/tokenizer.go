package mdtest

import (
	"encoding/json"

	"github.com/short-d/app/fw"
)

var _ fw.CryptoTokenizer = (*CryptoTokenizerFake)(nil)

type CryptoTokenizerFake struct {
}

func (c CryptoTokenizerFake) Encode(payload fw.TokenPayload) (string, error) {
	buf, err := json.Marshal(payload)
	return string(buf), err
}

func (c CryptoTokenizerFake) Decode(tokenStr string) (fw.TokenPayload, error) {
	payload := map[string]interface{}{}
	err := json.Unmarshal([]byte(tokenStr), &payload)
	return payload, err
}

func NewCryptoTokenizerFake() CryptoTokenizerFake {
	return CryptoTokenizerFake{}
}
