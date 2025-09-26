package storage

import (
	"context"
	"errors"
	"io"
)

var (
	ErrInvalidObjectName = errors.New("invalid object name")
	ErrObjectNotFound    = errors.New("object not found")
	ErrMkdirFailed       = errors.New("mkdir failed")
	ErrCreateFileFailed  = errors.New("create file failed")
	ErrSaveFileFailed    = errors.New("save file failed")
)

type SizeReadAtSeekCloser interface {
	io.Reader
	io.Seeker
	io.Closer
	io.ReaderAt
	Size() int64
}

type Storage interface {
	GetObject(ctx context.Context, objectName string) (SizeReadAtSeekCloser, error)
	PutObject(ctx context.Context, objectName string, objectSize int64, r io.Reader) error
	RemoveObject(ctx context.Context, objectName string) error
}
