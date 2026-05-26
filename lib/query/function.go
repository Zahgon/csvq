package query

import (
	"context"
	"errors"
	"hash"
	"regexp"
	"time"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

type BuiltInFunction func(parser.Function, []value.Primary, *option.Flags) (value.Primary, error)

var Functions = map[string]BuiltInFunction{
	"COALESCE":               Coalesce,
	"IF":                     If,
	"IFNULL":                 Ifnull,
	"NULLIF":                 Nullif,
	"CEIL":                   Ceil,
	"FLOOR":                  Floor,
	"ROUND":                  Round,
	"ABS":                    Abs,
	"ACOS":                   Acos,
	"ACOSH":                  Acosh,
	"ASIN":                   Asin,
	"ASINH":                  Asinh,
	"ATAN":                   Atan,
	"ATAN2":                  Atan2,
	"ATANH":                  Atanh,
	"CBRT":                   Cbrt,
	"COS":                    Cos,
	"COSH":                   Cosh,
	"EXP":                    Exp,
	"EXP2":                   Exp2,
	"EXPM1":                  Expm1,
	"IS_INF":                 IsInf,
	"IS_NAN":                 IsNaN,
	"LOG":                    MathLog,
	"LOG10":                  Log10,
	"LOG1P":                  Log1p,
	"LOG2":                   Log2,
	"LOGB":                   Logb,
	"POW":                    Pow,
	"SIN":                    Sin,
	"SINH":                   Sinh,
	"SQRT":                   Sqrt,
	"TAN":                    Tan,
	"TANH":                   Tanh,
	"BIN_TO_DEC":             BinToDec,
	"OCT_TO_DEC":             OctToDec,
	"HEX_TO_DEC":             HexToDec,
	"ENOTATION_TO_DEC":       EnotationToDec,
	"BIN":                    Bin,
	"OCT":                    Oct,
	"HEX":                    Hex,
	"ENOTATION":              Enotation,
	"NUMBER_FORMAT":          NumberFormat,
	"RAND":                   Rand,
	"TRIM":                   Trim,
	"LTRIM":                  Ltrim,
	"RTRIM":                  Rtrim,
	"UPPER":                  Upper,
	"LOWER":                  Lower,
	"BASE64_ENCODE":          Base64Encode,
	"BASE64_DECODE":          Base64Decode,
	"HEX_ENCODE":             HexEncode,
	"HEX_DECODE":             HexDecode,
	"LEN":                    Len,
	"BYTE_LEN":               ByteLen,
	"WIDTH":                  Width,
	"LPAD":                   Lpad,
	"RPAD":                   Rpad,
	"SUBSTRING":              Substring,
	"SUBSTR":                 Substr,
	"INSTR":                  Instr,
	"LIST_ELEM":              ListElem,
	"REPLACE":                ReplaceFn,
	"REGEXP_MATCH":           RegExpMatch,
	"REGEXP_FIND":            RegExpFind,
	"REGEXP_FIND_SUBMATCHES": RegExpFindSubMatches,
	"REGEXP_FIND_ALL":        RegExpFindAll,
	"REGEXP_REPLACE":         RegExpReplace,
	"TITLE_CASE":             TitleCase,
	"FORMAT":                 Format,
	"JSON_VALUE":             JsonValue,
	"MD5":                    Md5,
	"SHA1":                   Sha1,
	"SHA256":                 Sha256,
	"SHA512":                 Sha512,
	"MD5_HMAC":               Md5Hmac,
	"SHA1_HMAC":              Sha1Hmac,
	"SHA256_HMAC":            Sha256Hmac,
	"SHA512_HMAC":            Sha512Hmac,
	"DATETIME_FORMAT":        DatetimeFormat,
	"YEAR":                   Year,
	"MONTH":                  Month,
	"DAY":                    Day,
	"HOUR":                   Hour,
	"MINUTE":                 Minute,
	"SECOND":                 Second,
	"MILLISECOND":            Millisecond,
	"MICROSECOND":            Microsecond,
	"NANOSECOND":             Nanosecond,
	"WEEKDAY":                Weekday,
	"UNIX_TIME":              UnixTime,
	"UNIX_NANO_TIME":         UnixNanoTime,
	"DAY_OF_YEAR":            DayOfYear,
	"WEEK_OF_YEAR":           WeekOfYear,
	"ADD_YEAR":               AddYear,
	"ADD_MONTH":              AddMonth,
	"ADD_DAY":                AddDay,
	"ADD_HOUR":               AddHour,
	"ADD_MINUTE":             AddMinute,
	"ADD_SECOND":             AddSecond,
	"ADD_MILLI":              AddMilli,
	"ADD_MICRO":              AddMicro,
	"ADD_NANO":               AddNano,
	"TRUNC_MONTH":            TruncMonth,
	"TRUNC_DAY":              TruncDay,
	"TRUNC_TIME":             TruncTime,
	"TRUNC_HOUR":             TruncTime,
	"TRUNC_MINUTE":           TruncMinute,
	"TRUNC_SECOND":           TruncSecond,
	"TRUNC_MILLI":            TruncMilli,
	"TRUNC_MICRO":            TruncMicro,
	"TRUNC_NANO":             TruncNano,
	"DATE_DIFF":              DateDiff,
	"TIME_DIFF":              TimeDiff,
	"TIME_NANO_DIFF":         TimeNanoDiff,
	"UTC":                    UTC,
	"MILLI_TO_DATETIME":      MilliToDatetime,
	"NANO_TO_DATETIME":       NanoToDatetime,
	"STRING":                 String,
	"INTEGER":                Integer,
	"FLOAT":                  Float,
	"BOOLEAN":                Boolean,
	"TERNARY":                Ternary,
	"DATETIME":               Datetime,
}

type Direction string

const (
	RightDirection Direction = "R"
	LeftDirection            = "L"
)

type PaddingType string

const (
	PaddingRuneCount PaddingType = "LEN"
	PaddingByteCount PaddingType = "BYTE"
	PaddingWidth     PaddingType = "WIDTH"
)

func Coalesce(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func If(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Ifnull(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Nullif(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func roundParams(args []value.Primary) (number float64, place float64, isnull bool, argsErr bool) {
	_ = "STUB: not implemented"
	return 0, 0, false, false
}

func Ceil(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Floor(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func round(f float64, place float64) float64 { _ = "STUB: not implemented"; return 0 }

func Round(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execMath1Arg(fn parser.Function, args []value.Primary, mathf func(float64) float64) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execMath2Args(fn parser.Function, args []value.Primary, mathf func(float64, float64) float64) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Abs(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Acos(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Acosh(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Asin(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Asinh(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Atan(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Atan2(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Atanh(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Cbrt(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Cos(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Cosh(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Exp(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Exp2(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Expm1(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func IsInf(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func IsNaN(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func MathLog(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Log10(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Log1p(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Log2(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Logb(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Pow(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sin(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sinh(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sqrt(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Tan(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Tanh(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execParseInt(fn parser.Function, args []value.Primary, base int) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execFormatInt(fn parser.Function, args []value.Primary, base int) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func BinToDec(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func OctToDec(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func HexToDec(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func EnotationToDec(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Bin(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Oct(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Hex(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Enotation(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func NumberFormat(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Rand(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execStrings1Arg(fn parser.Function, args []value.Primary, stringsf func(string) string) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execStringsTrim(fn parser.Function, args []value.Primary, stringsf func(string, string) string) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func base64Encode(s string) string { _ = "STUB: not implemented"; return "" }

func base64Decode(s string) string { _ = "STUB: not implemented"; return "" }

func hexEncode(s string) string { _ = "STUB: not implemented"; return "" }

func hexDecode(s string) string { _ = "STUB: not implemented"; return "" }

func trim(s string, cutset string) string { _ = "STUB: not implemented"; return "" }

func ltrim(s string, cutset string) string { _ = "STUB: not implemented"; return "" }

func rtrim(s string, cutset string) string { _ = "STUB: not implemented"; return "" }

func execStringsLen(fn parser.Function, args []value.Primary, stringsf func(string) int) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execStringsPadding(fn parser.Function, args []value.Primary, direction Direction, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

// Do Nothing

func execCrypto(fn parser.Function, args []value.Primary, cryptof func() hash.Hash) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execCryptoHMAC(fn parser.Function, args []value.Primary, cryptof func() hash.Hash) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Trim(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Ltrim(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Rtrim(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Upper(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Lower(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Base64Encode(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Base64Decode(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func HexEncode(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func HexDecode(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Len(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func ByteLen(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Width(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Lpad(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Rpad(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func substr(fn parser.Function, args []value.Primary, zeroBasedIndex bool) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Substring(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Substr(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Instr(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func ListElem(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func ReplaceFn(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

var regExpStrIsNull = errors.New("cannot convert the value to string")

func prepareRegExp(fn parser.Function, str value.Primary, expr value.Primary, flags value.Primary) (string, *regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func prepareRegExpMatch(fn parser.Function, args []value.Primary) (string, *regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func prepareRegExpReplace(fn parser.Function, args []value.Primary) (string, *regexp.Regexp, string, error) {
	_ = "STUB: not implemented"
	return "", nil, "", nil
}

func RegExpMatch(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func RegExpFind(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func RegExpFindSubMatches(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func RegExpFindAll(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func RegExpReplace(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TitleCase(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Format(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func JsonValue(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Md5(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sha1(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sha256(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sha512(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Md5Hmac(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sha1Hmac(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sha256Hmac(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Sha512Hmac(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func DatetimeFormat(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execDatetimeToInt(fn parser.Function, args []value.Primary, timef func(time.Time) int64, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func execDatetimeAdd(fn parser.Function, args []value.Primary, timef func(time.Time, int) time.Time, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func year(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func month(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func day(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func hour(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func minute(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func second(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func millisecond(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func microsecond(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func nanosecond(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func weekday(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func unixTime(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func unixNanoTime(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func dayOfYear(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func weekOfYear(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func addYear(t time.Time, duration int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func addMonth(t time.Time, duration int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func addDay(t time.Time, duration int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func addHour(t time.Time, duration int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func addMinute(t time.Time, duration int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func addSecond(t time.Time, duration int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func addMilli(t time.Time, duration int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func addMicro(t time.Time, duration int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func addNano(t time.Time, duration int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func Year(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Month(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Day(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Hour(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Minute(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Second(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Millisecond(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Microsecond(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Nanosecond(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Weekday(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func UnixTime(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func UnixNanoTime(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func DayOfYear(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func WeekOfYear(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddYear(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddMonth(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddDay(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddHour(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddMinute(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddSecond(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddMilli(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddMicro(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func AddNano(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func truncateDate(fn parser.Function, args []value.Primary, place int8, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TruncMonth(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TruncDay(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TruncTime(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func truncateDuration(fn parser.Function, args []value.Primary, dur time.Duration, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TruncMinute(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TruncSecond(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TruncMilli(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TruncMicro(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TruncNano(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func DateDiff(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func timeDiff(fn parser.Function, args []value.Primary, durf func(time.Duration) value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func durationSeconds(dur time.Duration) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func durationNanoseconds(dur time.Duration) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func TimeDiff(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func TimeNanoDiff(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func UTC(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func MilliToDatetime(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func NanoToDatetime(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func String(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Integer(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Float(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Boolean(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Ternary(fn parser.Function, args []value.Primary, _ *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Datetime(fn parser.Function, args []value.Primary, flags *option.Flags) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Call(ctx context.Context, fn parser.Function, args []value.Primary) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func Now(scope *ReferenceScope, fn parser.Function, args []value.Primary) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func JsonObject(ctx context.Context, scope *ReferenceScope, fn parser.Function) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}
