package persistence

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"

	"github.com/wavesoftware/serverless-installer/internal/domain/contract"
)

func TestFilesystemWriter(t *testing.T) {
	// given
	tmp, err := ioutil.TempDir("", "test-filesystem-writer")
	checkNoError(err)
	defer os.RemoveAll(tmp)
	writer := filesystemWriter{}
	path := filepath.Join(tmp, uuid.New().String(), contract.AnswersFilename)
	contents := []byte("This is test")

	// when
	err = writer.Write(contents, path)

	// then
	if err != nil {
		t.Error(err)
	}
}

func checkNoError(err error) {
	if err != nil {
		panic(err)
	}
}
