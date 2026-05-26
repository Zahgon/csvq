package query

import (
	"bytes"

	"github.com/mithrandie/csvq/lib/value"
)

const EOF = -1

type StringFormatter struct {
	format    []rune
	formatPos int
	offset    int
	values    []value.Primary

	buf bytes.Buffer

	err error
}

func NewStringFormatter() *StringFormatter { _ = "STUB: not implemented"; return nil }

func (f *StringFormatter) Format(format string, values []value.Primary) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *StringFormatter) runes() []rune { _ = "STUB: not implemented"; return nil }

func (f *StringFormatter) literal() string { _ = "STUB: not implemented"; return "" }

func (f *StringFormatter) integer() int { _ = "STUB: not implemented"; return 0 }

func (f *StringFormatter) peek() rune { _ = "STUB: not implemented"; return 0 }

func (f *StringFormatter) next() rune { _ = "STUB: not implemented"; return 0 }

func (f *StringFormatter) isFlag(ch rune) bool { _ = "STUB: not implemented"; return false }

func (f *StringFormatter) isDecimal(ch rune) bool { _ = "STUB: not implemented"; return false }

func (f *StringFormatter) scanDecimal() { _ = "STUB: not implemented"; return }

func (f *StringFormatter) numericSign(flag rune, minus bool) string {
	_ = "STUB: not implemented"
	return ""
}
