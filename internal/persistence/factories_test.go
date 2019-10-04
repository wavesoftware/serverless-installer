package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFileWriter(t *testing.T) {
	// when
	writer := CreateFileWriter()

	// then
	assert.NotNil(t, writer)
}
