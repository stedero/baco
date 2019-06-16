package rio

import (
	"io"
	"log"
	"os"
)

// OpenFile opens a file for reading.
// If an error occurs then the program terminates with a panic message.
func OpenFile(filename string) io.ReadCloser {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", filename, err)
	}
	return file
}

// CreateFile creates a file for writing.
// If an error occurs then the program terminates with a panic message.
func CreateFile(filename string) io.WriteCloser {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed to create file %s: %v", filename, err)
	}
	return file
}
