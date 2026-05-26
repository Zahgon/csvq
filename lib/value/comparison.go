package value

import (
	"time"

	"github.com/mithrandie/ternary"
)

type ComparisonResult int

const (
	IsEqual ComparisonResult = iota
	IsBoolEqual
	IsNotEqual
	IsLess
	IsGreater
	IsIncommensurable
)

var comparisonResultLiterals = map[ComparisonResult]string{
	IsEqual:           "IsEqual",
	IsBoolEqual:       "IsBoolEqual",
	IsNotEqual:        "IsNotEqual",
	IsLess:            "IsLess",
	IsGreater:         "IsGreater",
	IsIncommensurable: "IsIncommensurable",
}

func (cr ComparisonResult) String() string { _ = "STUB: not implemented"; return "" }

func compareInteger(v1 int64, v2 int64) ComparisonResult {
	_ = "STUB: not implemented"
	return *new(ComparisonResult)
}

func compareFloat(v1 float64, v2 float64) ComparisonResult {
	_ = "STUB: not implemented"
	return *new(ComparisonResult)
}

func CompareCombinedly(p1 Primary, p2 Primary, datetimeFormats []string, location *time.Location) ComparisonResult {
	_ = "STUB: not implemented"
	return *new(ComparisonResult)
}

func Identical(p1 Primary, p2 Primary) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func Equal(p1 Primary, p2 Primary, datetimeFormats []string, location *time.Location) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func NotEqual(p1 Primary, p2 Primary, datetimeFormats []string, location *time.Location) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func Less(p1 Primary, p2 Primary, datetimeFormats []string, location *time.Location) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func Greater(p1 Primary, p2 Primary, datetimeFormats []string, location *time.Location) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func LessOrEqual(p1 Primary, p2 Primary, datetimeFormats []string, location *time.Location) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func GreaterOrEqual(p1 Primary, p2 Primary, datetimeFormats []string, location *time.Location) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

func Compare(p1 Primary, p2 Primary, operator string, datetimeFormats []string, location *time.Location) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

//case "<>", "!=":

func CompareRowValues(rowValue1 RowValue, rowValue2 RowValue, operator string, datetimeFormats []string, location *time.Location) (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}

func Equivalent(p1 Primary, p2 Primary, datetimeFormats []string, location *time.Location) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}
