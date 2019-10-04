package internal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNil(t *testing.T) {
	// given
	eid := "20191004:135407"
	var value error = nil

	// when
	CheckNil(eid, value)
}

func TestCheckNilOnValue(t *testing.T) {
	// given
	eid := "20191004:140503"
	var value error = errors.New("Some error")
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	// when
	CheckNil(eid, value)
}
