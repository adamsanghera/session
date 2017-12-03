package session

import "errors"

//ErrTokenExists ...
// This error is created when a call made to Create
//  terminates by identifying that an existing
//  session token already exists.
var ErrTokenExists = errors.New("A session token already exists for this user")
