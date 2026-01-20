package Dialogservice

import (
	"context"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
)

// Абстракция для cache
type DialogRepository interface {
	SetStateRepo(ctx context.Context, state domains.State) (error)
	GetStateRepo(ctx context.Context, fileUrl string) (*domains.State, error)
	DeleteStateRepo(ctx context.Context, fileUrl string) (error)
}

// Сервис управления состояниями диалога
type DialogService struct {
	dr	DialogRepository
}

// Конструктор
func NewDialogService(dr DialogRepository) (*DialogService) {
	return &DialogService{
		dr: dr,
	}
}

//====================================================================================================

// Создать сосотояние
func (p *DialogService) SetState(ctx context.Context, state domains.State) (error) {


	return nil
}

// Получить состояние
func (p *DialogService) GetState() (error) {

	return nil
}