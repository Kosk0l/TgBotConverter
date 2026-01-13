package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/Kosk0l/TgBotConverter/intrernal/models"
	"github.com/jackc/pgx/v5"
)

//====================================================================================================
// TODO: тестирование storage

// Взять из БД пользователя по id
func (p *Postgres) GetById(ctx context.Context, userId int64) (*models.User, error) {
	query := `
		SELECT userName, firstName, lastName, createdAt, lastSeenAt
		FROM users WHERE id = $1;
	`

	var user models.User
	err := p.pool.QueryRow(ctx, query, userId).Scan(
		&user.UserName, 
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.LastSeen,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found with id %d", userId)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user.ID = userId
	return &user, nil
}


// Создать пользователя
func (p *Postgres) CreareUser(ctx context.Context, user *models.User) (error){
	query := `
		INSERT INTO users (id, userName, firstName, lastName, createdAt, lastSeenAt) VALUES
		($1, $2, $3, $4, $5, $6);
	`
	t := time.Now()
	cmd, err := p.pool.Exec(ctx, query, user.ID, user.UserName, user.FirstName, user.LastName, t, t)
	if err != nil {
		return fmt.Errorf("error in create user %v", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("error in commandTag create")
	}

	return nil
}


// Обновить Пользователя
func (p *Postgres) UpdateUser(ctx context.Context, user *models.User) (error){
	query := `
		UPDATE users
		SET userName = $1, firstName = $2, lastName = $3
		WHERE id = $4;
	`

	cmd, err := p.pool.Exec(ctx, query, user.UserName, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return fmt.Errorf("error in update user %v", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("error in commandTag update")
	}

	return nil
}


// Обновить последний вход
func (p *Postgres) UpdateLastSeen(ctx context.Context, userId int64) (error) {
	query := `
		UPDATE users 
		SET  lastSeenAt = $1
		WHERE id = $2;
	`

	NewSeen := time.Now()
	cmd, err := p.pool.Exec(ctx, query, NewSeen, userId)
	if err != nil {
		return fmt.Errorf("error in update seen user %v", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("error in commandTag update seen")
	}

	return nil
}


// Удалить пользователя
func (p *Postgres) DeleteUser(ctx context.Context, userid int64) (error) {
	query := `
		DELETE FROM users WHERE id = $1;
	`

	cmd, err := p.pool.Exec(ctx, query, userid)
	if err != nil {
		return fmt.Errorf("error in update seen user %v", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("error in commandTag update seen")
	}

	return nil
}