package userentity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Solon", "solon@teste.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Solon", user.Name)
	assert.Equal(t, "solon@teste.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Solon", "solon@teste.com", "123456")
	assert.Nil(t, err)
	assert.NotEqual(t, user.Password, "123456")
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}