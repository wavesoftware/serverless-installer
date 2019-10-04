package logic

import (
	"github.com/wavesoftware/serverless-installer/internal/domain/contract"
	"github.com/wavesoftware/serverless-installer/internal/persistence"
)

// CreateAnswersSaver - creates an object of AnswersSaver
func CreateAnswersSaver() contract.AnswersSaver {
	return answersService{
		fileWriter: persistence.CreateFileWriter(),
	}
}
