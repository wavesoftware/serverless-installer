package persistence

import "github.com/wavesoftware/serverless-installer/internal/domain/gateway"

// CreateFileWriter - factory for FileWriter
func CreateFileWriter() gateway.FileWriter {
	return filesystemWriter{}
}
