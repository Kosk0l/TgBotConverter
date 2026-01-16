package minio

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

func(m *Minio) SetObject(ctx context.Context, jobId string, r io.Reader, size int64, contentType string) (error){
	_, err := m.client.PutObject(ctx, m.bucket, objectName(jobId), r, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("minio - error set object:%w", err)
	}

	return nil
}

func (m *Minio) GetObject(ctx context.Context, jobId string) (io.Reader, error){
	obj, err := m.client.GetObject(ctx, m.bucket, objectName(jobId), minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("minio - error get object:%w", err)
	}

	if _, err := obj.Stat(); err != nil {
		return nil, fmt.Errorf("minio - object not found:%w", err)
	}

	return obj, nil
}

func (m *Minio) DeleteObject(ctx context.Context, jobId string) (error){
	err := m.client.RemoveObject(ctx, m.bucket, objectName(jobId), minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("minio - error delete object:%w", err)
	}

	return nil
}

func (m *Minio) ExistObject(ctx context.Context, jobId string) (bool, error) {
	_, err := m.client.StatObject(ctx, m.bucket, objectName(jobId), minio.StatObjectOptions{})
	if err == nil {
		return true, nil
	}

	if minio.ToErrorResponse(err).Code == "NoSuchKey" {
		return false, nil
	}

	return false, err
}