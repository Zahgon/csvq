package excmd

import (
	"bytes"
)

type ArgumentScanner struct {
	src         []rune
	srcPos      int
	text        bytes.Buffer
	elementType ElementType

	err error
}

func (s *ArgumentScanner) Init(src string) *ArgumentScanner { _ = "STUB: not implemented"; return nil }

func (s *ArgumentScanner) Err() error { _ = "STUB: not implemented"; return nil }

func (s *ArgumentScanner) Text() string { _ = "STUB: not implemented"; return "" }

func (s *ArgumentScanner) ElementType() ElementType {
	_ = "STUB: not implemented"
	return *new(ElementType)
}

func (s *ArgumentScanner) peek() rune { _ = "STUB: not implemented"; return 0 }

func (s *ArgumentScanner) next() rune { _ = "STUB: not implemented"; return 0 }

func (s *ArgumentScanner) Scan() bool { _ = "STUB: not implemented"; return false }

func (s *ArgumentScanner) scanQuotedEnvironmentVariable(quote rune) {
	_ = "STUB: not implemented"
	return
}

func (s *ArgumentScanner) scanString() { _ = "STUB: not implemented"; return }

func (s *ArgumentScanner) scanVariable() { _ = "STUB: not implemented"; return }

func (s *ArgumentScanner) isVariableRune(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *ArgumentScanner) scanCsvqExpression() { _ = "STUB: not implemented"; return }
