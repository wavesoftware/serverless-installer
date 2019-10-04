package cli

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
	log "github.com/sirupsen/logrus"
)

type deployT struct {
	cli.Helper
	Answers clix.File `cli:"a,answers" usage:"an answers file produced by interactive shell"`
	Logging
}

func deploy() *cli.Command {
	return &cli.Command{
		Name: "deploy",
		Desc: "Deploys a Serverless onto Openshift cluster using ansers file",
		Argv: func() interface{} { return new(deployT) },
		Fn: func(ctx *cli.Context) error {
			argv := ctx.Argv().(*deployT)
			log.SetOutput(ctx.Writer())
			log.SetLevel(argv.Logging.Level.value)

			log.
				WithField("answers-file", argv.Answers).
				Info("Hello, from deploy command.")
			return nil
		},
	}
}
