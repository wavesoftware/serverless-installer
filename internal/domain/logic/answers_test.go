package logic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wavesoftware/serverless-installer/internal/domain/model"
)

var files []file

func TestSave(t *testing.T) {
	// given
	files = nil
	service := answersService{
		fileWriter: stubWriter{},
	}
	answers := model.Answers{}
	answers.Jeager.Install.Enabled = false
	path := "/tmp/answers.yaml"
	expectedYaml := strings.ReplaceAll(`jeager:
	install:
		enabled: false
elasticsearch:
	install:
		enabled: false
kiali:
	install:
		enabled: false
istio:
	install:
		enabled: false
serverless:
	install:
		enabled: false
`, "\t", "  ")

	// when
	err := service.Save(answers, path)

	// then
	assert.Nil(t, err)
	assert.Len(t, files, 1)
	first := files[0]
	assert.Equal(t, path, first.path)
	assert.Equal(t, expectedYaml, string(first.contents))
}

type file struct {
	contents []byte
	path     string
}

type stubWriter struct{}

func (w stubWriter) Write(contents []byte, path string) error {
	files = append(files, file{
		contents: contents,
		path:     path,
	})
	return nil
}
