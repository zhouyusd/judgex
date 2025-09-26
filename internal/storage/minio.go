package storage

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

var _ Storage = (*MinioStorage)(nil)

type MinioStorage struct {
	client     *minio.Client
	bucketName string
}

type minioObjectWrapper struct {
	*minio.Object
}

func (mbw *minioObjectWrapper) Size() int64 {
	info, err := mbw.Stat()
	if err != nil {
		return 0
	}
	return info.Size
}

func (ms *MinioStorage) GetObject(ctx context.Context, objectName string) (SizeReadAtSeekCloser, error) {
	obj, err := ms.client.GetObject(ctx, ms.bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &minioObjectWrapper{obj}, nil
}

func (ms *MinioStorage) PutObject(ctx context.Context, objectName string, objectSize int64, r io.Reader) error {
	defer io.Copy(io.Discard, r)
	if _, err := ms.client.PutObject(ctx, ms.bucketName, objectName, r, objectSize, minio.PutObjectOptions{}); err != nil {
		return ErrSaveFileFailed
	}
	return nil
}

func (ms *MinioStorage) RemoveObject(ctx context.Context, objectName string) error {
	return ms.client.RemoveObject(ctx, ms.bucketName, objectName, minio.RemoveObjectOptions{})
}

func NewMinioStorage(client *minio.Client, bucketName, location string) (*MinioStorage, error) {
	if exist, _ := client.BucketExists(context.Background(), bucketName); !exist {
		if err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location}); err != nil {
			return nil, err
		}
	}
	return &MinioStorage{client, bucketName}, nil
}
