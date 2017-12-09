package session

import (
	"crypto/rand"
	"time"
)

const (
	tokenLength    = 256
	expirationTime = time.Duration(time.Second * 300)
)

func genToken(length int) string {
	token := make([]byte, length)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}
	return string(token)
}

//Generates a new token for the session.
func (sesh *Session) genToken() string {
	return genToken(sesh.tokenLength)
}

func (sesh *SecureSession) genToken() string {
	return genToken(sesh.tokenLength)
}
