package userservice

import (
	"context"
	"fmt"

	"github.com/Kosk0l/TgBotConverter/intrernal/models"
)

type UserRepository interface {
	GetById(ctx context.Context, userId int64) (*models.User, error) //TODO:
	CreareUser(ctx context.Context, user *models.User) (int64, error)
	UpdateUser(ctx context.Context, user *models.User) (error)
	UpdateLastSeen(ctx context.Context, userId int64) (error)
	DeleteUser(ctx context.Context, userid int64) (error)
}

type UserService struct {
	repo UserRepository // Объект интерфейса - компановка
}

// Конструктор
func NewService(repo UserRepository) (*UserService) {
	return &UserService{
		repo: repo,
	}
}

//====================================================================================================

func (u *UserService) GetByIdService(ctx context.Context, userId int64) (*models.User, error) {
	user, err :=u.repo.GetById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("Bad get by id: %v", err)
	}

	return user, nil
}

func (u *UserService) CreateUserService(ctx context.Context, user *models.User) (int64, error) {
	u.repo.CreareUser(ctx, user)
	return 0, nil
}

func (u *UserService) UpdateUserService(ctx context.Context, user *models.User) (error) {
	err := u.repo.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("Bad update: %v", err)
	}

	return nil
}

func (u *UserService) UpdateLastSeenService(ctx context.Context, userId int64) (error) {
	err := u.repo.UpdateLastSeen(ctx, userId)
	if err != nil {
		return fmt.Errorf("bad update last seen : %v", err)
	}

	return nil
}

func (u *UserService) DeleteUserService(ctx context.Context, userId int64) (error) {
	err := u.repo.DeleteUser(ctx,userId)
	if err != nil {
		return fmt.Errorf("bad delete: %v", err)
	}

	return nil
}

