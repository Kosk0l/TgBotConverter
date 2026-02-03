package jobservice

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Kosk0l/TgBotConverter/internal/domains"
	"github.com/google/uuid"
)

// Абстракция для cache
type JobRepository interface {
	SetToList(ctx context.Context, jobId string) (error)
	SetToHash(ctx context.Context, job domains.Job) (error)

	GetFromList(ctx context.Context) (string, error)
	GetFromHash(ctx context.Context, jobId string) (domains.Job, error)

	DeleteKey(ctx context.Context, jobId string) (error)
}

// Абстракция для обработки сырых файлов
type FileRepository interface {
	SetObject(ctx context.Context, jobId string, fileUrl string, size int64, contentType string) (error)
	GetObject(ctx context.Context, jobId string) (io.Reader, error)
	DeleteObject(ctx context.Context, jobId string) (error)
	ExistObject(ctx context.Context, jobId string)(bool, error)
}

// Бизнес-логика для работы с запросами
type JobService struct {
	repo 	 JobRepository
	fileRepo FileRepository
}

// Конструктор
func NewJobService(repo JobRepository, fileRepo FileRepository) (*JobService) {
	return &JobService{
		repo: repo,
		fileRepo: fileRepo,
	}
}

//====================================================================================================

// Создать job 
func (js *JobService) CreateJob(ctx context.Context, job domains.Job, jobObj domains.Object) (string, error) {
	job.JobID = uuid.NewString() // уникальный id

	// Добавить file
	if err := js.fileRepo.SetObject(ctx, job.JobID, jobObj.FlieURL, jobObj.Size, jobObj.ContentType); err != nil {
		return "", fmt.Errorf("jobservice - error in setobject: %w", err)
	}

	// Положить в hash
	if err := js.repo.SetToHash(ctx, job); err != nil {
		if err2 := js.fileRepo.DeleteObject(ctx, job.JobID); err2 != nil {
			log.Printf("rollback error DeleteFile: %v", err2)
		}
		return "", fmt.Errorf("jobservice - error in settohash: %w", err)
	}

	// Добавить в очередь
	if err := js.repo.SetToList(ctx, job.JobID); err != nil {
		if err2 := js.fileRepo.DeleteObject(ctx, job.JobID); err2 != nil {
			log.Printf("rollback error DeleteFile: %v", err2)
		}
		if err2 := js.repo.DeleteKey(ctx, job.JobID); err2 != nil {
			log.Printf("rollback error DeleteKey: %v", err2)
		}
		return "", fmt.Errorf("jobservice - error in settolist: %w", err)
	}

	// вернуть id
	return job.JobID, nil
}

// Получить job
func (js *JobService) GetJob(ctx context.Context) (domains.Job, io.Reader, error) {
	// Получить JobId
	jobId, err := js.repo.GetFromList(ctx)
	if err != nil {
		return domains.Job{}, nil, fmt.Errorf("jobservice - error in getjob: %w", err)
	}

	// Получить file по jobId
	reader, err := js.fileRepo.GetObject(ctx, jobId)
	if err != nil {
		if err2 := js.repo.SetToList(ctx, jobId); err2 != nil {
			log.Printf("rollback error SetToListR: %v", err2)
		}
		return domains.Job{}, nil, fmt.Errorf("jobservice - error in getobject: %w", err)
	}

	// Получить метаданные по JobId Во что конвертировать
	job, err := js.repo.GetFromHash(ctx, jobId)
	if err != nil {
		if err2 := js.repo.SetToList(ctx, jobId); err2 != nil {
			log.Printf("rollback error SetToListR: %v", err2)
		}
		return domains.Job{}, nil, fmt.Errorf("jobservice - error in gethashdata: %w", err)
	}

	// Вернуть метаданные и reader
	return job, reader, nil
}