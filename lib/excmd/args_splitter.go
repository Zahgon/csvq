package excmd

import (
	"bytes"
)

type ArgsSplitter struct {
	src    []rune
	srcPos int
	text   bytes.Buffer

	err error
}

func (s *ArgsSplitter) Init(src string) *ArgsSplitter { _ = "STUB: not implemented"; return nil }

func (s *ArgsSplitter) Err() error { _ = "STUB: not implemented"; return nil }

func (s *ArgsSplitter) Text() string { _ = "STUB: not implemented"; return "" }

func (s *ArgsSplitter) peek() rune { _ = "STUB: not implemented"; return 0 }

func (s *ArgsSplitter) next() rune { _ = "STUB: not implemented"; return 0 }

func (s *ArgsSplitter) Scan() bool { _ = "STUB: not implemented"; return false }

func (s *ArgsSplitter) scanQuotedVariable(quote rune) { _ = "STUB: not implemented"; return }

func (s *ArgsSplitter) scanQuotedString(quote rune) { _ = "STUB: not implemented"; return }

func (s *ArgsSplitter) scanString() { _ = "STUB: not implemented"; return }

func (s *ArgsSplitter) scanExternalCommand() { _ = "STUB: not implemented"; return }
