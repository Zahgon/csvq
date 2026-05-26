package query

import (
	"time"

	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/ternary"
)

func Is(p1 value.Primary, p2 value.Primary) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func Like(p1 value.Primary, p2 value.Primary) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func matchText(text []rune, pattern []rune) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func matchCondition(pattern []rune) (anyRunesMinLen int, anyRunesMaxLen int, searchWord []rune, restPattern []rune) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

func InRowValueList(rowValue value.RowValue, list []value.RowValue, matchType int, operator string, datetimeFormats []string, location *time.Location) (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}

// parser.ALL

// parser.ALL

func Any(rowValue value.RowValue, list []value.RowValue, operator string, datetimeFormats []string, location *time.Location) (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}

func All(rowValue value.RowValue, list []value.RowValue, operator string, datetimeFormats []string, location *time.Location) (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}
