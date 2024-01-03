package entity

import (
	"github.com/Solon97/goexpert-api/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	GenerateFromPassword(password string, cost int) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
}

type User struct {
	ID       common.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	hasher   PasswordHasher
}

func NewUser(name, email, password string, hasher PasswordHasher) (*User, error) {
	hash, err := hasher.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       common.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
		hasher:   hasher,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := u.hasher.CompareHashAndPassword(u.Password, password)
	return err == nil
}
