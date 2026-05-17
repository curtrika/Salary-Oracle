package handler

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TelegramHandler struct {
	bot *bot.Bot
}

func NewTelegramHandler(b *bot.Bot) *TelegramHandler {
	return &TelegramHandler{
		bot: b,
	}
}

func (h *TelegramHandler) RegisterHandlers() {
	h.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, h.handleStart)

	h.bot.RegisterHandler(bot.HandlerTypeMessageText, "/ping", bot.MatchTypePrefix, h.handlePing)
}

func (h *TelegramHandler) handleStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := "👋 Привет!"

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   msg,
	})
}

func (h *TelegramHandler) handlePing(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("pong", update.Message.Text),
	})
}
