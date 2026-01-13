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
	m *minio.Client
}

func NewMinio(ctx context.Context, cfg *config.Config, bucket string) (*Minio, error) {

	client, err := minio.New("", &minio.Options{
		Creds: credentials.NewStaticV4("", "", ""),
		Secure: false,
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
		m: client,
	}, nil
}