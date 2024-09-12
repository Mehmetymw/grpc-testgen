package main

import (
	"fmt"
	"os"

	"github.com/mehmetymw/grpc-testgen/internal/config"
	"github.com/mehmetymw/grpc-testgen/internal/generator"

	"github.com/spf13/cobra"
)

func main() {
	var cfgFile string

	var rootCmd = &cobra.Command{
		Use:   "grpc-testgen",
		Short: "A CLI tool for generating gRPC test cases",
		Run: func(cmd *cobra.Command, args []string) {
			if err := run(cfgFile); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "configs/grpc-testgen.yaml", "config file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}

func run(cfgFile string) error {
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		return err
	}

	return generator.GenerateTestCases(cfg)
}
