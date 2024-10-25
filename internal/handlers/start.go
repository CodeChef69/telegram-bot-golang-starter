package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// StartHandler handles the /start command
func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Welcome to the bot! 🚀",
	}
	b.SendMessage(ctx, msg)
}
