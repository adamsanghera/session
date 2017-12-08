package session

import (
	bus "github.com/adamsanghera/redisBus"
)

// Revoke gives you the ability to revoke a given session token.
func (sesh *Session) Revoke(uname string) (bool, error) {
	// If res is 1, it means that it was deleted. 0 means not-found.
	// err will only be non-nil if there's some deeper issue (lost connection, etc.)
	_, err := bus.Client.Del(uname + "_" + sesh.id).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
