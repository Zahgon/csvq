package query

import (
	"github.com/mithrandie/csvq/lib/parser"
)

type AliasMap map[string]string

func (m AliasMap) Add(alias parser.Identifier, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m AliasMap) Get(alias parser.Identifier) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m AliasMap) Clear() { _ = "STUB: not implemented"; return }
