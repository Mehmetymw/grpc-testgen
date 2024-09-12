package utils

import (
	"os"
	"path/filepath"
)

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
