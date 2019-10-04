package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicRunnerFactory(t *testing.T) {
	// when
	runner := NewRunner()

	// then
	assert.NotNil(t, runner)
}

func TestHelpRun(t *testing.T) {
	// given
	args := []string{"cmd", "--help"}
	runner := NewRunnerWithArgs(args)

	// when
	retcode := runner.Run()

	// then
	assert.Equal(t, 0, retcode)
}

func TestInvalidRun(t *testing.T) {
	// given
	args := []string{"cmd", "invalid_command"}
	runner := NewRunnerWithArgs(args)

	// when
	retcode := runner.Run()

	// then
	assert.Equal(t, 1, retcode)
}
