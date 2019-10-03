package cli

import (
	"github.com/mkideal/cli"
	factory "github.com/wavesoftware/serverless-installer/internal/cli/interactive"
	"github.com/wavesoftware/serverless-installer/internal/domain/model"
)

// interactive command
type interactiveT struct {
	cli.Helper
}

var interactive = &cli.Command{
	Desc: `Installer for Openshift Serverless
	
Run without any options to enter interactive shell that will guide you
through installation process and produce an answers file. That answers
file can by applied on current Openshift cluster with deploy command.`,

	Argv: func() interface{} { return new(interactiveT) },
	Fn: func(ctx *cli.Context) error {
		answers := model.Answers{}
		screen := factory.Create(&answers)
		for {
			screen.Display()
			if screen.Next() != nil {
				screen = screen.Next()
			} else {
				break
			}
		}
		ctx.String("Answers are: %v", answers)
		return nil
	},
}
