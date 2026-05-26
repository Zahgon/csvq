package json

const (
	PathSeparator = '.'
	PathEscape    = '\\'
)

const EOF = -1

type PathScanner struct {
	src    []rune
	srcPos int
	offset int

	column int
}

func (s *PathScanner) Init(src string) *PathScanner { _ = "STUB: not implemented"; return nil }

func (s *PathScanner) peek() rune { _ = "STUB: not implemented"; return 0 }

func (s *PathScanner) next() rune { _ = "STUB: not implemented"; return 0 }

func (s *PathScanner) runes() []rune { _ = "STUB: not implemented"; return nil }

func (s *PathScanner) literal() string { _ = "STUB: not implemented"; return "" }

func (s *PathScanner) Scan() PathToken { _ = "STUB: not implemented"; return *new(PathToken) }

func (s *PathScanner) scanObjectMember() { _ = "STUB: not implemented"; return }

func (s *PathScanner) unescapeObjectMember(src string) string { _ = "STUB: not implemented"; return "" }
