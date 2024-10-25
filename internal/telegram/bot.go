package telegram

import (
	"context"
	"log"
	"tg-bot-go/internal/handlers"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// StartBot initializes and starts the bot
func StartBot(token string) error {
	// Create bot configuration
	b, err := bot.New(token, bot.WithDefaultHandler(defaultHandler))
	if err != nil {
		return err
	}
	// Register command handlers
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handlers.HelpHandler)

	// Start receiving updates
	log.Println("Bot is running...")
	b.Start(context.Background())

	return nil
}

// defaultHandler handles unexpected messages
func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Unknown command. Try /help.",
	}
	b.SendMessage(ctx, msg)
}
