package credentials

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func GenerateBCrypt(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate bcrypt hash")
	}

	return string(b), nil
}

func CompareBCrypt(password string, digest string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(password))
	switch err {
	case nil:
		return true, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return false, nil
	default:
		return false, err
	}
}
