package session

import (
	"github.com/adamsanghera/hashing"
	bus "github.com/adamsanghera/redisBus"
)

// retrieve returns the tokens indexed by the username/id
func (sesh *Session) retrieve(uname string) (string, error) {
	res, err := bus.Client.Get(uname + "_" + sesh.id).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

// secure retrieve simply retrieves by hashed key
func (sesh *SecureSession) retrieve(uname string) (string, error) {
	key := hashing.WithoutSalt(uname + "_" + sesh.id)

	res, err := bus.Client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
