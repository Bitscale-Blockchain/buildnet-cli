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

func DeleteAndCreateDirectory(dirPath string) error {
	// Check if the directory exists
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		// Directory exists, delete it
		err := os.RemoveAll(dirPath)
		if err != nil {
			return fmt.Errorf("failed to delete directory: %v", err)
		}
	}

	// Create the directory
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	return nil
}

func CreateDirectoryIfNotExists(dirPath string) error {
	// Check if the directory already exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	return nil
}

func GetCurrentDirectory() (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return cwd, nil
}
