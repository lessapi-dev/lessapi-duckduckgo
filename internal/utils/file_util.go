package utils

import (
	"io"
	"os"
)

// ReadLocalFile reads the content of the file with the given filename.
func ReadLocalFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
