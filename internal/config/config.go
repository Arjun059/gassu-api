package config

import "os"

type Config struct {
	GOOSE_DRIVER        string
	GOOSE_DBSTRING      string
	GOOSE_MIGRATION_DIR string
	GOOSE_TABLE         string
	POSTGRES_DB_URI     string
}

func Load() (*Config, error) {

	var config Config
	config.GOOSE_DRIVER = os.Getenv("GOOSE_DRIVER")
	config.GOOSE_DBSTRING = os.Getenv("GOOSE_DBSTRING")
	config.GOOSE_MIGRATION_DIR = os.Getenv("GOOSE_MIGRATION_DIR")
	config.GOOSE_TABLE = os.Getenv("GOOSE_TABLE")
	config.POSTGRES_DB_URI = os.Getenv("POSTGRES_DB_URI")

	return &config, nil
}
