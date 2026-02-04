package Dialogservice

import (
	"context"
	"fmt"

	"github.com/Kosk0l/TgBotConverter/internal/domains"
)
//go:generate mockery --name=DialogRepository --output=./mocks

// Абстракция для cache
type DialogRepository interface {
	SetStateRepo(ctx context.Context, state domains.State) (error)
	GetStateRepo(ctx context.Context, chatId int64) (domains.State, error)
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
	err := p.dr.SetStateRepo(ctx, state)
	if err != nil {
		return fmt.Errorf("dialogservice - error setstate: %w", err)
	}

	return nil
}

// Получить состояние 
func (p *DialogService) GetState(ctx context.Context, chatId int64) (domains.State, error) {
	state, err := p.dr.GetStateRepo(ctx, chatId)
	if err != nil {
		return domains.State{}, fmt.Errorf("dialogservice - error getstate: %w", err)
	}

	// TODO: потом вынести delete в отдельный метод
	if err := p.dr.DeleteStateRepo(ctx, chatId); err != nil {
		return state, fmt.Errorf("dialogservice - error getstate: %w", err)
	}

	return state, nil
}
