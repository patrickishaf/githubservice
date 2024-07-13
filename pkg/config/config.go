package config

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
