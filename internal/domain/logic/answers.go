package logic

import (
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"github.com/wavesoftware/serverless-installer/internal/domain/gateway"
	"github.com/wavesoftware/serverless-installer/internal/domain/model"
)

type answersService struct {
	fileWriter gateway.FileWriter
}

func (service answersService) Save(answers model.Answers, path string) error {
	log.
		WithField("path", path).
		WithField("answers", answers).
		Debug("Saved called")
	contents, err := yaml.Marshal(answers)
	if err != nil {
		return err
	}
	log.
		WithField("yaml", string(contents)).
		Trace("Yaml")
	return service.fileWriter.Write(contents, path)
}
