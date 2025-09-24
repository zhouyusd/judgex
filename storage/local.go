package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/fileutil"
)

var _ Storage = (*LocalStorage)(nil)

type LocalStorage struct {
	baseDir string
}

type localFileWrapper struct {
	*os.File
}

func (lfw *localFileWrapper) Size() int64 {
	info, err := lfw.Stat()
	if err != nil || info == nil {
		return 0
	}
	return info.Size()
}

func (ls *LocalStorage) GetObject(_ context.Context, objectName string) (SizeReadAtSeekCloser, error) {
	if len(objectName) < 4 {
		return nil, ErrInvalidObjectName
	}
	dst := filepath.Join(ls.baseDir, objectName[:2], objectName[2:4], objectName)
	if !fileutil.IsExist(dst) {
		return nil, ErrObjectNotFound
	}
	f, err := os.Open(dst)
	if err != nil {
		return nil, err
	}
	return &localFileWrapper{f}, nil
}

func (ls *LocalStorage) PutObject(_ context.Context, objectName string, _ int64, r io.Reader) error {
	defer io.Copy(io.Discard, r) // 确保 r 被完全读取
	if len(objectName) < 4 {
		return ErrInvalidObjectName
	}
	dst := filepath.Join(ls.baseDir, objectName[:2], objectName[2:4], objectName)
	if fileutil.IsExist(dst) {
		return nil // 已存在，忽略
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return ErrMkdirFailed
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		return ErrCreateFileFailed
	}
	defer dstFile.Close()
	if _, err = io.Copy(dstFile, r); err != nil {
		_ = os.Remove(dst)
		return ErrSaveFileFailed
	}
	return nil
}

func (ls *LocalStorage) RemoveObject(_ context.Context, objectName string) error {
	if len(objectName) < 4 {
		return ErrInvalidObjectName
	}
	dst := filepath.Join(ls.baseDir, objectName[:2], objectName[2:4], objectName)
	if !fileutil.IsExist(dst) {
		return ErrObjectNotFound
	}
	return os.Remove(dst)
}

func NewLocalStorage(baseDir string) (*LocalStorage, error) {
	return &LocalStorage{baseDir}, nil
}
