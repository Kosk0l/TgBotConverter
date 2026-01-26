package cache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	//"time"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
	"github.com/redis/go-redis/v9"
)

// Добавить состояние
func (r *RedisSt) SetStateRepo(ctx context.Context, state domains.State) (error) {
	keyQuery := "chat:" + strconv.FormatInt(state.ChatId, 10)

	// Проверить id
	if state.ChatId == 0 {
		return errors.New("chatId is required")
	}

	pipe := r.rdb.TxPipeline() // начало пайплайна
	// Добавить hash
	pipe.HSet(ctx, keyQuery,
		"user_id", state.UserId,
		"step", string(state.Step),
		"file_url", state.FileURL,
		"file_name", state.FileName,
		"size", state.Size,
		"content_type", state.ContentType,
	)
	// добавить таймер
	pipe.Expire(ctx, keyQuery, 10*time.Minute).Err() 
	_, err := pipe.Exec(ctx) // конец пайплайна
	if err != nil {
		return fmt.Errorf("redis - error pipline: %w", err)
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

	// Конвертация size
	size, err := strconv.ParseInt(values["size"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parse size: %w", err)
	}

	// конвертация userId
	userId, err := strconv.ParseInt(values["user_id"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parse user_id: %w", err)
	}

	return &domains.State{
		ChatId: chatId,
		UserId: userId,
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

	// Удалить ключ
	err := r.rdb.Del(ctx, keyQuery).Err()
	if err != nil {
		return fmt.Errorf("redis - failed delete key:%w", err)
	}

	return nil
}
