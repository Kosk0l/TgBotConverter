package userservice_test

import (
	"context"
	"testing"

	userservice "github.com/Kosk0l/TgBotConverter/internal/Services/userService"
	"github.com/Kosk0l/TgBotConverter/internal/Services/userService/mocks"
	"github.com/Kosk0l/TgBotConverter/internal/domains"
	"github.com/stretchr/testify/assert"
)

func TestGetByIdService_OK(t *testing.T) {
	// Arange:
	ctx := context.Background()
	repo := mocks.NewUserRepository(t)
	service := userservice.NewUserService(repo)
	userId := int64(123)
	user := domains.User{
		ID: userId,
	}

	// Act:
	repo.On("GetById", ctx, userId).Return(user, nil).Once()

	userActual, err := service.GetByIdService(ctx, userId)
	
	// Assert:
	assert.NoError(t, err)
	assert.NotEmpty(t, userActual)
}

func TestCreateUserService_OK(t *testing.T) {
	// Arange:
	ctx := context.Background()
	repo := mocks.NewUserRepository(t)
	service := userservice.NewUserService(repo)

	user := domains.User{
		ID: 123,
	}

	// Act:
	repo.On("CreareUser", ctx, user).Return(nil).Once()

	err := service.CreateUserService(ctx, user)

	// Assert:
	assert.NoError(t, err)
}

func TestUpdateUserService_OK(t *testing.T) {
	// Arange:
	ctx := context.Background()
	repo := mocks.NewUserRepository(t)
	service := userservice.NewUserService(repo)

	user := domains.User{
		ID: 123,
	}

	// Act:
	repo.On("UpdateUser", ctx, user).Return(nil).Once()

	err := service.UpdateUserService(ctx, user)

	// Assert:
	assert.NoError(t, err)
}

func TestUpdateLastSeenService_OK(t *testing.T) {
	// Arange:
	ctx := context.Background()
	repo := mocks.NewUserRepository(t)
	service := userservice.NewUserService(repo)

	userId := int64(123)

	// Act:
	repo.On("UpdateLastSeen", ctx, userId).Return(nil).Once()

	err := service.UpdateLastSeenService(ctx, userId)

	// Assert:
	assert.NoError(t, err)
}

func TestDeleteUserService_OK(t *testing.T) {
	// Arange:
	ctx := context.Background()
	repo := mocks.NewUserRepository(t)
	service := userservice.NewUserService(repo)

	userId := int64(123)

	// Act:
	repo.On("DeleteUser", ctx, userId).Return(nil).Once()

	err := service.DeleteUserService(ctx, userId)
	
	// Assert:
	assert.NoError(t, err)
}