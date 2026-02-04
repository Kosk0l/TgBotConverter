package Dialogservice_test

import (
	"context"
	"errors"
	"testing"

	Dialogservice "github.com/Kosk0l/TgBotConverter/internal/Services/DialogService"
	"github.com/Kosk0l/TgBotConverter/internal/Services/DialogService/mocks"
	"github.com/Kosk0l/TgBotConverter/internal/domains"
	"github.com/stretchr/testify/assert"
)

func TestSetState_OK(t *testing.T) {
	// Arange:
	ctx := context.Background()

	repo := mocks.NewDialogRepository(t)
	service := Dialogservice.NewDialogService(repo)

	state := domains.State{
		ChatId: 1,
		Step: domains.WaitingTargetType,

		FileURL: "qwerty",
		FileName: "qwerty",
		Size: 1000,
		ContentType: "qwerty",
	}
	
	// Act:
	repo.On("SetStateRepo", ctx, state).Return(nil).Once()

	err := service.SetState(ctx, state)

	// Assert:
	assert.NoError(t, err)
}

func TestSetState_Error(t *testing.T) {
	// Arange:
	ctx := context.Background()

	repo := mocks.NewDialogRepository(t)
	service := Dialogservice.NewDialogService(repo)

	state := domains.State{ChatId: 1}
	repoErr := errors.New("redis error")

	// Act:
	repo.On("SetStateRepo", ctx, state).Return(repoErr).Once()
	err := service.SetState(ctx, state)

	// Assert:
	assert.Error(t, err)
	assert.ErrorIs(t, err, repoErr) // проверка цепочки ошибок
}

func TestGetState_OK(t *testing.T) {
	// Arange:
	ctx := context.Background()
	
	repo := mocks.NewDialogRepository(t)
	service := Dialogservice.NewDialogService(repo)

	var chatId int64 = 42
	state := domains.State{
		ChatId: 1,
		Step: domains.WaitingTargetType,

		FileURL: "qwerty",
		FileName: "qwerty",
		Size: 1000,
		ContentType: "qwerty",
	}
	
	// Act:
	repo.On("GetStateRepo", ctx, chatId).Return(state, nil).Once()
	repo.On("DeleteStateRepo", ctx, chatId).Return(nil).Once()

	resultstate, err := service.GetState(ctx, chatId)

	// Assert:
	assert.NoError(t, err)
	assert.Equal(t, state, resultstate)
}

