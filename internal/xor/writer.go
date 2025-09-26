package xor

import "io"

type Writer struct {
	writer io.Writer
	key    []byte
	pos    int
}

func NewWriter(w io.Writer, key []byte) *Writer {
	return &Writer{writer: w, key: normalizeKey(key), pos: 0}
}

func (w *Writer) Close() error {
	if w.writer == nil {
		return nil
	}
	if closer, ok := w.writer.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if w.writer == nil {
		return 0, io.ErrClosedPipe
	}
	encrypted := make([]byte, len(p))
	for i, b := range p {
		encrypted[i] = b ^ w.key[w.pos%len(w.key)]
		w.pos++
	}
	return w.writer.Write(encrypted)
}
