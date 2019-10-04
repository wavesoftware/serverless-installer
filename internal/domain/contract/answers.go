package contract

import "github.com/wavesoftware/serverless-installer/internal/domain/model"

// AnswersFilename - a base filename of answers file
const AnswersFilename string = "serverless-installer-answers.yaml"

// LocationAwareAnswers - a struct that hold answers and location to saved them
type LocationAwareAnswers struct {
	*model.Answers
	Path string
}

// AnswersSaver - saves answers to a file
type AnswersSaver interface {
	Save(answers model.Answers, path string) error
}
