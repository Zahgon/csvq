package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/parser"
)

type InlineTableMap map[string]*View

func (it InlineTableMap) Set(ctx context.Context, scope *ReferenceScope, inlineTable parser.InlineTable) error {
	_ = "STUB: not implemented"
	return nil
}

func (it InlineTableMap) Exists(name parser.Identifier) bool {
	_ = "STUB: not implemented"
	return false
}

func (it InlineTableMap) Store(name parser.Identifier, view *View) error {
	_ = "STUB: not implemented"
	return nil
}

func (it InlineTableMap) Get(name parser.Identifier) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (it InlineTableMap) Clear() { _ = "STUB: not implemented"; return }
