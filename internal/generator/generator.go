package generator

import (
	"fmt"

	"github.com/mehmetymw/grpc-testgen/internal/config"
	"github.com/mehmetymw/grpc-testgen/internal/utils"
)

func GenerateTestCases(cfg *config.Config) error {
	protoFiles, err := utils.GetProtoFiles(cfg.InputDir)
	if err != nil {
		return err
	}

	for _, file := range protoFiles {
		err := generateForFile(file, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateForFile(file string, cfg *config.Config) error {
	fmt.Println("Generating test cases for", file)
	return nil
}
