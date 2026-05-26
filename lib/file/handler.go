package file

import (
	"context"
	"os"
	"time"
)

type OpenType int

const (
	ForRead OpenType = iota
	ForCreate
	ForUpdate
)

type Handler struct {
	path string
	fp   *os.File

	openType OpenType

	rlockFile *ControlFile
	lockFile  *ControlFile
	tempFile  *ControlFile

	closed bool
}

func NewHandlerWithoutLock(ctx context.Context, path string, defaultWaitTimeout time.Duration, retryDelay time.Duration) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHandlerForRead(ctx context.Context, path string, defaultWaitTimeout time.Duration, retryDelay time.Duration) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newHandlerForCreate(_ context.Context, path string, _ time.Duration, _ time.Duration) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHandlerForCreate(path string) (*Handler, error) { _ = "STUB: not implemented"; return nil, nil }

func NewHandlerForUpdate(ctx context.Context, path string, defaultWaitTimeout time.Duration, retryDelay time.Duration) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func closeIsolatedHandler(h *Handler, err error) error { _ = "STUB: not implemented"; return nil }

func (h *Handler) Path() string { _ = "STUB: not implemented"; return "" }

func (h *Handler) File() *os.File { _ = "STUB: not implemented"; return nil }

func (h *Handler) FileForUpdate() (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *Handler) close() error { _ = "STUB: not implemented"; return nil }

func (h *Handler) commit() error { _ = "STUB: not implemented"; return nil }

func (h *Handler) closeWithErrors() error { _ = "STUB: not implemented"; return nil }

func (h *Handler) CreateControlFileContext(ctx context.Context, fileType ControlFileType, retryDelay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//RLock

//RLock
