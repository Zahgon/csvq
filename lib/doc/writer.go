package doc

import (
	"bytes"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/go-text/color"
)

const DefaultPadding = 1

type Writer struct {
	Flags   *option.Flags
	Palette *color.Palette

	MaxWidth    int
	Padding     int
	Indent      int
	IndentWidth int

	Title1       string
	Title1Effect string
	Title2       string
	Title2Effect string

	buf bytes.Buffer

	subBlock  int
	lineWidth int
	Column    int
}

func NewWriter(screenWidth int, flags *option.Flags, palette *color.Palette) *Writer {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) Clear() { _ = "STUB: not implemented"; return }

func (w *Writer) WriteColorWithoutLineBreak(s string, effect string) {
	_ = "STUB: not implemented"
	return
}

func (w *Writer) WriteColor(s string, effect string) { _ = "STUB: not implemented"; return }

func (w *Writer) write(s string, effect string, withoutLineBreak bool) {
	_ = "STUB: not implemented"
	return
}

func (w *Writer) writeToBuf(s string) { _ = "STUB: not implemented"; return }

func (w *Writer) LeadingSpacesWidth() int { _ = "STUB: not implemented"; return 0 }

func (w *Writer) FitInLine(s string) bool { _ = "STUB: not implemented"; return false }

func (w *Writer) WriteWithoutLineBreak(s string) { _ = "STUB: not implemented"; return }

func (w *Writer) Write(s string) { _ = "STUB: not implemented"; return }

func (w *Writer) WriteWithAutoLineBreak(s string) { _ = "STUB: not implemented"; return }

func (w *Writer) WriteWithAutoLineBreakWithContinueMark(s string) {
	_ = "STUB: not implemented"
	return
}

func (w *Writer) writeWithAutoLineBreak(s string, useContinueMark bool, useBlock bool) {
	_ = "STUB: not implemented"
	return
}

func (w *Writer) WriteSpaces(l int) { _ = "STUB: not implemented"; return }

func (w *Writer) NewLine() { _ = "STUB: not implemented"; return }

func (w *Writer) BeginBlock() { _ = "STUB: not implemented"; return }

func (w *Writer) EndBlock() { _ = "STUB: not implemented"; return }

func (w *Writer) BeginSubBlock() { _ = "STUB: not implemented"; return }

func (w *Writer) EndSubBlock() { _ = "STUB: not implemented"; return }

func (w *Writer) ClearBlock() { _ = "STUB: not implemented"; return }

func (w *Writer) String() string { _ = "STUB: not implemented"; return "" }
