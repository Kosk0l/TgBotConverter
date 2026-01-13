package jobservice

import (
	"context"
	"github.com/Kosk0l/TgBotConverter/intrernal/models"
)

// Инверсия зависимостей
type JobRepository interface {
	SetToList(ctx context.Context, jobId int64) (error)
	SetToHash(ctx context.Context, job models.Job) (error)
	GetFromList(ctx context.Context) (int64, error)
	GetFromHash(ctx context.Context, jobId int64) (*models.Job, error)
}

// Бизнес-логика для работы с запросами
type JobService struct {
	repo JobRepository
}

func NewJobService(repo JobRepository) (*JobService) {
	return &JobService{
		repo: repo,
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