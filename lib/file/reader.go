package file

import (
	"bytes"
	"io"
)

type Reader struct {
	offset    int64
	reader    io.Reader
	headBytes *bytes.Reader
	headLen   int64
}

func NewReader(r io.Reader, headLen int) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) HeadBytes() (io.ReadSeeker, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), nil
}

func (r *Reader) Size() int64 { _ = "STUB: not implemented"; return 0 }
