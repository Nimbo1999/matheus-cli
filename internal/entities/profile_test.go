package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	profile, err := GetProfile("unexistant")
	assert.Nil(t, profile)
	assert.Error(t, err, ErrProfileNotFound)
}
