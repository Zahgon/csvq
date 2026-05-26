package query

import (
	"github.com/mithrandie/csvq/lib/parser"
)

func HasAggregateFunction(expr parser.QueryExpression, scope *ReferenceScope) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func HasAggregateFunctionInList(list []parser.QueryExpression, scope *ReferenceScope) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetValuesInOrderByClause(e parser.OrderByClause) []parser.QueryExpression {
	_ = "STUB: not implemented"
	return nil
}

func hasAggFuncInRowValueComparison(lhs parser.QueryExpression, values parser.QueryExpression, scope *ReferenceScope) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func hasAggFuncInRowValue(expr parser.QueryExpression, scope *ReferenceScope) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func SearchAnalyticFunctions(expr parser.QueryExpression) ([]parser.AnalyticFunction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SearchAnalyticFunctionsInList(list []parser.QueryExpression) ([]parser.AnalyticFunction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendAnalyticFunctionToListIfNotExist(list1 []parser.AnalyticFunction, list2 []parser.AnalyticFunction) []parser.AnalyticFunction {
	_ = "STUB: not implemented"
	return nil
}

func searchAnalyticFunctionsInRowValueComparison(lhs parser.QueryExpression, values parser.QueryExpression) ([]parser.AnalyticFunction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func searchAnalyticFunctionsInRowValue(expr parser.QueryExpression) ([]parser.AnalyticFunction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
