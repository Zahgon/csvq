package query

import (
	"context"
	"errors"

	"github.com/mithrandie/csvq/lib/file"
	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
)

var errTableNotLoaded = errors.New("table not loaded")

type ViewMap struct {
	*SyncMap
}

func NewViewMap() ViewMap { _ = "STUB: not implemented"; return *new(ViewMap) }

func (m ViewMap) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (m ViewMap) Store(identifier string, view *View) { _ = "STUB: not implemented"; return }

func (m ViewMap) LoadDirect(identifier string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m ViewMap) Load(identifier string) (*View, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m ViewMap) Delete(identifier string) { _ = "STUB: not implemented"; return }

func (m ViewMap) Exists(identifier string) bool { _ = "STUB: not implemented"; return false }

func (m ViewMap) Get(identifier string) (*View, error) { _ = "STUB: not implemented"; return nil, nil }

func (m ViewMap) GetWithInternalId(ctx context.Context, identifier string, flags *option.Flags) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m ViewMap) Set(view *View) { _ = "STUB: not implemented"; return }

func (m ViewMap) DisposeTemporaryTable(tablePath parser.QueryExpression) bool {
	_ = "STUB: not implemented"
	return false
}

func (m ViewMap) Dispose(container *file.Container, identifier string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m ViewMap) Clean(container *file.Container) error { _ = "STUB: not implemented"; return nil }

func (m ViewMap) CleanWithErrors(container *file.Container) error {
	_ = "STUB: not implemented"
	return nil
}
