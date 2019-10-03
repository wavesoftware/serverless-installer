package cli

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

type deployT struct {
	cli.Helper
	Answers clix.File `cli:"a,answers" usage:"an answers file produced by interactive shell"`
}

var deploy = &cli.Command{
	Name: "deploy",
	Desc: "Deploys a Serverless onto Openshift cluster using ansers file",
	Argv: func() interface{} { return new(deployT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*deployT)
		ctx.String("Hello, from deploy command. Answers file is %s\n", argv.Answers)
		return nil
	},
}
