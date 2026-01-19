package pendingservice

import (
	"context"

	
)

// Абстракция для cache
type DialogRepository interface {
	SetInquiry(ctx context.Context,) (error)
	GetInquiry(ctx context.Context, fileUrl string) (error)
	DeleteInquiry(ctx context.Context, fileUrl string) (error)
}

// Сервис управления состояниями диалога
type DialogStateService struct {
	pr	DialogRepository
}

// Конструктор
func NewDialogStateService(pr DialogRepository) (*DialogStateService) {
	return &DialogStateService{
		pr: pr,
	}
}

//====================================================================================================

func (p *DialogStateService) SetState() (error) {

	return nil
}

func (p *DialogStateService) GetState() (error) {

	return nil
}