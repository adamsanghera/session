package session

import (
	bus "github.com/adamsanghera/redisBus"
)

func (sesh *Session) retrieve(uname string) (string, error) {
	res, err := bus.Client.Get(uname + "_" + sesh.id).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
