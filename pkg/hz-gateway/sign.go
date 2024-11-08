package sign

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

type SignType string

var SM3 SignType = "SM3"

type IAuth interface {
	GetSignType() SignType
	Verify(sign string, data map[string]interface{}) bool
	Sign(data map[string]interface{}) string
}

func NewSignProvider(signType string, key string) (IAuth, error) {
	switch signType {
	case string(SM3):
		return &authSM3{key}, nil
	default:
		return nil, errors.New("unsupported sign type")
	}
}

type authSM3 struct {
	Key string
}

func (a *authSM3) GetSignType() SignType {
	return SM3
}

func (a *authSM3) Verify(sign string, data map[string]interface{}) bool {
	return true
}

func (a *authSM3) Sign(data map[string]interface{}) string {
	result := make([]byte, 10)
	rand.Read(result) //nolint:errcheck
	return hex.EncodeToString(result)
}