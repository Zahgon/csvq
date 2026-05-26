//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows

package terminal

import (
	"context"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/query"

	"github.com/mithrandie/readline-csvq"
)

type ReadLineTerminal struct {
	terminal  *readline.Instance
	fd        int
	prompt    *Prompt
	env       *option.Environment
	completer *Completer
	tx        *query.Transaction
}

func NewTerminal(ctx context.Context, scope *query.ReferenceScope) (query.VirtualTerminal, error) {
	_ = "STUB: not implemented"
	return *new(query.VirtualTerminal), nil
}

func (t ReadLineTerminal) Teardown() error { _ = "STUB: not implemented"; return nil }

func (t ReadLineTerminal) ReadLine() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (t ReadLineTerminal) Write(s string) error { _ = "STUB: not implemented"; return nil }

func (t ReadLineTerminal) WriteError(s string) error { _ = "STUB: not implemented"; return nil }

func (t ReadLineTerminal) SetPrompt(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t ReadLineTerminal) SetContinuousPrompt(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (t ReadLineTerminal) SaveHistory(s string) error { _ = "STUB: not implemented"; return nil }

func (t ReadLineTerminal) GetSize() (int, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func (t ReadLineTerminal) ReloadConfig() error { _ = "STUB: not implemented"; return nil }

func (t ReadLineTerminal) UpdateCompleter() { _ = "STUB: not implemented"; return }

func (t ReadLineTerminal) setCompleter() { _ = "STUB: not implemented"; return }

func (t ReadLineTerminal) setKillWholeLine() { _ = "STUB: not implemented"; return }

func (t ReadLineTerminal) setViMode() { _ = "STUB: not implemented"; return }

func HistoryFilePath(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }
