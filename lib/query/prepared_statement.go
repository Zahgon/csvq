package query

import (
	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
)

type PreparedStatementMap struct {
	*SyncMap
}

func NewPreparedStatementMap() PreparedStatementMap {
	_ = "STUB: not implemented"
	return *new(PreparedStatementMap)
}

func (m PreparedStatementMap) Store(name string, statement *PreparedStatement) {
	_ = "STUB: not implemented"
	return
}

func (m PreparedStatementMap) LoadDirect(name string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m PreparedStatementMap) Load(name string) (*PreparedStatement, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m PreparedStatementMap) Delete(name string) { _ = "STUB: not implemented"; return }

func (m PreparedStatementMap) Exists(name string) bool { _ = "STUB: not implemented"; return false }

func (m PreparedStatementMap) Prepare(flags *option.Flags, expr parser.StatementPreparation) error {
	_ = "STUB: not implemented"
	return nil
}

func (m PreparedStatementMap) Get(name parser.Identifier) (*PreparedStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m PreparedStatementMap) Dispose(expr parser.DisposeStatement) error {
	_ = "STUB: not implemented"
	return nil
}

type PreparedStatement struct {
	Name            string
	StatementString string
	Statements      []parser.Statement
	HolderNumber    int
}

func NewPreparedStatement(flags *option.Flags, expr parser.StatementPreparation) (*PreparedStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ReplaceValues struct {
	Values []parser.QueryExpression
	Names  map[string]int
}

func NewReplaceValues(replace []parser.ReplaceValue) *ReplaceValues {
	_ = "STUB: not implemented"
	return nil
}
