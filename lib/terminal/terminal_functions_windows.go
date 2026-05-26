//go:build windows

package terminal

import (
	"io"
)

func GetStdinForREPL() io.ReadCloser { _ = "STUB: not implemented"; return *new(io.ReadCloser) }
