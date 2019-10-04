package gateway

// FileWriter - writes contents to given file path creating directory
// structure if needed
type FileWriter interface {
	Write(contents []byte, path string) error
}
