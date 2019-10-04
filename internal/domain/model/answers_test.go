package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T) {
	// when
	answers := Answers{}

	// then
	assert.NotNil(t, answers)
}
