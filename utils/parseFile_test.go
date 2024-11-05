package utils

import (
	"os"
)

// Mock data for tests
var testFileData = `
5
##start
A 0 0
##end
B 5 5
C 2 2
A-B
A-C
C-B
`

// Helper function to create a temporary file for testing
func createTempFile(data string) (*os.File, error) {
	tmpfile, err := os.CreateTemp("", "antcolony")
	if err != nil {
		return nil, err
	}
	if _, err := tmpfile.Write([]byte(data)); err != nil {
		tmpfile.Close()
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}
	return tmpfile, nil
}
