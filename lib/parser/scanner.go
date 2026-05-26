package parser

import (
	"bytes"
	"errors"
)

const (
	EOF = -(iota + 1)
	Uncategorized
)

const (
	TokenFrom   = IDENTIFIER
	TokenTo     = SUBSTITUTION_OP
	KeywordFrom = SELECT
	KeywordTo   = JSON_OBJECT
)

const (
	VariableSign            = '@'
	EnvironmentVariableSign = '%'
	ExternalCommandSign     = '$'
	RuntimeInformationSign  = '#'

	SubstitutionOperator = ":="

	BeginExpression = '{'
	EndExpression   = '}'

	IdentifierDelimiter = ':'
)

var errTokenIsNotKeyword = errors.New("token is not keyword")
var errInvalidConstantSyntax = errors.New("invalid constant syntax")

var comparisonOperators = []string{
	">",
	"<",
	">=",
	"<=",
	"<>",
	"!=",
	"==",
}

var stringOperators = []string{
	"||",
}

var runesNotIncludedInUrl = []rune{
	'{',
	'}',
	'|',
	'\\',
	'^',
	'[',
	']',
	'`',
}

var aggregateFunctions = []string{
	"MIN",
	"MAX",
	"SUM",
	"AVG",
	"STDEV",
	"STDEVP",
	"VARP",
	"MEDIAN",
}

var listFunctions = []string{
	"LISTAGG",
	"JSON_AGG",
}

var analyticFunctions = []string{
	"ROW_NUMBER",
	"RANK",
	"DENSE_RANK",
	"CUME_DIST",
	"PERCENT_RANK",
	"NTILE",
}

var functionsNth = []string{
	"FIRST_VALUE",
	"LAST_VALUE",
	"NTH_VALUE",
}

var functionsWithIgnoreNulls = []string{
	"LAG",
	"LEAD",
}

var ConstantDelimiter = string(IdentifierDelimiter) + string(IdentifierDelimiter)

func TokenLiteral(token int) string { _ = "STUB: not implemented"; return "" }

func KeywordLiteral(token int) (string, error) { _ = "STUB: not implemented"; return "", nil }

type Scanner struct {
	src     []rune
	srcPos  int
	literal bytes.Buffer

	line       int
	char       int
	sourceFile string

	forPrepared bool
	ansiQuotes  bool

	holderOrdinal int
	holderNames   []string
	holderNumber  int
}

func (s *Scanner) Init(src string, sourceFile string, forPrepared bool, ansiQuotes bool) *Scanner {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) HolderNumber() int { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) holderNameExists(name string) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) peek() rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) peekFurtherAhead(n int) rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) peekNextLetter(n int) rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) next() rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) checkNewLine(ch rune) rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) Scan() (Token, error) { _ = "STUB: not implemented"; return *new(Token), nil }

func (s *Scanner) scanString(quote rune) error { _ = "STUB: not implemented"; return nil }

func (s *Scanner) scanIdentifier(head rune) { _ = "STUB: not implemented"; return }

func (s *Scanner) scanConstant() error { _ = "STUB: not implemented"; return nil }

func (s *Scanner) scanUrl() int { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) isRuneNotIncludedInUrl(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isIdentRune(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isDecimal(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanNumber(head rune) (rune, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Scanner) scanOperator(head rune) { _ = "STUB: not implemented"; return }

func (s *Scanner) isOperatorRune(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) searchKeyword(str string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Scanner) isAggregateFunctions(str string) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isListaggFunctions(str string) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isAnalyticFunctions(str string) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isFunctionsNth(str string) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isFunctionsWithIgnoreNulls(str string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Scanner) isComparisonOperators(str string) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isStringOperators(str string) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) isCommentRune(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanComment() { _ = "STUB: not implemented"; return }

func (s *Scanner) isLineCommentRune(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *Scanner) scanLineComment() { _ = "STUB: not implemented"; return }

func (s *Scanner) scanExternalCommand() { _ = "STUB: not implemented"; return }

func (s *Scanner) scanExternalCommandQuotedString(quote rune) { _ = "STUB: not implemented"; return }

func (s *Scanner) scanExternalCommandCSVQExpression() { _ = "STUB: not implemented"; return }
