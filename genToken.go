package session

import (
	"crypto/rand"
	"time"
)

const (
	tokenLength    = 256
	expirationTime = time.Duration(time.Second * 300)
)

func genToken(size int) string {
	token := make([]byte, size)
	_, err := rand.Read(token)
	if err != nil {
		// What to do?
		// 	 panic(err)
		//   return err
	}
	return string(token)
}
