package jobservice

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Kosk0l/TgBotConverter/intrernal/models"
	"github.com/google/uuid"
)

// Абстракция для cache
type JobRepository interface {
	SetToList(ctx context.Context, jobId string) (error)
	SetToHash(ctx context.Context, job models.Job) (error)
	GetFromList(ctx context.Context) (string, error)
	GetFromHash(ctx context.Context, jobId string) (*models.Job, error)
	DeleteKey(ctx context.Context, jobId string) (error)
	SetToListR(ctx context.Context, jobId string) (error)
}

// Абстракция для обработки сырых файлов
type FileRepository interface {
	SetObject(ctx context.Context, jobId string, r io.Reader, size int64, contentType string) (error)
	GetObject(ctx context.Context, jobId string) (io.Reader, error)
	DeleteFile(ctx context.Context, jobId string) (error)
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
func (js *JobService) CreateJob(ctx context.Context, job models.Job) (string, error) {
	job.JobID = uuid.NewString()

	if err := js.fileRepo.SetObject(ctx, job.JobID, nil, 0, ""); err != nil {
		return job.JobID, fmt.Errorf("jobservice - error in setobject: %w", err)
	}

	if err := js.repo.SetToHash(ctx, job); err != nil {
		if err2 := js.fileRepo.DeleteFile(ctx, job.JobID); err2 != nil {
			log.Printf("rollback error DeleteFile: %v", err2)
		}
		return job.JobID, fmt.Errorf("jobservice - error in settohash: %w", err)
	}

	if err := js.repo.SetToList(ctx, job.JobID); err != nil {
		if err2 := js.fileRepo.DeleteFile(ctx, job.JobID); err2 != nil {
			log.Printf("rollback error DeleteFile: %v", err2)
		}
		if err2 := js.repo.DeleteKey(ctx, job.JobID); err2 != nil {
			log.Printf("rollback error DeleteKey: %v", err2)
		}
		return job.JobID, fmt.Errorf("jobservice - error in settolist: %w", err)
	}

	return job.JobID, nil
}

// Получить job //TODO: реализовать проверку наличия данных в list
func (js *JobService) GetJob(ctx context.Context) (*models.Job, error) {
	jobId, err := js.repo.GetFromList(ctx)
	if err != nil {
		return nil, fmt.Errorf("jobservice - error in getjob: %w", err)
	}

	//TODO: добавить reader;
	if _, err := js.fileRepo.GetObject(ctx, jobId); err != nil {
		if err2 := js.repo.SetToList(ctx, jobId); err2 != nil {
			log.Printf("rollback error SetToList: %v", err2)
		}
		return nil, fmt.Errorf("jobservice - error in getobject: %w", err)
	}

	job, err := js.repo.GetFromHash(ctx, jobId)
	if err != nil {
		if err2 := js.repo.SetToList(ctx, jobId); err2 != nil {
			log.Printf("rollback error SetToList: %v", err2)
		}
		return nil, fmt.Errorf("jobservice - error in gethashdata: %w", err)
	}

	return job, nil
}