package testcase

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/minio/minio-go/v7"
)

var (
	_ Storage = (*LocalStorage)(nil)
	_ Storage = (*MinioStorage)(nil)
)

var (
	ErrInvalidObjectName = errors.New("invalid object name")
	ErrObjectNotFound    = errors.New("object not found")
	ErrCalcMD5Failed     = errors.New("calc md5 failed")
	ErrMkdirFailed       = errors.New("mkdir failed")
	ErrResetSeekFailed   = errors.New("reset seek failed")
	ErrCreateFileFailed  = errors.New("create file failed")
	ErrSaveFileFailed    = errors.New("save file failed")
)

type ReadAtSeekCloser interface {
	io.Reader
	io.Seeker
	io.Closer
	io.ReaderAt
}

type Storage interface {
	GetObject(ctx context.Context, objectName string) (ReadAtSeekCloser, error)
	PutObject(ctx context.Context, src io.ReadSeekCloser, objectSize int64) error
	RemoveObject(ctx context.Context, objectName string) error
}

type MinioStorage struct {
	client     *minio.Client
	bucketName string
}

func (ms *MinioStorage) GetObject(ctx context.Context, objectName string) (ReadAtSeekCloser, error) {
	if len(objectName) < 6 {
		return nil, ErrInvalidObjectName
	}
	if filepath.Ext(objectName) != ".zip" {
		objectName += ".zip"
	}
	return ms.client.GetObject(ctx, ms.bucketName, objectName, minio.GetObjectOptions{})
}

func (ms *MinioStorage) PutObject(ctx context.Context, src io.ReadSeekCloser, objectSize int64) error {
	defer src.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		return ErrCalcMD5Failed
	}
	objectName := fmt.Sprintf("%x", hash.Sum(nil)) + ".zip"
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return ErrResetSeekFailed
	}
	if _, err := ms.client.PutObject(ctx, ms.bucketName, objectName, src, objectSize, minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"MD5Sum":     objectName[:32],
			"UploadTime": time.Now().Format(time.RFC3339),
		},
		ContentType: "application/zip",
	}); err != nil {
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

type LocalStorage struct {
	baseDir string
}

func (ls *LocalStorage) GetObject(_ context.Context, objectName string) (ReadAtSeekCloser, error) {
	if len(objectName) < 6 {
		return nil, ErrInvalidObjectName
	}
	if filepath.Ext(objectName) != ".zip" {
		objectName += ".zip"
	}
	dst := filepath.Join(ls.baseDir, objectName[:2], objectName[2:4], objectName)
	if !fileutil.IsExist(dst) {
		return nil, ErrObjectNotFound
	}
	return os.Open(dst)
}

func (ls *LocalStorage) PutObject(_ context.Context, src io.ReadSeekCloser, _ int64) error {
	defer src.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		return ErrCalcMD5Failed
	}
	objectName := fmt.Sprintf("%x", hash.Sum(nil)) + ".zip"
	dst := filepath.Join(ls.baseDir, objectName[:2], objectName[2:4], objectName)
	if fileutil.IsExist(dst) {
		return nil // 已存在，忽略
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return ErrMkdirFailed
	}
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return ErrResetSeekFailed
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		return ErrCreateFileFailed
	}
	defer dstFile.Close()
	if _, err = io.Copy(dstFile, src); err != nil {
		_ = os.Remove(dst)
		return ErrSaveFileFailed
	}
	return nil
}

func (ls *LocalStorage) RemoveObject(_ context.Context, objectName string) error {
	dst := filepath.Join(ls.baseDir, objectName[:2], objectName[2:4], objectName)
	if !fileutil.IsExist(dst) {
		return ErrObjectNotFound
	}
	return os.Remove(dst)
}

func NewLocalStorage(baseDir string) (*LocalStorage, error) {
	return &LocalStorage{baseDir}, nil
}
