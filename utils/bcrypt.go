package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func BcryptEndoderPassword(password []byte) string {
	password_hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(password_hash)
}

func CompareHashAndPassword(password []byte, password_compare []byte) bool {
	err := bcrypt.CompareHashAndPassword(password, password_compare)

	return err == nil
}
