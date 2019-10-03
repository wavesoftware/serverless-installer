package interactive

import (
	survey "github.com/AlecAivazis/survey/v2"
)

type scopeScreen struct {
	screenData
}

const jaeger string = "Jaeger Operator"
const elasticsearch string = "Elasticsearch Operator"
const kiali string = "Kiali Operator"
const istio string = "Maistra Istio Operator"
const serverless string = "Serverless Operator"

func (s scopeScreen) Display() {
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
	s.screenData.answers.Jeager.Install = contains(selected, jaeger)
	s.screenData.answers.Elasticsearch.Install = contains(selected, elasticsearch)
	s.screenData.answers.Kiali.Install = contains(selected, kiali)
	s.screenData.answers.Istio.Install = contains(selected, istio)
	s.screenData.answers.Serverless.Install = contains(selected, serverless)
}

func contains(elements []string, searched string) bool {
	for _, element := range elements {
		if element == searched {
			return true
		}
	}
	return false
}
