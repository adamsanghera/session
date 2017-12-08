package session

import (
	"crypto/rand"
	"time"
)

const (
	tokenLength    = 256
	expirationTime = time.Duration(time.Second * 300)
)

//Generates a new token for the session.
func (sesh *Session) genToken() string {
	token := make([]byte, sesh.tokenLength)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}
	return string(token)
}
