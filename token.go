package session

import (
	"crypto/rand"
	"time"
)

const (
	tokenLength    = 256
	expirationTime = time.Duration(time.Second * 300)
)

func genToken() string {
	token := make([]byte, tokenLength)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}
	return string(token)
}
