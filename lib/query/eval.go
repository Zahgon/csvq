package query

import (
	"bytes"
	"context"

	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

func Evaluate(ctx context.Context, scope *ReferenceScope, expr parser.QueryExpression) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evaluateSequentialRoutine(ctx context.Context, scope *ReferenceScope, view *View, fn func(*ReferenceScope, int) error, thIdx int, gm *GoroutineTaskManager) {
	_ = "STUB: not implemented"
	return
}

func EvaluateSequentially(ctx context.Context, scope *ReferenceScope, view *View, fn func(*ReferenceScope, int) error) error {
	_ = "STUB: not implemented"
	return nil
}

func evalFieldReference(expr parser.QueryExpression, scope *ReferenceScope) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalArithmetic(ctx context.Context, scope *ReferenceScope, expr parser.Arithmetic) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalUnaryArithmetic(ctx context.Context, scope *ReferenceScope, expr parser.UnaryArithmetic) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalConcat(ctx context.Context, scope *ReferenceScope, expr parser.Concat) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalRowOrSingleValue(ctx context.Context, scope *ReferenceScope, expr parser.QueryExpression) (value.RowValue, value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.RowValue), *new(value.Primary), nil
}

func evalComparison(ctx context.Context, scope *ReferenceScope, expr parser.Comparison) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalIs(ctx context.Context, scope *ReferenceScope, expr parser.Is) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalBetween(ctx context.Context, scope *ReferenceScope, expr parser.Between) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func valuesForRowValueListComparison(ctx context.Context, scope *ReferenceScope, lhs parser.QueryExpression, values parser.QueryExpression) (value.RowValue, []value.RowValue, error) {
	_ = "STUB: not implemented"
	return *new(value.RowValue), nil, nil
}

func evalIn(ctx context.Context, scope *ReferenceScope, expr parser.In) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalAny(ctx context.Context, scope *ReferenceScope, expr parser.Any) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalAll(ctx context.Context, scope *ReferenceScope, expr parser.All) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalLike(ctx context.Context, scope *ReferenceScope, expr parser.Like) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalExists(ctx context.Context, scope *ReferenceScope, expr parser.Exists) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalSubqueryForValue(ctx context.Context, scope *ReferenceScope, expr parser.Subquery) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalFunction(ctx context.Context, scope *ReferenceScope, expr parser.Function) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalAggregateFunction(ctx context.Context, scope *ReferenceScope, expr parser.AggregateFunction) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalListFunction(ctx context.Context, scope *ReferenceScope, expr parser.ListFunction) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

// LISTAGG

func evalAnalyticFunction(scope *ReferenceScope, expr parser.AnalyticFunction) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func checkArgsForListFunction(ctx context.Context, scope *ReferenceScope, expr parser.ListFunction) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func checkArgsForJsonAgg(expr parser.ListFunction) error { _ = "STUB: not implemented"; return nil }

func evalCaseExpr(ctx context.Context, scope *ReferenceScope, expr parser.CaseExpr) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalLogic(ctx context.Context, scope *ReferenceScope, expr parser.Logic) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalUnaryLogic(ctx context.Context, scope *ReferenceScope, expr parser.UnaryLogic) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalCursorStatus(expr parser.CursorStatus, scope *ReferenceScope) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalCursorAttribute(expr parser.CursorAttrebute, scope *ReferenceScope) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func evalPlaceholder(ctx context.Context, scope *ReferenceScope, expr parser.Placeholder) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

// EvalRowValue returns single or multiple fields, single record
func EvalRowValue(ctx context.Context, scope *ReferenceScope, expr parser.QueryExpression) (value.RowValue, error) {
	_ = "STUB: not implemented"
	return *new(value.RowValue), nil
}

// evalRowValueList returns multiple fields, multiple records
func evalRowValueList(ctx context.Context, scope *ReferenceScope, expr parser.QueryExpression) ([]value.RowValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
 * Returns single fields, multiple records
 */
func evalArray(ctx context.Context, scope *ReferenceScope, expr parser.QueryExpression) ([]value.RowValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func evalSubqueryForRowValue(ctx context.Context, scope *ReferenceScope, expr parser.Subquery) (value.RowValue, error) {
	_ = "STUB: not implemented"
	return *new(value.RowValue), nil
}

func evalJsonQueryForRowValue(ctx context.Context, scope *ReferenceScope, expr parser.JsonQuery) (value.RowValue, error) {
	_ = "STUB: not implemented"
	return *new(value.RowValue), nil
}

func evalValueList(ctx context.Context, scope *ReferenceScope, expr parser.ValueList) (value.RowValue, error) {
	_ = "STUB: not implemented"
	return *new(value.RowValue), nil
}

func evalSubqueryForRowValueList(ctx context.Context, scope *ReferenceScope, expr parser.Subquery) ([]value.RowValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func evalJsonQueryForRowValueList(ctx context.Context, scope *ReferenceScope, expr parser.JsonQuery) ([]value.RowValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func evalSubqueryForArray(ctx context.Context, scope *ReferenceScope, expr parser.Subquery) ([]value.RowValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func evalJsonQueryForArray(ctx context.Context, scope *ReferenceScope, expr parser.JsonQuery) ([]value.RowValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func evalJsonQueryParameters(ctx context.Context, scope *ReferenceScope, expr parser.JsonQuery) (value.Primary, value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), *new(value.Primary), nil
}

func EvaluateEmbeddedString(ctx context.Context, scope *ReferenceScope, embedded string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func writeEmbeddedExpression(ctx context.Context, scope *ReferenceScope, buf *bytes.Buffer, expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}
