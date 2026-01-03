package main

import (
	"context"
	"log"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/Kosk0l/TgBotConverter/intrernal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()

	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	app.Run(ctx)
}