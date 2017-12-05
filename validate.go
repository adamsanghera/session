package session

func Validate(challenge string, uname string) (bool, error) {
	answer, err := retrieve(uname)
	if err != nil {
		return false, err
	}
	return challenge == answer, nil
}
