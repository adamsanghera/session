package session

import (
	"time"

	bus "github.com/adamsanghera/redisBus"
)

/*
Examination of possible states
	There are the following possible session states:
		1. User is not in the database
		2. User is in database
			2.1 User has no associated token
			2.2 User has an associated token
				2.2.1 Token is expired
				2.2.2 Token is still valid

	Thanks to Redis' support for self-expiring keys, 2.1 and 2.2.1 are impossible states.
	Thus, this function need only consider 1 and 2.2.2 as possible states.

	If 1,
		Create a new KVP, with fresh token
	If 2.2.2,
		Do nothing.
*/

// Session is a type that allows the user to configure their own kinds of token.
// tokenLength stipulates the length of the token.
// timeToLive stipulates the amount of time that a session's token is valid.
// id allows users to create different kinds of sessions.
//   for example, one can make a token for free users, and a token for paid users.
type Session struct {
	tokenLength int
	timeToLive  time.Duration
	id          int
}

// Returns an instance of the Session struct, which allows users to create tokens
// and validate their own sessions.
func NewSession(tokenLength int, timeToLive time.Duration, id int) *Session {
	return &Session{
		tokenLength: tokenLength,
		timeToLive:  timeToLive,
		id:          id,
	}
}

// CreateToken allows one to instantiate a new session, given a username.
// If a valid session instance already exists for the given user, no action is taken.
func (*Session) CreateToken(uname string) (string, time.Duration, error) {
	// Make a new token!
	token := genToken()

	// Try to Set the key
	result, err := bus.Client.SetNX(uname, token, expirationTime).Result()
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
