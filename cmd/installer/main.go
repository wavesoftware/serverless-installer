package main

import (
	"os"

	"github.com/wavesoftware/serverless-installer/internal/cli"
)

var exit func(int) = func(exitcode int) {
	os.Exit(exitcode)
}

func main() {
	exit(cli.NewRunner().Run())
}
