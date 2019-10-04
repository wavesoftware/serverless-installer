package contract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T) {
	// when
	answers := LocationAwareAnswers{}

	// then
	assert.NotNil(t, answers)
}

func TestFilename(t *testing.T) {
	// then
	assert.NotEmpty(t, AnswersFilename)
}
