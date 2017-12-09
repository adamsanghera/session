package session

import (
	"time"
)

const (
	defaultTokenLength    = 256
	defaultExpirationTime = time.Duration(time.Second * 300)
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
	id          string
}

// SecureSession is the same as Session except that it stores
// hashes of usernames and ids in redis, instead of plaintext
type SecureSession struct {
	tokenLength int
	timeToLive  time.Duration
	id          string
}

//NewSession returns an instance of the Session struct.
// It's up to the user to make sure that they maintain unique ids.
func NewSession(tokenLength int, timeToLive time.Duration, id string) *Session {
	return &Session{
		tokenLength: tokenLength,
		timeToLive:  timeToLive,
		id:          id,
	}
}

func NewBasicSession() *Session {
	return &Session{
		tokenLength: defaultTokenLength,
		timeToLive:  defaultExpirationTime,
		id:          "",
	}
}
