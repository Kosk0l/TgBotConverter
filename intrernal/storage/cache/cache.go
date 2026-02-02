package cache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
	"github.com/redis/go-redis/v9"
)

// Добавить в list(очередь) jobId
func (r *RedisSt) SetToList(ctx context.Context, jobId string) (error) {
	job := fmt.Sprintf("job:%s", jobId)

	err := r.rdb.RPush(ctx, "queue:jobs", job).Err()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		return fmt.Errorf("redis: lpush(list) job %s failed:%w", jobId, err)
	}

	return nil
}


// Добавить в hash параметры запроса
func (r *RedisSt) SetToHash(ctx context.Context, job domains.Job) (error) {
	query := fmt.Sprintf("job:%s", job.JobID)
	
	err := r.rdb.HSet(ctx, query,
		"chat_id", job.ChatID,
		"file_to", string(job.FileTypeTo),
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
	result, err := r.rdb.BLPop(ctx, 0, "queue:jobs").Result()
	if err != nil {
		return "", fmt.Errorf("redis: blpop failed: %w", err)
	}

	if len(result) != 2 {
		return "", fmt.Errorf("redis: blpop(len) failed: %w", err)
	}

	jobId := strings.TrimPrefix(result[1], "job:")
	return jobId, nil
}


// Получить из hash данные запроса
func (r *RedisSt) GetFromHash(ctx context.Context, jobId string) (domains.Job, error) {
	keyQuery := fmt.Sprintf("job:%s", jobId)

	values, err := r.rdb.HGetAll(ctx, keyQuery).Result()
	if err != nil {
		return domains.Job{}, fmt.Errorf("redis - hgetall job %s failed: %w", jobId, err)
	}

	if len(values) == 0 {
		return domains.Job{}, redis.Nil
	}

	chatId, err := strconv.ParseInt(values["chat_id"], 10, 64)
	if err != nil {
		return domains.Job{}, fmt.Errorf("error parse chat_id: %w", err)
	}

	return domains.Job{
		JobID:      jobId,
		ChatID:     chatId,
		FileTypeTo: domains.FileType(values["file_to"]),
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
