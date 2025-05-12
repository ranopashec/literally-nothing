package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("TOKEN"), opts...)
	if err != nil {
		panic(err)
	}
	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.BusinessMessage != nil && update.BusinessMessage.From.Username != "ranopashec" {
		fmt.Println(update.BusinessMessage.Text)
		b.SendMessage(ctx, &bot.SendMessageParams{
			BusinessConnectionID: update.BusinessMessage.BusinessConnectionID,
			ChatID:               update.BusinessMessage.Chat.ID,
			Text:                 "The user ignores incoming messages, please [book a time](https://calendar.app.google/iHgxhxw3WfCATTuH7) for the call.",
			ParseMode:            models.ParseMode("Markdown"),
		})
	}
}
