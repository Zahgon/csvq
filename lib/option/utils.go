package option

import (
	"github.com/mithrandie/go-text"
	txjson "github.com/mithrandie/go-text/json"
)

func EscapeString(s string) string { _ = "STUB: not implemented"; return "" }

func UnescapeString(s string, quote rune) string { _ = "STUB: not implemented"; return "" }

func EscapeIdentifier(s string) string { _ = "STUB: not implemented"; return "" }

func UnescapeIdentifier(s string, quote rune) string { _ = "STUB: not implemented"; return "" }

func QuoteString(s string) string { _ = "STUB: not implemented"; return "" }

func QuoteIdentifier(s string) string { _ = "STUB: not implemented"; return "" }

func VariableSymbol(s string) string { _ = "STUB: not implemented"; return "" }

func FlagSymbol(s string) string { _ = "STUB: not implemented"; return "" }

func EnvironmentVariableSymbol(s string) string { _ = "STUB: not implemented"; return "" }

func EnclosedEnvironmentVariableSymbol(s string) string { _ = "STUB: not implemented"; return "" }

func MustBeEnclosed(s string) bool { _ = "STUB: not implemented"; return false }

func RuntimeInformationSymbol(s string) string { _ = "STUB: not implemented"; return "" }

func FormatInt(i int, thousandsSeparator string) string { _ = "STUB: not implemented"; return "" }

func FormatNumber(f float64, precision int, decimalPoint string, thousandsSeparator string, decimalSeparator string) string {
	_ = "STUB: not implemented"
	return ""
}

func ParseEncoding(s string) (text.Encoding, error) {
	_ = "STUB: not implemented"
	return *new(text.Encoding), nil
}

func ParseLineBreak(s string) (text.LineBreak, error) {
	_ = "STUB: not implemented"
	return *new(text.LineBreak), nil
}

func ParseDelimiter(s string) (rune, error) { _ = "STUB: not implemented"; return 0, nil }

func ParseDelimiterPositions(s string) ([]int, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func ParseFormat(s string, et txjson.EscapeType) (Format, txjson.EscapeType, error) {
	_ = "STUB: not implemented"
	return *new(Format), *new(txjson.EscapeType), nil
}

func ParseJsonEscapeType(s string) (txjson.EscapeType, error) {
	_ = "STUB: not implemented"
	return *new(txjson.EscapeType), nil
}

func AppendStrIfNotExist(list []string, elem string) []string {
	_ = "STUB: not implemented"
	return nil
}

func TextWidth(s string, flags *Flags) int { _ = "STUB: not implemented"; return 0 }

func RuneWidth(r rune, flags *Flags) int { _ = "STUB: not implemented"; return 0 }

func TrimSpace(s string) string { _ = "STUB: not implemented"; return "" }
