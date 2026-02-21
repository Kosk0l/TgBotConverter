package app

import (
	"context"
	"log/slog"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/Kosk0l/TgBotConverter/internal/lib/logger"
	"github.com/Kosk0l/TgBotConverter/internal/storage/cache"
	"github.com/Kosk0l/TgBotConverter/internal/storage/minio"
	"github.com/Kosk0l/TgBotConverter/internal/storage/postgres"
)

type Infrastructure struct {
	Logger 	*slog.Logger
	DB 		*postgres.Postgres
	Cache 	*cache.RedisSt
	Minio 	*minio.Minio
}

func initInfrastructure(ctx context.Context, cfg config.Config) (*Infrastructure, error) {
    // Объект логгера
    logger := logger.NewLogger(cfg)


	// объект постгреса 
    dsn := config.LoadDsn(cfg)
    pool, err := postgres.NewPostgres(ctx, dsn)
    if err != nil {
        return nil, err
    }

    // объект редиса
    cache, err := cache.NewRedis(ctx, cfg)
    if err != nil {
        return nil, err
    }

    // объект минио
    minioClient, err := minio.NewMinio(ctx, cfg, "files")
    if err != nil {
        return nil, err
    }

    return &Infrastructure{
        Logger: logger,
        DB:     pool,
        Cache:  cache,
        Minio:  minioClient,
    }, nil
}