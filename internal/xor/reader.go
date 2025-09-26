package xor

import (
	"io"
)

type Reader struct {
	reader io.Reader
	key    []byte
	pos    int
}

func NewReader(r io.Reader, key []byte) *Reader {
	return &Reader{reader: r, key: normalizeKey(key), pos: 0}
}

func (r *Reader) Close() error {
	if closer, ok := r.reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	if n > 0 {
		for i := 0; i < n; i++ {
			p[i] ^= r.key[r.pos%len(r.key)]
			r.pos++
		}
	}
	return n, err
}
