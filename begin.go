package session

import (
	"time"

	bus "github.com/adamsanghera/redisBus"
)

// Begin allows one to instantiate a new session, given a username.
// If a valid session instance already exists for the given user, no action is taken.
func (sesh *Session) Begin(uname string) (string, time.Duration, error) {
	// Make a new token!
	token := sesh.genToken()

	// Try to Set the key
	result, err := bus.Client.SetNX(uname+"_"+sesh.id, token, sesh.timeToLive).Result()
	if err != nil {
		return "", 0, err
	}

	// If unset, then the user already has a valid token.
	//   While it might be reasonable to renew the token automatically here,
	//   such an operation is beyond the scope of this function.
	if result == false {
		return "", time.Duration(0), ErrTokenExists
	}

	// If the key is set, then the user has just been logged in!
	return token, expirationTime, nil
}
