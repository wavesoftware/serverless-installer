package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// given
	oldExit := exit
	oldArgs := os.Args
	defer func() {
		exit = oldExit
		os.Args = oldArgs
	}()
	exitcode := ^int(0)
	exit = func(received int) {
		exitcode = received
	}
	os.Args = []string{"cmd", "--help"}

	// when
	main()

	assert.Equal(t, 0, exitcode)
}

func TestExit(t *testing.T) {
	exit(0)
}
