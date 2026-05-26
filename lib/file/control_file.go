package file

import (
	"context"
	"os"
	"time"
)

type ControlFileType int

const (
	RLock ControlFileType = iota
	Lock
	Temporary
)

var controlFileTypeLit = map[ControlFileType]string{
	RLock:     "read lock",
	Lock:      "lock",
	Temporary: "temporary",
}

func (t ControlFileType) String() string { _ = "STUB: not implemented"; return "" }

type ControlFile struct {
	path string
	fp   *os.File
}

func NewControlFile(path string, fp *os.File) *ControlFile { _ = "STUB: not implemented"; return nil }

func (m *ControlFile) Close() error { _ = "STUB: not implemented"; return nil }

func (m *ControlFile) CloseWithErrors() []error { _ = "STUB: not implemented"; return nil }

func CreateControlFileContext(ctx context.Context, filePath string, fileType ControlFileType, retryDelay time.Duration) (*ControlFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// try again

func tryCreateControlFile(filePath string, fileType ControlFileType) (*ControlFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//RLock

func TryCreateRLockFile(filePath string) (controlFile *ControlFile, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TryCreateLockFile(filePath string) (*ControlFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TryCreateTempFile(filePath string) (*ControlFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
