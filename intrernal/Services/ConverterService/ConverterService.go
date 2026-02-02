package converterservice

import (
	"context"
	"io"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
)

// Бизнес-логика конвертации файлов
type Converter struct {
	//TODO:
}

// Конструктор
func NewConverterService() (*Converter) {
	return &Converter{}
}

func (c *Converter) GetJob(ctx context.Context, job domains.Job, reader io.Reader) (error) {
	//TODO: реализация

	return nil
}