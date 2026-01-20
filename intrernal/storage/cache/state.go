package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
)

// Добавить состояние
func (r *RedisSt) SetStateRepo(ctx context.Context, state domains.State) (error) {
	keyQuery := ""

	err := r.rdb.Set(ctx, keyQuery, "", 10*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("redis - error set state: %w", err)
	}

	return nil
}

// Получить состояние
func (r *RedisSt) GetStateRepo(ctx context.Context, fileUrl string) (*domains.State, error) {
	keyQuery := ""

	err := r.rdb.Get(ctx, keyQuery).Err()
	if err != nil {
		return nil, fmt.Errorf("redis - error get state: %w", err)
	}

	return nil, nil
}

// Удалить ключ
func (r *RedisSt) DeleteStateRepo(ctx context.Context, fileUrl string) (error) {
	keyQuery := ""

	err := r.rdb.Del(ctx, keyQuery).Err()
	if err != nil {
		return fmt.Errorf("redis - failed delete key:%w", err)
	}


	return nil
}
