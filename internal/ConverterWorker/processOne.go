package converterworker

import (
	"context"
	"fmt"
	"log"
	"log/slog"
)

// Делигировать метаданные и reader
func (w *Worker) processOne(ctx context.Context) (error) {
	// Получить 
	job, reader, err := w.js.GetJob(ctx)
	if err != nil {
		w.logger.Error("worker error get the job",
			slog.Any("error", err),
			slog.String("job_id", job.JobID),
		)
		return fmt.Errorf("get job error: %w", err)
	}
	log.Printf("worker got job: %s", job.JobID)

	// Обработка
	if err := w.cn.GetJob(ctx, job, reader); err != nil {
		w.logger.Error("converter error get the job",
			slog.Any("error", err),
			slog.String("job_id", job.JobID),
		)
		return fmt.Errorf("handler job %s error: %v", job.JobID, err)
	}

	w.logger.Info("worker success convert the job",
		slog.String("job_id", job.JobID),
	)
	return nil
}
