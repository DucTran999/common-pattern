package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func BuildFilePath(fileName string) (string, error) {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %w", err)
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

	return "", errors.New("go.mod not found")
}
