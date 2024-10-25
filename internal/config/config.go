package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config stores configuration for the bot.
type Config struct {
	Token string
}

// LoadConfig loads configuration from the .env file and environment variables.
func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from environment variables")
	}

	// Fetch the token from environment variables
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	return &Config{
		Token: token,
	}
}
