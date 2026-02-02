package userservice

import (
	"context"
	"fmt"
	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
)

// Контенкст создать в хендлере для сервисов // Контекст не должен жить долго
type UserRepository interface {
	GetById(ctx context.Context, userId int64) (domains.User, error)
	CreareUser(ctx context.Context, user domains.User) (error)
	UpdateUser(ctx context.Context, user domains.User) (error)
	UpdateLastSeen(ctx context.Context, userId int64) (error)
	DeleteUser(ctx context.Context, userid int64) (error)
}

type UserService struct {
	repo UserRepository // Объект интерфейса - компановка
}

// Конструктор
func NewUserService(repo UserRepository) (*UserService) {
	return &UserService{
		repo: repo,
	}
}

//====================================================================================================

func (u *UserService) GetByIdService(ctx context.Context, userId int64) (domains.User, error) {
	user, err :=u.repo.GetById(ctx, userId)
	if err != nil {
		return domains.User{}, fmt.Errorf("Bad get by id: %w", err)
	}

	return user, nil
}

func (u *UserService) CreateUserService(ctx context.Context, user domains.User) (error) {
	err := u.repo.CreareUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error in create: %w", err)
	}

	return nil
}

func (u *UserService) UpdateUserService(ctx context.Context, user domains.User) (error) {
	err := u.repo.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("Bad update: %w", err)
	}

	return nil
}

func (u *UserService) UpdateLastSeenService(ctx context.Context, userId int64) (error) {
	err := u.repo.UpdateLastSeen(ctx, userId)
	if err != nil {
		return fmt.Errorf("bad update last seen : %w", err)
	}

	return nil
}

func (u *UserService) DeleteUserService(ctx context.Context, userId int64) (error) {
	err := u.repo.DeleteUser(ctx,userId)
	if err != nil {
		return fmt.Errorf("bad delete: %w", err)
	}

	return nil
}

