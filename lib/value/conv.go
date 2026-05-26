package value

import (
	"sync"
	"time"
)

var DatetimeFormats = NewDatetimeFormatMap()

type DatetimeFormatMap struct {
	m   *sync.Map
	mtx *sync.Mutex
}

func NewDatetimeFormatMap() DatetimeFormatMap {
	_ = "STUB: not implemented"
	return *new(DatetimeFormatMap)
}

func (dfmap DatetimeFormatMap) store(key string, value string) { _ = "STUB: not implemented"; return }

func (dfmap DatetimeFormatMap) load(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (dfmap DatetimeFormatMap) Get(s string) string { _ = "STUB: not implemented"; return "" }

func StrToTime(s string, formats []string, location *time.Location) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func ConvertDatetimeFormat(format string) string { _ = "STUB: not implemented"; return "" }

func Float64ToTime(f float64, location *time.Location) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func Int64ToStr(i int64) string { _ = "STUB: not implemented"; return "" }

func Float64ToStr(f float64, useScientificNotation bool) string {
	_ = "STUB: not implemented"
	return ""
}

func ToInteger(p Primary) Primary { _ = "STUB: not implemented"; return *new(Primary) }

func ToIntegerStrictly(p Primary) Primary { _ = "STUB: not implemented"; return *new(Primary) }

func ToFloat(p Primary) Primary { _ = "STUB: not implemented"; return *new(Primary) }

func ToDatetime(p Primary, formats []string, location *time.Location) Primary {
	_ = "STUB: not implemented"
	return *new(Primary)
}

func TimeFromUnixTime(sec int64, nano int64, location *time.Location) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func ToBoolean(p Primary) Primary { _ = "STUB: not implemented"; return *new(Primary) }

func ToString(p Primary) Primary { _ = "STUB: not implemented"; return *new(Primary) }
