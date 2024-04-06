package util

import (
	bycrpt "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bycrpt.GenerateFromPassword([]byte(password), bycrpt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bycrpt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
