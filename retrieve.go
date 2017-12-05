package session

import (
	bus "github.com/adamsanghera/redisBus"
)

func retrieve(uname string) (string, error) {
	res, err := bus.Client.Get(uname).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
