package generator

import (
	"fmt"
	"path/filepath"

	"github.com/mehmetymw/grpc-testgen/internal/config"
	"github.com/mehmetymw/grpc-testgen/internal/utils"
)

func GenerateTestCases(cfg *config.Config) error {
	protoFiles, err := utils.GetProtoFiles(cfg.Input)
	if err != nil {
		return err
	}

	for _, protoFile := range protoFiles {
		prompt := fmt.Sprintf("Generate test cases for the following .proto file: %s", protoFile)
		tests, err := getChatGPTResponses(cfg, prompt)
		if err != nil {
			return err
		}

		outputFile := fmt.Sprintf("%s/%s_test.go", cfg.Output, filepath.Base(protoFile))
		err = utils.CreateFile(outputFile, []byte(tests))
		if err != nil {
			return err
		}
	}

	return nil
}
