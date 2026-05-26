package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

type View struct {
	Header    Header
	RecordSet RecordSet
	FileInfo  *FileInfo

	selectFields []int
	selectLabels []string
	isGrouped    bool

	comparisonKeysInEachRecord []string
	sortValuesInEachCell       [][]*SortValue
	sortValuesInEachRecord     []SortValues
	sortDirections             []int
	sortNullPositions          []int

	offset int
}

func NewView() *View { _ = "STUB: not implemented"; return nil }

func NewDualView() *View { _ = "STUB: not implemented"; return nil }

func NewViewFromGroupedRecord(ctx context.Context, flags *option.Flags, referenceRecord ReferenceRecord) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (view *View) IsUpdatable() bool { _ = "STUB: not implemented"; return false }

func (view *View) Where(ctx context.Context, scope *ReferenceScope, clause parser.WhereClause) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) filter(ctx context.Context, scope *ReferenceScope, condition parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) GroupBy(ctx context.Context, scope *ReferenceScope, clause parser.GroupByClause) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) group(ctx context.Context, scope *ReferenceScope, items []parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) groupAll(ctx context.Context, flags *option.Flags) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) Having(ctx context.Context, scope *ReferenceScope, clause parser.HavingClause) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) Select(ctx context.Context, scope *ReferenceScope, clause parser.SelectClause) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) GenerateComparisonKeys(ctx context.Context, flags *option.Flags) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) SelectAllColumns(ctx context.Context, scope *ReferenceScope) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) OrderBy(ctx context.Context, scope *ReferenceScope, clause parser.OrderByClause) error {
	_ = "STUB: not implemented"
	return nil
}

//parser.DESC

func (view *View) numberOfColumnsToBeAdded(exprs []parser.QueryExpression, funcs []parser.AnalyticFunction) int {
	_ = "STUB: not implemented"
	return 0
}

func (view *View) ExtendRecordCapacity(ctx context.Context, scope *ReferenceScope, exprs []parser.QueryExpression, funcs []parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) evalColumn(ctx context.Context, scope *ReferenceScope, obj parser.QueryExpression, alias string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) evalAnalyticFunction(ctx context.Context, scope *ReferenceScope, expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) Offset(ctx context.Context, scope *ReferenceScope, clause parser.OffsetClause) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) Limit(ctx context.Context, scope *ReferenceScope, clause parser.LimitClause) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) InsertValues(ctx context.Context, scope *ReferenceScope, fields []parser.QueryExpression, list []parser.QueryExpression) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) InsertFromQuery(ctx context.Context, scope *ReferenceScope, fields []parser.QueryExpression, query parser.SelectQuery) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) ReplaceValues(ctx context.Context, scope *ReferenceScope, fields []parser.QueryExpression, list []parser.QueryExpression, keys []parser.QueryExpression) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) ReplaceFromQuery(ctx context.Context, scope *ReferenceScope, fields []parser.QueryExpression, query parser.SelectQuery, keys []parser.QueryExpression) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) convertListToRecordValues(ctx context.Context, scope *ReferenceScope, fields []parser.QueryExpression, list []parser.QueryExpression) ([][]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (view *View) convertResultSetToRecordValues(ctx context.Context, scope *ReferenceScope, fields []parser.QueryExpression, query parser.SelectQuery) ([][]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (view *View) convertRecordValuesToRecordSet(ctx context.Context, fields []parser.QueryExpression, recordValues [][]value.Primary) (RecordSet, error) {
	_ = "STUB: not implemented"
	return *new(RecordSet), nil
}

func (view *View) insert(ctx context.Context, fields []parser.QueryExpression, recordValues [][]value.Primary) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) replace(ctx context.Context, flags *option.Flags, fields []parser.QueryExpression, recordValues [][]value.Primary, keys []parser.QueryExpression) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) Fix(ctx context.Context, flags *option.Flags) error {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) Union(ctx context.Context, flags *option.Flags, calcView *View, all bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) Except(ctx context.Context, flags *option.Flags, calcView *View, all bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) Intersect(ctx context.Context, flags *option.Flags, calcView *View, all bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (view *View) ListValuesForAggregateFunctions(ctx context.Context, scope *ReferenceScope, expr parser.QueryExpression, arg parser.QueryExpression, distinct bool) ([]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (view *View) RestoreHeaderReferences() error { _ = "STUB: not implemented"; return nil }

func (view *View) FieldIndex(fieldRef parser.QueryExpression) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) FieldIndices(fields []parser.QueryExpression) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (view *View) FieldViewName(fieldRef parser.QueryExpression) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (view *View) InternalRecordId(ref string, recordIndex int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (view *View) CreateRestorePoint() { _ = "STUB: not implemented"; return }

func (view *View) Restore() { _ = "STUB: not implemented"; return }

func (view *View) FieldLen() int { _ = "STUB: not implemented"; return 0 }

func (view *View) RecordLen() int { _ = "STUB: not implemented"; return 0 }

func (view *View) Len() int { _ = "STUB: not implemented"; return 0 }

func (view *View) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (view *View) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (view *View) Copy() *View { _ = "STUB: not implemented"; return nil }
