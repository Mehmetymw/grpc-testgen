package main

import (
	"os"

	"github.com/mehmetymw/grpc-testgen/internal/config"
	"github.com/mehmetymw/grpc-testgen/internal/generator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
}

func main() {
	initLogger()

	var cfgFile string

	var rootCmd = &cobra.Command{
		Use:   "grpc-testgen",
		Short: "A CLI tool for generating gRPC test cases",
		Run: func(cmd *cobra.Command, args []string) {
			if err := run(cfgFile); err != nil {
				log.Fatal().Err(err).Msg("Error running grpc-testgen")
			}
		},
	}

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "configs/grpc-testgen.yaml", "Config file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Error executing command")
	}
}

func run(cfgFile string) error {
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		return err
	}

	return generator.GenerateTestCases(cfg)
}
