package session

func (sesh *Session) Validate(challenge string, uname string) (bool, error) {
	answer, err := sesh.retrieve(uname)
	if err != nil {
		return false, err
	}
	return challenge == answer, nil
}
