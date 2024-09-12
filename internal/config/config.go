package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Language     string `yaml:"language"`
	InputDir     string `yaml:"input"`
	OutputDir    string `yaml:"output"`
	ChatGPTKey   string `yaml:"chatgpt_api_key"`
	ChatGPTModel string `yaml:"chatgpt_model"`
	GRPCServer   string `yaml:"grpc_server"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
