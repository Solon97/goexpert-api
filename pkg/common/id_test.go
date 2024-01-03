package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := NewID()
	assert.NotEmpty(t, id)
	assert.NotEqual(t, id, ID{})
	assert.Equal(t, 36, len(id.String()))
}

func TestID_Parse(t *testing.T) {
	id := NewID()
	parsedID, err := ParseID(id.String())
	assert.Nil(t, err)
	assert.Equal(t, id, parsedID)
}
