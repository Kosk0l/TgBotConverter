package logger

import (
	"log/slog"
	"os"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/lmittmann/tint"
)

// Конструктор
func NewLogger(cfg config.Config) (*slog.Logger) {
	var handler slog.Handler
	mode := cfg.Log.Mode

	// Выбор Mode
	switch mode {
	case "dev":
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level: slog.LevelDebug,
			TimeFormat: "15:04:05",
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