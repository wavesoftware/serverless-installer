package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAnswersSaver(t *testing.T) {
	// when
	saver := CreateAnswersSaver()

	// then
	assert.NotNil(t, saver)
}
