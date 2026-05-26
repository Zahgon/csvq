package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/parser"
)

func ParseJoinCondition(join parser.Join, view *View, joinView *View) (parser.QueryExpression, []parser.FieldReference, []parser.FieldReference, error) {
	_ = "STUB: not implemented"
	return *new(parser.QueryExpression), nil, nil, nil
}

func CrossJoin(ctx context.Context, scope *ReferenceScope, view *View, joinView *View) error {
	_ = "STUB: not implemented"
	return nil
}

func InnerJoin(ctx context.Context, scope *ReferenceScope, view *View, joinView *View, condition parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

func OuterJoin(ctx context.Context, scope *ReferenceScope, view *View, joinView *View, condition parser.QueryExpression, direction int) error {
	_ = "STUB: not implemented"
	return nil
}

func CalcMinimumRequired(i1 int, i2 int, defaultMinimumRequired int) int {
	_ = "STUB: not implemented"
	return 0
}
