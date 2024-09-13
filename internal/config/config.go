package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Language          string `mapstructure:"language"`
	Input             string `mapstructure:"input"`
	Output            string `mapstructure:"output"`
	ChatGPTAPIKey     string `mapstructure:"chatgpt_api_key"`
	ChatGPTModel      string `mapstructure:"chatgpt_model"`
	GRPCServerAddress string `mapstructure:"grpc_server_address"`
}

func LoadConfig(filePath string) (*Config, error) {
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read environment variables

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config file: %w", err)
	}

	return &cfg, nil
}
