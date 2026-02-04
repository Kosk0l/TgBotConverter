package logger

import (
	"log/slog"
	"os"

	"github.com/Kosk0l/TgBotConverter/config"
	
)

// Конструктор
func NewLogger(cfg config.Config) (*slog.Logger) {
	var handler slog.Handler
	mode := cfg.Log.Mode

	// Выбор Mode
	switch mode {
	case "dev":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		return slog.New(handler)

	case "prod":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		return slog.New(handler)
	}

	return slog.Default()
}