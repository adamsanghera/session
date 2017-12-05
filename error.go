package session

import (
	"errors"
)

//ErrTokenExists ...
// This error is created when a call made to Create
//  terminates by identifying that an existing
//  session token already exists.
var ErrTokenExists = errors.New("A session token already exists for this user")

// catchError is meant to be used as a deferred function to turn a panic(sessionError) into a
// plain error. It overwrites the error return of the function that deferred its call.
func catchError(err *error) {
	if recoveredErr := recover(); recoveredErr != nil {
		s, ok := recoveredErr.(string)
		if !ok {
			panic(s)
		}
		*err = errors.New(s)
	}
	return
}
