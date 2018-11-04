package model

import (
	"github.com/realityone/vericry/pkg/bcrypt"
)

// Salt is
func (t *Token) Salt() (string, error) {
	hashb, err := bcrypt.Base64Decode([]byte(t.HashData))
	if err != nil {
		return "", err
	}

	hashed, err := bcrypt.NewFromHash(hashb)
	if err != nil {
		return "", err
	}

	return string(hashed.Salt()), nil
}
