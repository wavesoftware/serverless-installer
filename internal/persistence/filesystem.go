package persistence

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type filesystemWriter struct{}

func (p filesystemWriter) Write(contents []byte, path string) error {
	log.
		WithField("size", len(contents)).
		WithField("path", path).
		Debug("Writing file")
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, contents, 0644)
	return err
}
