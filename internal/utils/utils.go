package utils

import (
	"os"
	"path/filepath"
)

// GetProtoFiles retrieves all .proto files from the input directory
func GetProtoFiles(inputDir string) ([]string, error) {
	var protoFiles []string

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return protoFiles, nil
}

// CreateDirIfNotExists creates a directory if it does not exist
func CreateDirIfNotExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateFile creates a file with the specified content
func CreateFile(filePath string, content []byte) error {
	err := os.WriteFile(filePath, content, 0644)
	return err
}
