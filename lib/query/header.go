package query

import (
	"errors"

	"github.com/mithrandie/csvq/lib/parser"
)

const InternalIdColumn = "@__internal_id"

type HeaderField struct {
	View         string
	Identifier   string
	Column       string
	Aliases      []string
	Number       int
	IsFromTable  bool
	IsJoinColumn bool
	IsGroupKey   bool
}

var errFieldAmbiguous = errors.New("field ambiguous")
var errFieldNotExist = errors.New("field not exists")

type Header []HeaderField

func NewHeaderWithId(view string, words []string) Header {
	_ = "STUB: not implemented"
	return *new(Header)
}

func NewHeader(view string, words []string) Header { _ = "STUB: not implemented"; return *new(Header) }

func NewHeaderWithAutofill(view string, words []string) Header {
	_ = "STUB: not implemented"
	return *new(Header)
}

func NewEmptyHeader(len int) Header { _ = "STUB: not implemented"; return *new(Header) }

func AddHeaderField(h Header, identifier string, column string, alias string) (header Header, index int) {
	_ = "STUB: not implemented"
	return *new(Header), 0
}

func (h Header) Len() int { _ = "STUB: not implemented"; return 0 }

func (h Header) TableColumns() []parser.QueryExpression { _ = "STUB: not implemented"; return nil }

func (h Header) TableColumnNames() []string { _ = "STUB: not implemented"; return nil }

func (h Header) ContainsObject(obj parser.QueryExpression) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (h Header) SearchIndex(fieldRef parser.QueryExpression) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h Header) FieldNumberIndex(number parser.ColumnNumber) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h Header) FieldIndex(fieldRef parser.FieldReference) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h Header) ContainsInternalId(viewName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h Header) Update(reference string, fields []parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

func (h Header) Merge(h2 Header) Header { _ = "STUB: not implemented"; return *new(Header) }

func (h Header) Copy() Header { _ = "STUB: not implemented"; return *new(Header) }
