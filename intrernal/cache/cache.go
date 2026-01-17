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

// Добавить в list(очередь) jobId
func (r *RedisSt) SetToList(ctx context.Context, jobId string) (error) {
	job := fmt.Sprintf("job:%s", jobId)

	err := r.rdb.LPush(ctx, "queue", job).Err()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		return fmt.Errorf("redis: lpush(list) job %s failed:%w", jobId, err)
	}

	return nil
}


// Добавить в hash параметры запроса
func (r *RedisSt) SetToHash(ctx context.Context, job models.Job) (error) {
	query := fmt.Sprintf("job:%s", job.JobID)
	job.StatusJob = models.InQueue
	
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
		return fmt.Errorf("redis: hset(hash) job %s failed:%w", job.JobID, err)
	}

	return nil
}


// Получить последний JobId из очереди
func (r *RedisSt) GetFromList(ctx context.Context) (string, error) {
	result, err := r.rdb.RPop(ctx, "queue").Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil // очередь пуста
		}
		return "", fmt.Errorf("redis: rpop failed: %w", err)
	}

	jobId := strings.TrimPrefix(result, "job:")
	return jobId, nil
}


// Получить из hash данные запроса
func (r *RedisSt) GetFromHash(ctx context.Context, jobId string) (*models.Job, error) {
	keyQuery := fmt.Sprintf("job:%s", jobId)

	values, err := r.rdb.HGetAll(ctx, keyQuery).Result()
	if err != nil {
		return nil, fmt.Errorf("redis - hgetall job %s failed: %w", jobId, err)
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
		JobID:      jobId,
		UserID:     userId,
		ChatID:     chatId,
		FileTypeIn: values["file_in"],
		FileTypeTo: values["file_to"],
		StatusJob: models.ProcessedNow,
	}, nil
}


// Удалить ключ
func(r *RedisSt) DeleteKey(ctx context.Context, jobId string) (error) {
	query := fmt.Sprintf("job:%s", jobId)
	err := r.rdb.Del(ctx, query).Err()
	if err != nil {
		return fmt.Errorf("redis - failed delete key:%w", err)
	}

	return nil
}


// Вернуть данные в List справа
func (r *RedisSt) SetToListR(ctx context.Context, jobId string) (error) {
	job := fmt.Sprintf("job:%s", jobId)

	err := r.rdb.RPush(ctx, "queue", job).Err()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		return fmt.Errorf("redis: Rpush(list) job %s failed:%w", jobId, err)
	}

	return nil
}