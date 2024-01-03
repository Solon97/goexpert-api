package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/Solon97/goexpert-api/pkg/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountValidator struct {
	mock.Mock
}

func (m *MockAccountValidator) IsValidAccount(username, email, token string) (bool, error) {
	args := m.Called(username, email, token)
	return args.Bool(0), args.Error(1)
}

func TestNewOrganizzeAccount(t *testing.T) {
	validator := new(MockAccountValidator)
	validator.On("IsValidAccount", "testUser", "test@example.com", "token123").Return(true, nil)

	isoDate := entity.ISODate(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	organizzeAccount, err := NewOrganizzeAccount("testUser", "test@example.com", "token123", isoDate, validator)
	assert.Nil(t, err)
	assert.NotNil(t, organizzeAccount)
	assert.NotEmpty(t, organizzeAccount.ID)
	assert.Equal(t, "testUser", organizzeAccount.Username)
	assert.Equal(t, "test@example.com", organizzeAccount.Email)
	assert.Equal(t, "token123", organizzeAccount.BasicToken)
	assert.Equal(t, isoDate.String(), organizzeAccount.InitialDate.String())
}

func TestInvalidCredentials(t *testing.T) {
	validator := new(MockAccountValidator)
	validator.On("IsValidAccount", "testUser", "test@example.com", "token123").Return(false, nil)
	isoDate := entity.ISODate(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	_, err := NewOrganizzeAccount("testUser", "test@example.com", "token123", isoDate, validator)
	assert.NotNil(t, err)
	assert.Error(t, err, "invalid credentials")
}

func TestIsValidAccountError(t *testing.T) {
	validator := new(MockAccountValidator)
	validator.On("IsValidAccount", "testUser", "test@example.com", "token123").Return(false, fmt.Errorf("some error"))
	isoDate := entity.ISODate(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	_, err := NewOrganizzeAccount("testUser", "test@example.com", "token123", isoDate, validator)
	assert.NotNil(t, err)
	assert.Error(t, err, "account validation error: some error")
}
