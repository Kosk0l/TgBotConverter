package converterworker

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
)

// Делигировать метаданные и reader
func (w *Worker) processOne( ctx context.Context) (error) {
	// Получить 
	job, reader, err := w.js.GetJob(ctx)
	if err != nil {
		return fmt.Errorf("get job error: %w", err)
	}
	log.Printf("worker got job: %s", job.JobID)

	// Обработка
	if err := w.handleJob(ctx, job, reader); err != nil {
		return fmt.Errorf("handler job %s error: %v", job.JobID, err)
	}

	log.Printf("job %s compleated", job.JobID)
	return nil
}

// передать в слой бизнес-логики
func (w *Worker) handleJob(ctx context.Context, job *domains.Job, reader io.Reader) (error) {
	//TODO: реализация

	return nil
}