package session

import "time"

func (sesh *Session) Renew(uname string, challenge string) (string, time.Duration, error) {
	valid, err := sesh.Validate(challenge, uname)
	if err != nil {
		return "", 0, err
	}
	if valid {
		return sesh.CreateToken(uname)
	}
	return "", 0, err
}
