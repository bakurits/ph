package ph

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", errors.Wrap(err, "error while hashing")
	}
	return string(hash), nil
}

func Compare(hashed string, plain string) bool {
	byteHash := []byte(hashed)
	return bcrypt.CompareHashAndPassword(byteHash, []byte(plain)) == nil
}
