package query

import (
	"bytes"
	"context"
	"io"
	"os"
	"sync"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
)

var (
	screenFd                = os.Stdin.Fd()
	stdin    io.ReadCloser  = os.Stdin
	stdout   io.WriteCloser = os.Stdout
	stderr   io.WriteCloser = os.Stderr
)

const DefaultScreenWidth = 75

func isNamedPipe(fp *os.File) bool { _ = "STUB: not implemented"; return false }

type Discard struct {
}

func NewDiscard() *Discard { _ = "STUB: not implemented"; return nil }

func (d Discard) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (d Discard) Close() error { _ = "STUB: not implemented"; return nil }

type Input struct {
	reader io.Reader
}

func NewInput(r io.Reader) *Input { _ = "STUB: not implemented"; return nil }

func (r *Input) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Input) Close() error { _ = "STUB: not implemented"; return nil }

type Output struct {
	bytes.Buffer
}

func NewOutput() *Output { _ = "STUB: not implemented"; return nil }

func (w *Output) Close() error { _ = "STUB: not implemented"; return nil }

type StdinLocker struct {
	mtx      *sync.Mutex
	locked   bool
	rlockCnt int32
}

func NewStdinLocker() *StdinLocker { _ = "STUB: not implemented"; return nil }

func (cl *StdinLocker) Lock() error { _ = "STUB: not implemented"; return nil }

func (cl *StdinLocker) LockContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *StdinLocker) Unlock() (err error) { _ = "STUB: not implemented"; return nil }

func (cl *StdinLocker) RLock() error { _ = "STUB: not implemented"; return nil }

func (cl *StdinLocker) RLockContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *StdinLocker) RUnlock() (err error) { _ = "STUB: not implemented"; return nil }

func (cl *StdinLocker) tryLock() bool { _ = "STUB: not implemented"; return false }

func (cl *StdinLocker) tryRLock() bool { _ = "STUB: not implemented"; return false }

func (cl *StdinLocker) lockContext(ctx context.Context, fn func() bool) error {
	_ = "STUB: not implemented"
	return nil
}

// try again

type Session struct {
	screenFd uintptr
	stdin    io.ReadCloser
	stdout   io.WriteCloser
	stderr   io.WriteCloser
	outFile  io.Writer
	terminal VirtualTerminal

	CanReadStdin    bool
	CanOutputToPipe bool

	stdinViewMap ViewMap
	stdinLocker  *StdinLocker

	mtx *sync.Mutex
}

func NewSession() *Session { _ = "STUB: not implemented"; return nil }

func (sess *Session) ScreenFd() uintptr { _ = "STUB: not implemented"; return 0 }

func (sess *Session) Stdin() io.ReadCloser { _ = "STUB: not implemented"; return *new(io.ReadCloser) }

func (sess *Session) Stdout() io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func (sess *Session) Stderr() io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func (sess *Session) OutFile() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (sess *Session) Terminal() VirtualTerminal {
	_ = "STUB: not implemented"
	return *new(VirtualTerminal)
}

func (sess *Session) ScreenWidth() int { _ = "STUB: not implemented"; return 0 }

func (sess *Session) SetStdin(r io.ReadCloser) error { _ = "STUB: not implemented"; return nil }

func (sess *Session) SetStdinContext(ctx context.Context, r io.ReadCloser) error {
	_ = "STUB: not implemented"
	return nil
}

func (sess *Session) SetStdout(w io.WriteCloser) { _ = "STUB: not implemented"; return }

func (sess *Session) SetStderr(w io.WriteCloser) { _ = "STUB: not implemented"; return }

func (sess *Session) SetOutFile(w io.Writer) { _ = "STUB: not implemented"; return }

func (sess *Session) SetTerminal(t VirtualTerminal) { _ = "STUB: not implemented"; return }

func (sess *Session) GetStdinView(ctx context.Context, flags *option.Flags, fileInfo *FileInfo, expr parser.Stdin) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sess *Session) updateStdinView(view *View) { _ = "STUB: not implemented"; return }

func (sess *Session) WriteToStdout(s string) (err error) { _ = "STUB: not implemented"; return nil }

func (sess *Session) WriteToStdoutWithLineBreak(s string) error {
	_ = "STUB: not implemented"
	return nil
}

func (sess *Session) WriteToStderr(s string) (err error) { _ = "STUB: not implemented"; return nil }

func (sess *Session) WriteToStderrWithLineBreak(s string) error {
	_ = "STUB: not implemented"
	return nil
}
