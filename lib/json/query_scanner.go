package json

const AliasSpecifier = "AS"

type QueryScanner struct {
	src    []rune
	srcPos int
	offset int

	column int

	err error
}

func (s *QueryScanner) Init(src string) *QueryScanner { _ = "STUB: not implemented"; return nil }

func (s *QueryScanner) peek() rune { _ = "STUB: not implemented"; return 0 }

func (s *QueryScanner) next() rune { _ = "STUB: not implemented"; return 0 }

func (s *QueryScanner) runes() []rune { _ = "STUB: not implemented"; return nil }

func (s *QueryScanner) literal() string { _ = "STUB: not implemented"; return "" }

func (s *QueryScanner) trimQuotes() string { _ = "STUB: not implemented"; return "" }

func (s *QueryScanner) Scan() (QueryToken, error) {
	_ = "STUB: not implemented"
	return *new(QueryToken), nil
}

func (s *QueryScanner) skipSpaces() rune { _ = "STUB: not implemented"; return 0 }

func (s *QueryScanner) isDecimal(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *QueryScanner) isIdentRune(ch rune) bool { _ = "STUB: not implemented"; return false }

func (s *QueryScanner) scanIdentifier() { _ = "STUB: not implemented"; return }

func (s *QueryScanner) scanString(quote rune) { _ = "STUB: not implemented"; return }

func (s *QueryScanner) scanDecimal() { _ = "STUB: not implemented"; return }
