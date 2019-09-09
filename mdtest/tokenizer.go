package mdtest

import (
	"encoding/json"

	"github.com/byliuyang/app/fw"
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
