package twsinput

import (
	"fmt"
	"os"
)

func OpenDevice(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open device on path %s, %w", path, err)
	}

	return file, err
}

func CloseDevice(file *os.File) error {
  return file.Close()
}
