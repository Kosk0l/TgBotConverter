package cache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
	"github.com/redis/go-redis/v9"
)

// Добавить состояние
func (r *RedisSt) SetStateRepo(ctx context.Context, state domains.State) (error) {
	keyQuery := "chat:" + strconv.FormatInt(state.ChatId, 10)

	err := r.rdb.HSet(ctx, keyQuery, 
		"step", state.Step,
		"file_url", state.FileURL,
		"file_name", state.FileName,
		"size", state.Size,
		"content_type", state.ContentType,
		10*time.Minute,
	).Err()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		return fmt.Errorf("redis - error set state: %w", err)
	}

	return nil
}

// Получить состояние
func (r *RedisSt) GetStateRepo(ctx context.Context, chatId int64) (*domains.State, error) {
	keyQuery := "chat:" + strconv.FormatInt(chatId, 10)

	// Получаем мапу из redis
	values, err := r.rdb.HGetAll(ctx, keyQuery).Result()
	if err != nil {
		return nil, fmt.Errorf("redis - error get state: %w", err)
	}

	// Проверка мапы
	if len(values) == 0 {
		return nil, redis.Nil
	}

	size, err := strconv.ParseInt(values["size"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parse size: %w", err)
	}

	return &domains.State{
		ChatId: chatId,
		Step: domains.Step(values["step"]),
		FileURL: values["file_url"],
		FileName: values["file_name"],
		Size: size,
		ContentType: values["content_type"],
	}, nil
}

// Удалить ключ
func (r *RedisSt) DeleteStateRepo(ctx context.Context, chatId int64) (error) {
	keyQuery := "chat:" + strconv.FormatInt(chatId, 10)

	err := r.rdb.Del(ctx, keyQuery).Err()
	if err != nil {
		return fmt.Errorf("redis - failed delete key:%w", err)
	}

	return nil
}
