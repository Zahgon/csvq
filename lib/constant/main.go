package constant

import (
	"errors"

	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

var ErrInvalidType = errors.New("invalid constant type")
var ErrUndefined = errors.New("constant is not defined")

func Get(expr parser.Constant) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func ConvertConstantToPrivamryValue(c interface{}) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Count() int { _ = "STUB: not implemented"; return 0 }
