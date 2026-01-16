package jobservice

import (
	"context"
	"fmt"
	"io"

	"github.com/Kosk0l/TgBotConverter/intrernal/models"
)

// Абстракция для cache
type JobRepository interface {
	SetToList(ctx context.Context, jobId int64) (error)
	SetToHash(ctx context.Context, job models.Job) (error)
	GetFromList(ctx context.Context) (int64, error)
	GetFromHash(ctx context.Context, jobId int64) (*models.Job, error)
	DeleteKey(ctx context.Context, jobId int64) (error)
	SetToListR(ctx context.Context, jobId int64) (error)
}

// Абстракция для обработки сырых файлов
type FileRepository interface {
	SetObject(ctx context.Context, jobId int64, r io.Reader, size int64, contentType string) (error)
	GetObject(ctx context.Context, jobId int64) (io.Reader, error)
	DeleteFile(ctx context.Context, jobId int64) (error)
	ExistObject(ctx context.Context, jobId int64)(bool, error)
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
func (js *JobService) CreateJob(ctx context.Context, job models.Job) (int64, error) {
	if err := js.fileRepo.SetObject(ctx, job.JobID, nil, 0, ""); err != nil {
		return 0, fmt.Errorf("jobservice - error in setobject: %w", err)
	}

	if err := js.repo.SetToHash(ctx, job); err != nil {
		js.fileRepo.DeleteFile(ctx, job.JobID)
		return 0, fmt.Errorf("jobservice - error in settohash: %w", err)
	}

	if err := js.repo.SetToList(ctx, job.JobID); err != nil {
		js.fileRepo.DeleteFile(ctx, job.JobID)
		js.repo.DeleteKey(ctx, job.JobID)
		//TODO: методы необходимо логировать
		return 0, fmt.Errorf("jobservice - error in settolist: %w", err)
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
		js.repo.SetToList(ctx, jobId)
		return nil, fmt.Errorf("jobservice - error in getobject: %w", err)
	}

	job, err := js.repo.GetFromHash(ctx, jobId)
	if err != nil {
		js.repo.SetToList(ctx, jobId)
		return nil, fmt.Errorf("jobservice - error in gethashdata: %w", err)
	}

	return job, nil
}