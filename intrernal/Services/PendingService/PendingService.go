package pendingservice

import (
	"context"

	"github.com/Kosk0l/TgBotConverter/intrernal/models"
)

// Абстракция для cache
type PendingRepository interface {
	SetInquiry(ctx context.Context, inq models.Inquiry) (error)
	GetInquiry(ctx context.Context, fileUrl string) (*models.Inquiry, error)
	DeleteInquiry(ctx context.Context, fileUrl string) (error)
}

// Бизнес-логика для запросов
type PendingService struct {
	pr	PendingRepository
}

// Конструктор
func NewPendingService(pr PendingRepository) (*PendingService) {
	return &PendingService{
		pr: pr,
	}
}

//====================================================================================================

