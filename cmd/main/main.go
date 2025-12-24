package main

import (
	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/Kosk0l/TgBotConverter/intrernal/app"
)

func main() {
	cfg := config.Load()
	app := app.NewApp(cfg)
	app.Run()
}