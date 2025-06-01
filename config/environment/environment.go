package environment

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	ServerPort       = getEnv("SERVER_PORT")
	ServerMode       = getEnv("SERVER_MODE")
	DatabaseDriver   = getEnv("DATABASE_DRIVER")
	DatabaseUsername = getEnv("DATABASE_USERNAME")
	DatabasePassword = getEnv("DATABASE_PASSWORD")
	DatabaseHost     = getEnv("DATABASE_HOST")
	DatabasePort     = getEnv("DATABASE_PORT")
	DatabaseName     = getEnv("DATABASE_NAME")
)

func getEnv(key string) string {
	LoadEnv()
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable not set: " + key)
	}
	return value
}
func LoadEnv() {
	// Try loading the .env file from the root directory
	if err := godotenv.Load(".env"); err != nil {
		if err := godotenv.Load("../../.env"); err != nil {
			panic("Warning: No .env file found in root or ../../.env, using default values")
		}
	}
}
