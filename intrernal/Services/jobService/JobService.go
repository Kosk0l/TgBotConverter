package jobservice

import (
	"context"
	"io"

	"github.com/Kosk0l/TgBotConverter/intrernal/models"
)

// Абстракция для cache
type JobRepository interface {
	SetToList(ctx context.Context, jobId int64) (error)
	SetToHash(ctx context.Context, job models.Job) (error)
	GetFromList(ctx context.Context) (int64, error)
	GetFromHash(ctx context.Context, jobId int64) (*models.Job, error)
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

	return 0, nil
}

// Получить job
func (js *JobService) GetJob(ctx context.Context, jobId int64) (*models.Job, error) {

	return &models.Job{

	}, nil
}
	
// Удалить job
func (js *JobService) DeleteJob(ctx context.Context, jobId int64) (error) {

	return nil
}