package persistence

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/wavesoftware/serverless-installer/internal/domain/contract"
)

func TestFilesystemWriter(t *testing.T) {
	// given
	tmp, err := ioutil.TempDir("", "test-filesystem-writer")
	assert.Nil(t, err)
	defer os.RemoveAll(tmp)
	writer := filesystemWriter{}
	path := filepath.Join(tmp, uuid.New().String(), contract.AnswersFilename)
	contents := []byte("This is test")

	// when
	err = writer.Write(contents, path)

	// then
	assert.Nil(t, err)
}

func TestFilesystemWriterOnInvalidPath(t *testing.T) {
	// given
	writer := filesystemWriter{}
	path := filepath.Join(
		// invalid character in path 0x0
		string([]byte{0x0}),
		uuid.New().String(), contract.AnswersFilename,
	)
	contents := []byte("This is test")

	// when
	err := writer.Write(contents, path)

	// then
	assert.NotNil(t, err)
}
