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
