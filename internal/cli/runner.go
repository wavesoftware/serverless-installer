package cli

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

var help = cli.HelpCommand("display help information")

type runner struct {
	args []string
}

// Runner - runs main app
type Runner interface {
	Run() int
}

// NewRunnerWithArgs - creates a runner from args
func NewRunnerWithArgs(args []string) Runner {
	return runner{
		args: args,
	}
}

// NewRunner - creates a default runner
func NewRunner() Runner {
	return NewRunnerWithArgs(os.Args)
}

func (r runner) Run() int {
	if err := cli.Root(
		interactive,
		cli.Tree(help),
		cli.Tree(deploy),
	).Run(r.args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
