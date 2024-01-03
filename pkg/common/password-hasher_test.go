package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestBcryptPasswordHasher(t *testing.T) {
	hasher := BcryptPasswordHasher{}
	password := "123456"
	hash, err := hasher.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.Nil(t, err)
	assert.NotEmpty(t, hash)
}

func TestBcryptPasswordHasher_CompareHashAndPassword(t *testing.T) {
	hasher := BcryptPasswordHasher{}
	password := "123456"
	hash, _ := hasher.GenerateFromPassword(password, bcrypt.DefaultCost)
	err := hasher.CompareHashAndPassword(hash, password)
	assert.Nil(t, err)
}

func TestBcryptPasswordHasher_CompareHashAndPassword_Error(t *testing.T) {
	hasher := BcryptPasswordHasher{}
	password := "123456"
	hash, _ := hasher.GenerateFromPassword(password, bcrypt.DefaultCost)
	err := hasher.CompareHashAndPassword(hash, "654321")
	assert.NotNil(t, err)
}
