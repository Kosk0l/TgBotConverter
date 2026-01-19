package bot

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/Kosk0l/TgBotConverter/intrernal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext( // root context
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	cfg := config.Load()

	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	app.Run(ctx)
}