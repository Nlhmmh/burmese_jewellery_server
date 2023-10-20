package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(dbPw string, reqPw string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(dbPw),
		[]byte(reqPw),
	)
}

func HashPassword(pw string) (string, error) {
	hashPW, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		return "", err
	}
	return string(hashPW), nil
}
