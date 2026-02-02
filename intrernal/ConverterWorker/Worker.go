package converterworker

import (
	"context"
	"io"
	"log"

	jobservice "github.com/Kosk0l/TgBotConverter/intrernal/Services/jobService"
	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
)

type ConverterRepository interface {
	GetJob(ctx context.Context, job domains.Job, reader io.Reader) (error)
}

// В бесконечном цикле смотрит очередь на появление новых джоб
// Верхний(адаптер) слой архитектуры
type Worker struct {
	js *jobservice.JobService
	cn ConverterRepository
}

// Конструктор
func NewWorker(js *jobservice.JobService, cn ConverterRepository) (*Worker) {
	return &Worker{
		js: js,
		cn: cn,
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