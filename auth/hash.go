package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// CompareHashAndPassword - Compare Hash and Password
// Use bcrypt to compare the hashed password with the plain text password.
func CompareHashAndPassword(dbPw string, reqPw string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(dbPw),
		[]byte(reqPw),
	)
}

// HashPassword - Hash Password
// Use bcrypt to hash the password with a cost of 10.
// The higher the cost, the more secure the hash, but it will also take longer to compute.
func HashPassword(pw string) (string, error) {
	hashPW, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		return "", err
	}
	return string(hashPW), nil
}
