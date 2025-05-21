package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func BuildFilePath(fileName string) (string, error) {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %v", err)
	}

	for {
		rootPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(rootPath); err == nil {
			fullPath := filepath.Join(dir, fileName)
			return fullPath, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("go.mod not found")
}
