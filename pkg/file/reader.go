package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ReadJSON reads JSON data from a file and unmarshals it into the provided struct
func ReadJSON(filePath string, v interface{}) error {
	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Parse JSON
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to parse JSON from %s: %w", filePath, err)
	}

	return nil
}

// GetAbsolutePath returns the absolute path of a file relative to the current working directory
func GetAbsolutePath(relativePath string) (string, error) {
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}
	return absPath, nil
}

// FileExists checks if a file exists at the given path
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
