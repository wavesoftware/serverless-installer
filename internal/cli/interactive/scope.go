package interactive

import (
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/wavesoftware/serverless-installer/internal/domain/contract"
)

type scope struct {
	screen
}

func newScope(answers *contract.LocationAwareAnswers) scope {
	val := scope{
		screen{
			answers: answers,
			next:    nil,
		},
	}
	return val
}

const jaeger string = "Jaeger Operator"
const elasticsearch string = "Elasticsearch Operator"
const kiali string = "Kiali Operator"
const istio string = "Maistra Istio Operator"
const serverless string = "Serverless Operator"

func (s scope) Display() {
	possible := []string{
		serverless,
		istio,
		jaeger,
		elasticsearch,
		kiali,
	}

	selected := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select what operators to install:",
		Help: `Serverless Operator needs Istio to be installed. Maistra Istio
 Operator needs Jeager, Elasticsearech, and Kiali to be installed
 beforehand. If you have some of them installed already, uncheck them
 from selection.`,
		Options: possible,
		Default: possible,
	}
	survey.AskOne(prompt, &selected)
	answers := s.screen.answers.Answers
	answers.Jeager.Install.Enabled = contains(selected, jaeger)
	answers.Elasticsearch.Install.Enabled = contains(selected, elasticsearch)
	answers.Kiali.Install.Enabled = contains(selected, kiali)
	answers.Istio.Install.Enabled = contains(selected, istio)
	answers.Serverless.Install.Enabled = contains(selected, serverless)
}

func contains(elements []string, searched string) bool {
	for _, element := range elements {
		if element == searched {
			return true
		}
	}
	return false
}
