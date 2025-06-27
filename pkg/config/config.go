package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServerConfig struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env              string `yaml:"env" env-required:"true" env-default:"dev"`
	StoragePath      string `yaml:"storage_path" env-required:"true"`
	HTTPServerConfig `yaml:"http_server"`
}

func MustLoadConfig() *Config {
	var configPath string
	configPath = os.Getenv("MIYUKI_CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "config/prod.yaml", "Path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Configuration file path is not set. Please set the MIYUKI_CONFIG_PATH environment variable or use the --config flag.")
		}

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist at path: %s", configPath)
	}

	var config Config
	err := cleanenv.ReadConfig(configPath, &config)

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	return &config
}
