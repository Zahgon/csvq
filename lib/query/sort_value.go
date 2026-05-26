package query

import (
	"bytes"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/ternary"
)

type SortValueType int

const (
	NullType SortValueType = iota
	IntegerType
	FloatType
	DatetimeType
	BooleanType
	StringType
)

type SortValues []*SortValue

func (values SortValues) Less(compareValues SortValues, directions []int, nullPositions []int) bool {
	_ = "STUB: not implemented"
	return false
}

func (values SortValues) EquivalentTo(compareValues SortValues) bool {
	_ = "STUB: not implemented"
	return false
}

func (values SortValues) Serialize(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

type SortValue struct {
	Type SortValueType

	SerializedKey *bytes.Buffer

	Integer  int64
	Float    float64
	Datetime int64
	String   string
}

func NewSortValue(val value.Primary, flags *option.Flags) *SortValue {
	_ = "STUB: not implemented"
	return nil
}

func (v *SortValue) Less(compareValue *SortValue) ternary.Value {
	_ = "STUB: not implemented"
	return *new(ternary.Value)
}

// math.IsNaN(compareValue.Float)

func (v *SortValue) EquivalentTo(compareValue *SortValue) bool {
	_ = "STUB: not implemented"
	return false
}
