package converterworker

import (
	"context"
	"io"
	"log"
	"log/slog"

	jobservice "github.com/Kosk0l/TgBotConverter/internal/Services/jobService"
	"github.com/Kosk0l/TgBotConverter/internal/domains"
)

type ConverterRepository interface {
	GetJob(ctx context.Context, job domains.Job, reader io.Reader) (error)
}

// В бесконечном цикле смотрит очередь на появление новых джоб
// Верхний(адаптер) слой архитектуры
type Worker struct {
	js *jobservice.JobService
	cn ConverterRepository
	logger *slog.Logger
}

// Конструктор
func NewWorker(js *jobservice.JobService, cn ConverterRepository, logger *slog.Logger) (*Worker) {
	return &Worker{
		js: js,
		cn: cn,
		logger: logger,
	}
}

//====================================================================================================

// Запуск воркера
func (w *Worker) Run(ctx context.Context) () {
	log.Println("worker running..")
	for {
		select {
		case <-ctx.Done():
			log.Println("Worker stopped")
			return
		default:
			err := w.processOne(ctx)
			if err != nil {
				log.Printf("worker error: %v", err)
			}
		}
	}
}