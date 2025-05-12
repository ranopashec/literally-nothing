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
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.BusinessMessage != nil {
		fmt.Println(update.BusinessMessage.Text)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:               update.BusinessMessage.Chat.ID,
			BusinessConnectionID: update.BusinessMessage.BusinessConnectionID,
			Text:                 "Sorry, this user is in the process of discovering his true potential, please contact him later (in 3-4 months).",
		})
	}
}
