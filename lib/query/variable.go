package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

type VariableMap struct {
	*SyncMap
}

func NewVariableMap() VariableMap { _ = "STUB: not implemented"; return *new(VariableMap) }

func (m VariableMap) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (m VariableMap) Store(name string, val value.Primary) { _ = "STUB: not implemented"; return }

func (m VariableMap) LoadDirect(name string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m VariableMap) Load(name string) (value.Primary, bool) {
	_ = "STUB: not implemented"
	return *new(value.Primary), false
}

func (m VariableMap) Delete(name string) { _ = "STUB: not implemented"; return }

func (m VariableMap) Exists(name string) bool { _ = "STUB: not implemented"; return false }

func (m VariableMap) Add(variable parser.Variable, val value.Primary) error {
	_ = "STUB: not implemented"
	return nil
}

func (m VariableMap) Set(variable parser.Variable, val value.Primary) bool {
	_ = "STUB: not implemented"
	return false
}

func (m VariableMap) Get(variable parser.Variable) (value.Primary, bool) {
	_ = "STUB: not implemented"
	return *new(value.Primary), false
}

func (m VariableMap) Dispose(variable parser.Variable) bool {
	_ = "STUB: not implemented"
	return false
}

func (m VariableMap) Declare(ctx context.Context, scope *ReferenceScope, declaration parser.VariableDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}
