package entity

import (
	"testing"
	"time"

	"github.com/Solon97/goexpert-api/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewOrganizzeAccount(t *testing.T) {
	isoDate := entity.ISODate(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	organizzeAccount, err := NewOrganizzeAccount("Solon", "solonb20@gmail.com", "c29sb25iMjBAZ21haWwuY29tOjQyYTRkZjM4MDVkNmM3MmZjMDA3YjcyZDVkMjBmZjU1NTI4ZjU4OTA=", isoDate)
	assert.Nil(t, err)
	assert.NotNil(t, organizzeAccount)
	assert.NotEmpty(t, organizzeAccount.ID)
	assert.Equal(t, "Solon", organizzeAccount.Username)
	assert.Equal(t, "solonb20@gmail.com", organizzeAccount.Email)
	assert.Equal(t, "c29sb25iMjBAZ21haWwuY29tOjQyYTRkZjM4MDVkNmM3MmZjMDA3YjcyZDVkMjBmZjU1NTI4ZjU4OTA=", organizzeAccount.BasicToken)
	assert.Equal(t, isoDate.String(), organizzeAccount.InitialDate.String())
}

func TestInvalidCredentials(t *testing.T) {
	isoDate := entity.ISODate(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	_, err := NewOrganizzeAccount("Solon", "solonb20@gmail.com", "123", isoDate)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
}
