package main

import (
	"os"

	"github.com/wavesoftware/serverless-installer/internal/cli"
)

func main() {
	os.Exit(cli.NewRunner().Run())
}
