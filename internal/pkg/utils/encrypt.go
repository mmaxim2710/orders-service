package utils

import "golang.org/x/crypto/bcrypt"

func EncryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ComparePassword(hashPassword string, inputPassword string) bool {
	byteHash := []byte(hashPassword)
	byteInput := []byte(inputPassword)

	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}
