package common

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct{}

func (BcryptPasswordHasher) GenerateFromPassword(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash), err
}

func (BcryptPasswordHasher) CompareHashAndPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
