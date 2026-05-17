package main

import (
	"context"
	"log"
	"oracle/internal/handler"
	"oracle/pkg/config"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	if cfg.TelegramToken == "" || cfg.TelegramToken == "YOUR_BOT_TOKEN_HERE" {
		log.Fatal("Please set TELEGRAM_TOKEN in .env file or replace in config")
	}

	b, err := bot.New(cfg.TelegramToken)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	h := handler.NewTelegramHandler(b)
	h.RegisterHandlers()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Println("bot started")
	b.Start(ctx)
}
