package converterworker

import (
	"context"
	"log"

	jobservice "github.com/Kosk0l/TgBotConverter/intrernal/Services/jobService"
)

// В бесконечном цикле смотрит очередь на появление новых джоб
// Верхний(адаптер) слой архитектуры
type Worker struct {
	js *jobservice.JobService
}

// Конструктор
func NewWorker(js *jobservice.JobService) (*Worker) {
	return &Worker{
		js: js,
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