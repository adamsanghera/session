package session

// Validate is used to verify that a provided token for a given user matches the stored token
func (sesh *Session) Validate(challenge string, uname string) (bool, error) {
	answer, err := sesh.retrieve(uname)
	if err != nil {
		return false, err
	}
	return challenge == answer, nil
}

// Validate for SecureSession is the same deal
func (sesh *SecureSession) Validate(challenge string, uname string) (bool, error) {
	answer, err := sesh.retrieve(uname)
	if err != nil {
		return false, err
	}
	return challenge == answer, nil
}
