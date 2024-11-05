package utils

import (
	"os"
	"testing"
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

// Test for ParseFile function
func TestParseFile(t *testing.T) {
	tmpfile, err := createTempFile(testFileData)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	colony, err := ParseFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	// Assertions
	if colony.NumberOfAnts != 5 {
		t.Errorf("Expected 5 ants, got %d", colony.NumberOfAnts)
	}
	if colony.Start != "A" {
		t.Errorf("Expected start room 'A', got '%s'", colony.Start)
	}
	if colony.End != "B" {
		t.Errorf("Expected end room 'B', got '%s'", colony.End)
	}
	if len(colony.Rooms) != 3 {
		t.Errorf("Expected 3 rooms, got %d", len(colony.Rooms))
	}
	if len(colony.Links["A"]) != 2 {
		t.Errorf("Expected room 'A' to have 2 links, got %d", len(colony.Links["A"]))
	}
}
