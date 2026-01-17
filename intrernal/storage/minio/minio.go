package minio

import (
	"context"
	"fmt"
	"time"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// S3 хранилище
type Minio struct {
	client *minio.Client
	bucket string
}

// Конструктор
func NewMinio(ctx context.Context, cfg *config.Config, bucket string) (*Minio, error) {
	client, err := minio.New(cfg.Mi.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Mi.AccessKey, cfg.Mi.SecretKey, ""),
		Secure: cfg.Mi.Secure,
	})
	if err != nil {
		return nil, fmt.Errorf("error - config up minio:%w", err)
	}

	ctxMinio, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	exists, err := client.BucketExists(ctxMinio, bucket)
	if err != nil {
		return nil, fmt.Errorf("error - check bucket: %w", err)
	}
	if !exists {
		if err := client.MakeBucket(ctxMinio, bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("error - create bucket: %w", err)
		}
	}

	return &Minio{
		client: client,
		bucket: bucket,
	}, nil
}

func objectName(jobID string) string {
	return fmt.Sprintf("job_%s", jobID)
}