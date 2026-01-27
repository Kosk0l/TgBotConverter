package minio

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/minio/minio-go/v7"
)

// Добавить объект
func(m *Minio) SetObject(ctx context.Context, jobId string, fileUrl string, size int64, contentType string) (error){

	Reader, err := http.Get(fileUrl)
	if err != nil {
		return fmt.Errorf("minio - error get file url: %w", err)
	}
	defer Reader.Body.Close()

	if _, err := m.client.PutObject(ctx, m.bucket, objectName(jobId), Reader.Body, size, minio.PutObjectOptions{
		ContentType: contentType,
	}); err != nil {
		return fmt.Errorf("minio - error set object:%w", err)
	}

	return nil
}

// Получить объект
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

// Удалить Объект
func (m *Minio) DeleteObject(ctx context.Context, jobId string) (error){
	err := m.client.RemoveObject(ctx, m.bucket, objectName(jobId), minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("minio - error delete object:%w", err)
	}

	return nil
}

// Проверить наличие объекта
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