package cli

import (
	log "github.com/sirupsen/logrus"

	"github.com/mkideal/cli"
	factory "github.com/wavesoftware/serverless-installer/internal/cli/interactive"
	"github.com/wavesoftware/serverless-installer/internal/domain/contract"
	"github.com/wavesoftware/serverless-installer/internal/domain/logic"
	"github.com/wavesoftware/serverless-installer/internal/domain/model"
)

// interactive command
type interactiveT struct {
	cli.Helper
	Logging
}

var interactive = &cli.Command{
	Desc: `Installer for Openshift Serverless
	
Run without any options to enter interactive shell that will guide you
through installation process and produce an answers file. That answers
file can by applied on current Openshift cluster with deploy command.`,

	Argv: func() interface{} { return new(interactiveT) },
	Fn:   operate,
}

func operate(ctx *cli.Context) error {
	argv := ctx.Argv().(*interactiveT)
	log.SetOutput(ctx.Writer())
	log.SetLevel(argv.Logging.Level.value)

	answers := contract.LocationAwareAnswers{
		Answers: &model.Answers{},
	}
	screen := factory.Create(&answers)
	display(screen)
	return save(answers)
}

func display(screen factory.Screen) {
	for {
		screen.Display()
		if screen.Next() != nil {
			screen = screen.Next()
		} else {
			break
		}
	}
}

func save(answers contract.LocationAwareAnswers) error {
	saver := logic.CreateAnswersSaver()
	err := saver.Save(*answers.Answers, answers.Path)
	if err == nil {
		log.
			WithField("Answers file", answers.Path).
			Info("Answers file has been saved.")
	}
	return err
}
