package main

import (
	"log"
	"tg-bot-go/internal/config"
	"tg-bot-go/internal/telegram"
)

func main() {
	// Load the configuration (like the bot token)
	cfg := config.LoadConfig()

	// Initialize and run the bot
	if err := telegram.StartBot(cfg.Token); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}
