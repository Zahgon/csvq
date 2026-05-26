package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

type UserDefinedFunctionMap struct {
	*SyncMap
}

func NewUserDefinedFunctionMap() UserDefinedFunctionMap {
	_ = "STUB: not implemented"
	return *new(UserDefinedFunctionMap)
}

func (m UserDefinedFunctionMap) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (m UserDefinedFunctionMap) Store(name string, val *UserDefinedFunction) {
	_ = "STUB: not implemented"
	return
}

func (m UserDefinedFunctionMap) LoadDirect(name string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m UserDefinedFunctionMap) Load(name string) (*UserDefinedFunction, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m UserDefinedFunctionMap) Delete(name string) { _ = "STUB: not implemented"; return }

func (m UserDefinedFunctionMap) Exists(name string) bool { _ = "STUB: not implemented"; return false }

func (m UserDefinedFunctionMap) Declare(expr parser.FunctionDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}

func (m UserDefinedFunctionMap) DeclareAggregate(expr parser.AggregateDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}

func (m UserDefinedFunctionMap) parseParameters(parameters []parser.VariableAssignment) ([]parser.Variable, map[string]parser.QueryExpression, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

func (m UserDefinedFunctionMap) CheckDuplicate(name parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (m UserDefinedFunctionMap) Get(name string) (*UserDefinedFunction, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m UserDefinedFunctionMap) Dispose(name parser.Identifier) bool {
	_ = "STUB: not implemented"
	return false
}

type UserDefinedFunction struct {
	Name         parser.Identifier
	Statements   []parser.Statement
	Parameters   []parser.Variable
	Defaults     map[string]parser.QueryExpression
	RequiredArgs int

	IsAggregate bool
	Cursor      parser.Identifier // For Aggregate Functions
}

func (fn *UserDefinedFunction) Execute(ctx context.Context, scope *ReferenceScope, args []value.Primary) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func (fn *UserDefinedFunction) ExecuteAggregate(ctx context.Context, scope *ReferenceScope, values []value.Primary, args []value.Primary) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func (fn *UserDefinedFunction) CheckArgsLen(expr parser.QueryExpression, name string, argsLen int) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn *UserDefinedFunction) execute(ctx context.Context, scope *ReferenceScope, args []value.Primary) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}
