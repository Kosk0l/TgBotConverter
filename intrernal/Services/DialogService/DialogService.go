package Dialogservice

import (
	"context"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
)

// Абстракция для cache
type DialogRepository interface {
	SetStateRepo(ctx context.Context, state domains.State) (error)
	GetStateRepo(ctx context.Context, chatId int64) (*domains.State, error)
	DeleteStateRepo(ctx context.Context, chatId int64) (error)
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

// Создать состояние
func (p *DialogService) SetState(ctx context.Context, state domains.State) (error) {


	return nil
}

// Получить состояние
func (p *DialogService) GetState(ctx context.Context, chatId int64) (*domains.State, error) {
	//TODO: по chatId получить состояние

	return &domains.State{

	}, nil
}