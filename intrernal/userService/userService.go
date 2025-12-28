package userservice

import (
	"context"

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

func (u *UserService) GetByIdService(userId int64) (*models.User, error) {
	

	return &models.User{
		ID: 0,
	}, nil
}

func (u *UserService) CreateUserService(user *models.User) (int64, error) {
	return 0, nil
}

func (u *UserService) UpdateUserService(user *models.User) (error) {
	return nil
}

func (u *UserService) UpdateLastSeenService(userId int64) (error) {
	return nil
}

func (u *UserService) DeleteUserService(iserId int64) (error) {
	return nil
}

