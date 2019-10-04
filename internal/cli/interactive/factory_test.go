package interactive

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wavesoftware/serverless-installer/internal/domain/contract"
)

func TestCreate(t *testing.T) {
	// given
	answers := contract.LocationAwareAnswers{}

	// when
	screen := Create(&answers)

	// then
	assert.NotNil(t, screen)
}
