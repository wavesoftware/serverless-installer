package interactive

import "github.com/wavesoftware/serverless-installer/internal/domain/model"

// Create - creates a screens object that points to it's first element
func Create(answers *model.Answers) Screen {
	root := scopeScreen{
		screenData: screenData{
			answers: answers,
			next:    nil,
		},
	}
	return root
}
