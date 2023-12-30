package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jon Doe", "jd@jd.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jon Doe", user.Name)
	assert.Equal(t, "jd@jd.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	password := "123456"
	user, err := NewUser("Jon Doe", "jd@jd.com", password)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Password)
	assert.NotEqual(t, password, user.Password)
	assert.True(t, user.ValidatePassword(password))
	assert.False(t, user.ValidatePassword("654321"))
}
