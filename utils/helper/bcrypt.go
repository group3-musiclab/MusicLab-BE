package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(inputPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 14)
	return string(hashedPassword), err
}

func CompareHashPassword(inputPassword, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword))
	return err == nil
}
