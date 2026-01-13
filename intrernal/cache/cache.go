package cache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Kosk0l/TgBotConverter/intrernal/models"
	"github.com/redis/go-redis/v9"
)

//TODO: тестирование redis

// Добавить в list(очередь) jobId
func (r *RedisSt) SetToList(ctx context.Context, jobId int64) (error) {
	job := fmt.Sprintf("job:%d", jobId)

	err := r.rdb.LPush(ctx, "queue", job).Err()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		return fmt.Errorf("redis: lpush(list) job %d failed:%w", jobId, err)
	}

	return nil
}


// Добавить в hash параметры запроса
func (r *RedisSt) SetToHash(ctx context.Context, job models.Job) (error) {
	query := fmt.Sprintf("job:%d", job.JobID)

	err := r.rdb.HSet(ctx, query,
		"user_id", job.UserID,
		"chat_id", job.ChatID,
		"file_in", job.FileTypeIn,
		"file_to", job.FileTypeTo,
		"status", job.StatusJob,
	).Err()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		return fmt.Errorf("redis: hset(hash) job %d failed:%w", job.JobID, err)
	}

	return nil
}


// Получить последний JobId из очереди
func(r *RedisSt) GetFromList(ctx context.Context) (int64, error) {
	result, err := r.rdb.RPop(ctx, "queue").Result()
	if err != nil {
		return 0, fmt.Errorf("redis: failed get job from list: %w", err)
	}

	jobId, err := strconv.ParseInt(strings.TrimPrefix(result, "job:"), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed parse job_id: %w", err)
	}

	return jobId, nil
}


// Получить из hash данные запроса
func (r *RedisSt) GetFromHash(ctx context.Context, jobId int64) (*models.Job, error) {
	keyQuery := fmt.Sprintf("job:%d", jobId)

	values, err := r.rdb.HGetAll(ctx, keyQuery).Result()
	if err != nil {
		return nil, fmt.Errorf("redis: failed hgetall job %d: %w", jobId, err)
	}

	if len(values) == 0 {
		return nil, redis.Nil
	}

	userId, err := strconv.ParseInt(values["user_id"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parse user_id: %w", err)
	}

	chatId, err := strconv.ParseInt(values["chat_id"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parse chat_id: %w", err)
	}

	return &models.Job{
		JobID: jobId,
		UserID: userId,
		ChatID: chatId,
		FileTypeIn: values["file_in"],
		FileTypeTo: values["file_to"],
		StatusJob: values["status"],
	}, nil
}
