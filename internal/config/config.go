package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type Config struct{
	MongoURI string
	MongoDBName string
	Port string
}

func LoadConfig() (*Config, error) {

	// godotenv.Load() Load environment variables from .env file and set to process environment
	err := godotenv.Load()
	if err != nil {
		return &Config{}, fmt.Errorf("failed to load .env file: %v", err)
	}


	mongoURI, err := extractConfigValue("MONGO_URI")
	if err != nil {
		return &Config{}, err
	}

	mongoDBName, err := extractConfigValue("MONGO_DB_NAME")
	if err != nil {
		return &Config{}, err
	}

	port, err := extractConfigValue("SERVER_PORT")
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		MongoURI: mongoURI,
		MongoDBName: mongoDBName,
		Port: port,
	}, nil
}


func extractConfigValue(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}
