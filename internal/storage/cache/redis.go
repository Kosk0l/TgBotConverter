package cache

import (
	"context"
	"time"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/redis/go-redis/v9"
)

type RedisSt struct {
	rdb *redis.Client
}

func NewRedis(ctx context.Context, cfg config.Config) (*RedisSt, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Re.Addr,
		Password: cfg.Re.Password,
		DB: cfg.Re.DB,
	})

	ctxRedis, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := rdb.Ping(ctxRedis).Err(); err != nil {
		return nil, err
	}

	return &RedisSt{
		rdb: rdb,
	}, nil
}