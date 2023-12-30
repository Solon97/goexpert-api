package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ISODateTest struct {
	Date ISODate `json:"date"`
}

func TestISODate_UnmarshalJSON(t *testing.T) {
	date := ISODateTest{}
	err := date.Date.UnmarshalJSON([]byte(`"2022-01-01"`))
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "2022-01-01", date.Date.String())
}
