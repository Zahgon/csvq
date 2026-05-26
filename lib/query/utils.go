package query

import (
	"bytes"
	"sync"
	"time"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/ternary"
)

const LimitToUseUintSlicePool = 20

type UintPool struct {
	limitToUseSlice int
	m               map[uint]bool
	values          []uint
}

func NewUintPool(initCap int, limitToUseSlice int) *UintPool { _ = "STUB: not implemented"; return nil }

func (c *UintPool) Exists(val uint) bool { _ = "STUB: not implemented"; return false }

func (c *UintPool) Add(val uint) { _ = "STUB: not implemented"; return }

func (c *UintPool) Range(fn func(idx int, value uint) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *UintPool) Len() int { _ = "STUB: not implemented"; return 0 }

func InStrSliceWithCaseInsensitive(s string, list []string) bool {
	_ = "STUB: not implemented"
	return false
}

func Distinguish(list []value.Primary, flags *option.Flags) []value.Primary {
	_ = "STUB: not implemented"
	return nil
}

func FormatCount(i int, obj string) string { _ = "STUB: not implemented"; return "" }

var comparisonKeysBufPool = &sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func GetComparisonKeysBuf() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func PutComparisonkeysBuf(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func SerializeComparisonKeys(buf *bytes.Buffer, values []value.Primary, flags *option.Flags) {
	_ = "STUB: not implemented"
	return
}

func SerializeKey(buf *bytes.Buffer, val value.Primary, flags *option.Flags) {
	_ = "STUB: not implemented"
	return
}

func SerializeIdenticalKey(buf *bytes.Buffer, val value.Primary) { _ = "STUB: not implemented"; return }

func serializeNull(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func serializeInteger(buf *bytes.Buffer, s string) { _ = "STUB: not implemented"; return }

func serializeFloat(buf *bytes.Buffer, s string) { _ = "STUB: not implemented"; return }

func serializeDatetime(buf *bytes.Buffer, t time.Time) { _ = "STUB: not implemented"; return }

func serializeDatetimeFromUnixNano(buf *bytes.Buffer, t int64) { _ = "STUB: not implemented"; return }

func serializeString(buf *bytes.Buffer, s string) { _ = "STUB: not implemented"; return }

func serializeCaseSensitiveString(buf *bytes.Buffer, s string) { _ = "STUB: not implemented"; return }

func serializeBoolean(buf *bytes.Buffer, b bool) { _ = "STUB: not implemented"; return }

func serializeTernary(buf *bytes.Buffer, t ternary.Value) { _ = "STUB: not implemented"; return }
