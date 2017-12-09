package session

import (
	"github.com/adamsanghera/hashing"
	bus "github.com/adamsanghera/redisBus"
)

// Revoke gives you the ability to revoke a given session token.
// returns true iff (1) privilege was revoked (2) privilege was never assigned.
// returns false iff there's some deeper error with redis.
func (sesh *Session) Revoke(uname string) (bool, error) {
	// If res is 1, it means that it was deleted. 0 means not-found.
	// err will only be non-nil if there's some deeper issue (lost connection, etc.)
	_, err := bus.Client.Del(uname + "_" + sesh.id).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Revoke for SecureSession is the same, just uses a hashed key.
func (sesh *SecureSession) Revoke(uname string) (bool, error) {
	key := hashing.WithoutSalt(uname + "_" + sesh.id)

	_, err := bus.Client.Del(key).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
