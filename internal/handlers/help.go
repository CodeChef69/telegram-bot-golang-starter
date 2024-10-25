package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// HelpHandler handles the /help command
func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Here is the help message.",
	}
	b.SendMessage(ctx, msg)
}
