package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	MongoConfig struct {
		URI             string
		DatabaseName    string
		Collection      string
		TimeoutDuration string
	}
	Config struct {
		Env   string
		Port  string
		Mongo MongoConfig
	}
)

func loadConfig() *Config {
	return &Config{
		Env:  os.Getenv("ENV"),
		Port: os.Getenv("PORT"),
		Mongo: MongoConfig{
			URI:             os.Getenv("MONGO_URI"),
			DatabaseName:    os.Getenv("MONGO_DB_NAME"),
			Collection:      os.Getenv("MONGO_COLLECTION"),
			TimeoutDuration: os.Getenv("MONGO_TIMEOUT_DURATION"),
		},
	}
}

func NewConfig(path string) (*Config, error) {
	cfg := loadConfig()

	if path == "" {
		path = ".env"
	}

	if cfg.Env == "" || cfg.Env == "LOCAL" {
		cfg.Env = "LOCAL"

		err := godotenv.Load(path)
		if err != nil {
			return nil, err
		}
		cfg = loadConfig()

	}
	fmt.Println(cfg)
	return cfg, nil
}
