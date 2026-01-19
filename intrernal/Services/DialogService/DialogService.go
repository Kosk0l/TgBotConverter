package Dialogservice

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
type DialogService struct {
	pr	DialogRepository
}

// Конструктор
func NewDialogService(pr DialogRepository) (*DialogService) {
	return &DialogService{
		pr: pr,
	}
}

//====================================================================================================

func (p *DialogService) SetState() (error) {

	return nil
}

func (p *DialogService) GetState() (error) {

	return nil
}