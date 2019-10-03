package interactive

import (
	"fmt"
	"os"
	"path/filepath"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/wavesoftware/serverless-installer/internal/domain/contract"
)

type answersfile struct {
	screen
}

func newAnswersfile(parent *screen) answersfile {
	curr := answersfile{
		screen{
			answers: parent.answers,
			next:    nil,
		},
	}
	return curr
}

func (a answersfile) Display() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := cwd
	prompt := &survey.Input{
		Message: "Select where to save a answers file:",
		Help: fmt.Sprintf(
			`Answers file will be named "%s" 
and will be placed where you asked for. If directory path don't exists, 
installer will try to create this path for you. If answers file exists at that 
location, it will be overwritten with new contents.`,
			contract.AnswersFilename,
		),
		Default: cwd,
	}
	survey.AskOne(prompt, &path)
	a.answers.Path = filepath.Join(path, contract.AnswersFilename)
}
