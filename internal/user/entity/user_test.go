package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockPasswordHasher struct {
	mock.Mock
}

func (m *MockPasswordHasher) GenerateFromPassword(password string, cost int) (string, error) {
	args := m.Called(password, cost)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockPasswordHasher) CompareHashAndPassword(hashedPassword string, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

func TestNewUser(t *testing.T) {
	mockHasher := new(MockPasswordHasher)
	mockHasher.On("GenerateFromPassword", mock.Anything, mock.Anything).Return("senha", nil)

	user, err := NewUser("Jon Doe", "jd@jd.com", "senha", mockHasher)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jon Doe", user.Name)
	assert.Equal(t, "jd@jd.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	password := "123456"
	hashedPassword := "askjasdhgb"
	mockHasher := new(MockPasswordHasher)
	mockHasher.On("GenerateFromPassword", password, bcrypt.DefaultCost).Return(hashedPassword, nil)
	mockHasher.On("CompareHashAndPassword", hashedPassword, password).Return(nil)
	mockHasher.On("CompareHashAndPassword", mock.Anything, mock.Anything).Return(errors.New("invalid password"))

	user, err := NewUser("Jon Doe", "jd@jd.com", password, mockHasher)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Password)
	assert.NotEqual(t, password, user.Password)
	assert.True(t, user.ValidatePassword(password))
	assert.False(t, user.ValidatePassword("654321"))
}

func TestNewUser_HashError(t *testing.T) {
	mockHasher := new(MockPasswordHasher)
	mockHasher.On("GenerateFromPassword", mock.Anything, mock.Anything).Return("", errors.New("hash error"))

	user, err := NewUser("Jon Doe", "jd@jd.com", "senha", mockHasher)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}
