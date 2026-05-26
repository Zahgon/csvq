package query

import (
	"context"
	"errors"
	"sync"

	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/ternary"
)

var errUndeclaredCursor = errors.New("undeclared cursor")
var errPseudoCursor = errors.New("unpermmitted pseudo cursor usage")
var errCursorClosed = errors.New("cursor cloased")

type CursorMap struct {
	*SyncMap
}

func NewCursorMap() CursorMap { _ = "STUB: not implemented"; return *new(CursorMap) }

func (m CursorMap) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (m CursorMap) Store(name string, val *Cursor) { _ = "STUB: not implemented"; return }

func (m CursorMap) LoadDirect(name string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m CursorMap) Load(name string) (*Cursor, bool) { _ = "STUB: not implemented"; return nil, false }

func (m CursorMap) Delete(name string) { _ = "STUB: not implemented"; return }

func (m CursorMap) Exists(name string) bool { _ = "STUB: not implemented"; return false }

func (m CursorMap) Declare(expr parser.CursorDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}

func (m CursorMap) AddPseudoCursor(name parser.Identifier, values []value.Primary) error {
	_ = "STUB: not implemented"
	return nil
}

func (m CursorMap) Dispose(name parser.Identifier) error { _ = "STUB: not implemented"; return nil }

func (m CursorMap) Open(ctx context.Context, scope *ReferenceScope, name parser.Identifier, values []parser.ReplaceValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (m CursorMap) Close(name parser.Identifier) error { _ = "STUB: not implemented"; return nil }

func (m CursorMap) Fetch(name parser.Identifier, position int, number int) ([]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m CursorMap) IsOpen(name parser.Identifier) (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}

func (m CursorMap) IsInRange(name parser.Identifier) (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}

func (m CursorMap) Count(name parser.Identifier) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type Cursor struct {
	Name      string
	query     parser.SelectQuery
	statement parser.Identifier
	view      *View
	index     int
	fetched   bool

	isPseudo bool

	mtx *sync.Mutex
}

func NewCursor(e parser.CursorDeclaration) *Cursor { _ = "STUB: not implemented"; return nil }

func NewPseudoCursor(name string, values []value.Primary) *Cursor {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cursor) Open(ctx context.Context, scope *ReferenceScope, name parser.Identifier, values []parser.ReplaceValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cursor) Close(name parser.Identifier) error { _ = "STUB: not implemented"; return nil }

func (c *Cursor) Fetch(name parser.Identifier, position int, number int) ([]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NEXT

func (c *Cursor) IsOpen() ternary.Value { _ = "STUB: not implemented"; return *new(ternary.Value) }

func (c *Cursor) IsInRange() (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}

func (c *Cursor) Count() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Cursor) Pointer() (int, error) { _ = "STUB: not implemented"; return 0, nil }
