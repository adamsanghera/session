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

//Create ...
// Creates a new session for a given user, if one does not already exist.
// If a valid session already exists, no action is taken.
func Create(uname string) (string, time.Duration, error) {
	// Make a new token!
	token := genToken(tokenLength)

	// Try to Set the key
	result, err := bus.Client.SetNX(uname, token, expirationTime).Result()
	if err != nil {
		// We really don't
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
