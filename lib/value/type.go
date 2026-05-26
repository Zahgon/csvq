package value

import (
	"time"

	"github.com/mithrandie/ternary"
)

var ternaryTrue = &Ternary{value: ternary.TRUE}
var ternaryFalse = &Ternary{value: ternary.FALSE}
var ternaryUnknown = &Ternary{value: ternary.UNKNOWN}
var booleanTrue = &Boolean{value: true}
var booleanFalse = &Boolean{value: false}
var null = &Null{}

func IsTrue(v Primary) bool { _ = "STUB: not implemented"; return false }

func IsFalse(v Primary) bool { _ = "STUB: not implemented"; return false }

func IsUnknown(v Primary) bool { _ = "STUB: not implemented"; return false }

func IsNull(v Primary) bool { _ = "STUB: not implemented"; return false }

type RowValue []Primary

type Primary interface {
	String() string
	Ternary() ternary.Value
}

type String struct {
	literal string
}

func (s String) String() string { _ = "STUB: not implemented"; return "" }

func NewString(s string) *String { _ = "STUB: not implemented"; return nil }

func (s String) Raw() string { _ = "STUB: not implemented"; return "" }

func (s String) Ternary() ternary.Value { _ = "STUB: not implemented"; return *new(ternary.Value) }

type Integer struct {
	value int64
}

func NewIntegerFromString(s string) *Integer { _ = "STUB: not implemented"; return nil }

func NewInteger(i int64) *Integer { _ = "STUB: not implemented"; return nil }

func (i Integer) String() string { _ = "STUB: not implemented"; return "" }

func (i Integer) Raw() int64 { _ = "STUB: not implemented"; return 0 }

func (i Integer) Ternary() ternary.Value { _ = "STUB: not implemented"; return *new(ternary.Value) }

type Float struct {
	value float64
}

func NewFloatFromString(s string) *Float { _ = "STUB: not implemented"; return nil }

func NewFloat(f float64) *Float { _ = "STUB: not implemented"; return nil }

func (f Float) String() string { _ = "STUB: not implemented"; return "" }

func (f Float) Raw() float64 { _ = "STUB: not implemented"; return 0 }

func (f Float) Ternary() ternary.Value { _ = "STUB: not implemented"; return *new(ternary.Value) }

type Boolean struct {
	value bool
}

func NewBoolean(b bool) *Boolean { _ = "STUB: not implemented"; return nil }

func (b Boolean) String() string { _ = "STUB: not implemented"; return "" }

func (b Boolean) Raw() bool { _ = "STUB: not implemented"; return false }

func (b Boolean) Ternary() ternary.Value { _ = "STUB: not implemented"; return *new(ternary.Value) }

type Ternary struct {
	value ternary.Value
}

func NewTernaryFromString(s string) *Ternary { _ = "STUB: not implemented"; return nil }

func NewTernary(t ternary.Value) *Ternary { _ = "STUB: not implemented"; return nil }

func (t Ternary) String() string { _ = "STUB: not implemented"; return "" }

func (t Ternary) Ternary() ternary.Value { _ = "STUB: not implemented"; return *new(ternary.Value) }

type Datetime struct {
	value time.Time
}

func NewDatetimeFromString(s string, formats []string, location *time.Location) *Datetime {
	_ = "STUB: not implemented"
	return nil
}

func NewDatetime(t time.Time) *Datetime { _ = "STUB: not implemented"; return nil }

func (dt Datetime) String() string { _ = "STUB: not implemented"; return "" }

func (dt Datetime) Raw() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (dt Datetime) Ternary() ternary.Value { _ = "STUB: not implemented"; return *new(ternary.Value) }

func (dt Datetime) Format(s string) string { _ = "STUB: not implemented"; return "" }

type Null struct{}

func NewNull() *Null { _ = "STUB: not implemented"; return nil }

func (n Null) String() string { _ = "STUB: not implemented"; return "" }

func (n Null) Ternary() ternary.Value { _ = "STUB: not implemented"; return *new(ternary.Value) }
