package cli

import (
	"os"

	"github.com/mkideal/cli"
	"github.com/sirupsen/logrus"
)

func help() *cli.Command {
	return cli.HelpCommand("display help information")
}

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
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05.000",
	})
	if err := cli.Root(
		interactive(),
		cli.Tree(help()),
		cli.Tree(deploy()),
	).Run(r.args[1:]); err != nil {
		logrus.Error(err)
		return 1
	}
	return 0
}
