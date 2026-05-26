package query

import (
	"errors"

	"github.com/mithrandie/csvq/lib/value"
)

var errIntegerDevidedByZero = errors.New("integer devided by zero")

func Calculate(p1 value.Primary, p2 value.Primary, operator int) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func calculateInteger(i1 int64, i2 int64, operator int) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func calculateFloat(f1 float64, f2 float64, operator int) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}
