package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/parser"
)

type DataObject struct {
	*parser.BaseExpr
	Raw string
}

func (o DataObject) String() string { _ = "STUB: not implemented"; return "" }

type HttpObject struct {
	*parser.BaseExpr
	URL string
}

func (o HttpObject) String() string { _ = "STUB: not implemented"; return "" }

func ParseTableName(ctx context.Context, scope *ReferenceScope, table parser.Table) (parser.Identifier, error) {
	_ = "STUB: not implemented"
	return *new(parser.Identifier), nil
}

// Do Nothing

func NormalizeTableObject(ctx context.Context, scope *ReferenceScope, tableObject parser.QueryExpression) (parser.QueryExpression, error) {
	_ = "STUB: not implemented"
	return *new(parser.QueryExpression), nil
}

func ConvertTableFunction(ctx context.Context, scope *ReferenceScope, tableFunction parser.TableFunction) (parser.QueryExpression, error) {
	_ = "STUB: not implemented"
	return *new(parser.QueryExpression), nil
}

func ConvertUrlExpr(urlExpr parser.Url) (parser.QueryExpression, error) {
	_ = "STUB: not implemented"
	return *new(parser.QueryExpression), nil
}
