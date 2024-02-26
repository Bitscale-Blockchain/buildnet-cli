package utils

import (
	"fmt"
	"io"
	"os"
)

// ReadFile reads the JSON file and unmarshals it into the Application struct
func ReadFile(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	return data, nil

}
