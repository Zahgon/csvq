package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

var AnalyticFunctions = map[string]AnalyticFunction{
	"ROW_NUMBER":   RowNumber{},
	"RANK":         Rank{},
	"DENSE_RANK":   DenseRank{},
	"CUME_DIST":    CumeDist{},
	"PERCENT_RANK": PercentRank{},
	"NTILE":        NTile{},
	"FIRST_VALUE":  FirstValue{},
	"LAST_VALUE":   LastValue{},
	"NTH_VALUE":    NthValue{},
	"LAG":          Lag{},
	"LEAD":         Lead{},
	"LISTAGG":      AnalyticListAgg{},
	"JSON_AGG":     AnalyticJsonAgg{},
}

type AnalyticFunction interface {
	CheckArgsLen(expr parser.AnalyticFunction) error
	Execute(context.Context, *ReferenceScope, Partition, parser.AnalyticFunction) (map[int]value.Primary, error)
}

type Partition []int

func (p Partition) Reverse() { _ = "STUB: not implemented"; return }

type Partitions map[string]Partition

func Analyze(ctx context.Context, scope *ReferenceScope, view *View, fn parser.AnalyticFunction, partitionIndices []int) error {
	_ = "STUB: not implemented"
	return nil
}

//User Defined Function

type WindowFrame struct {
	Low     int
	High    int
	Records []int
}

func WindowFrameSet(partition Partition, expr parser.AnalyticClause) []WindowFrame {
	_ = "STUB: not implemented"
	return nil
}

func windowValues(ctx context.Context, scope *ReferenceScope, frame WindowFrame, partition Partition, expr parser.AnalyticFunction, valueCache map[int]value.Primary) ([]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CheckArgsLen(expr parser.AnalyticFunction, length []int) error {
	_ = "STUB: not implemented"
	return nil
}

type RowNumber struct{}

func (fn RowNumber) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn RowNumber) Execute(_ context.Context, _ *ReferenceScope, partition Partition, _ parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Rank struct{}

func (fn Rank) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn Rank) Execute(_ context.Context, scope *ReferenceScope, partition Partition, _ parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DenseRank struct{}

func (fn DenseRank) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn DenseRank) Execute(_ context.Context, scope *ReferenceScope, partition Partition, _ parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CumeDist struct{}

func (fn CumeDist) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn CumeDist) Execute(_ context.Context, scope *ReferenceScope, partition Partition, _ parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PercentRank struct{}

func (fn PercentRank) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn PercentRank) Execute(_ context.Context, scope *ReferenceScope, partition Partition, _ parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func perseCumulativeGroups(partition Partition, view *View) [][]int {
	_ = "STUB: not implemented"
	return nil
}

type NTile struct{}

func (fn NTile) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn NTile) Execute(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FirstValue struct{}

func (fn FirstValue) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn FirstValue) Execute(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LastValue struct{}

func (fn LastValue) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn LastValue) Execute(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type NthValue struct{}

func (fn NthValue) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn NthValue) Execute(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setNthValue(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction, n int) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Lag struct{}

func (fn Lag) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn Lag) Execute(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Lead struct{}

func (fn Lead) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn Lead) Execute(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setLag(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AnalyticListAgg struct{}

func (fn AnalyticListAgg) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn AnalyticListAgg) Execute(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AnalyticJsonAgg struct{}

func (fn AnalyticJsonAgg) CheckArgsLen(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (fn AnalyticJsonAgg) Execute(ctx context.Context, scope *ReferenceScope, partition Partition, expr parser.AnalyticFunction) (map[int]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
