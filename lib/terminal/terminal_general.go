//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris && !windows

package terminal

import (
	"context"
	"io"

	"github.com/mithrandie/csvq/lib/query"

	"golang.org/x/crypto/ssh/terminal"
)

type SSHTerminal struct {
	terminal  *terminal.Terminal
	stdin     int
	origState *terminal.State
	rawState  *terminal.State
	prompt    *Prompt
	tx        *query.Transaction
}

func NewTerminal(ctx context.Context, scope *query.ReferenceScope) (query.VirtualTerminal, error) {
	_ = "STUB: not implemented"
	return *new(query.VirtualTerminal), nil
}

func (t SSHTerminal) Teardown() error { _ = "STUB: not implemented"; return nil }

func (t SSHTerminal) RestoreRawMode() error { _ = "STUB: not implemented"; return nil }

func (t SSHTerminal) RestoreOriginalMode() error { _ = "STUB: not implemented"; return nil }

func (t SSHTerminal) ReadLine() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (t SSHTerminal) Write(s string) error { _ = "STUB: not implemented"; return nil }

func (t SSHTerminal) WriteError(s string) error { _ = "STUB: not implemented"; return nil }

func (t SSHTerminal) SetPrompt(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t SSHTerminal) SetContinuousPrompt(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t SSHTerminal) SaveHistory(s string) error { _ = "STUB: not implemented"; return nil }

func (t SSHTerminal) GetSize() (int, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func (t SSHTerminal) ReloadConfig() error { _ = "STUB: not implemented"; return nil }

type StdIO struct {
	reader io.Reader
	writer io.Writer
}

func (t SSHTerminal) UpdateCompleter() {
	_ = "STUB: not implemented"
	// Do Nothing
	return
}

func (sh *StdIO) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (sh *StdIO) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func NewStdIO(sess *query.Session) *StdIO { _ = "STUB: not implemented"; return nil }
