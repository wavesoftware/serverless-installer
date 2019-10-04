package interactive

import "github.com/wavesoftware/serverless-installer/internal/domain/contract"

// Create - creates a screens object that points to it's first element
func Create(answers *contract.LocationAwareAnswers) Screen {
	scope := newScope(answers)
	answersfile := newAnswersfile(&scope.screen)
	scope.next = answersfile
	return scope
}
