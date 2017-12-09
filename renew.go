package session

import "time"

func (sesh *Session) Renew(uname string, challenge string) (string, time.Duration, error) {
	valid, err := sesh.Validate(challenge, uname)
	if err != nil {
		return "", 0, err
	}
	if valid {
		// We have to revoke, because begin will return an error if the key already exists.
		_, err := sesh.Revoke(uname)
		if err != nil {
			return "", 0, err
		}
		return sesh.Begin(uname)
	}
	return "", 0, err
}
