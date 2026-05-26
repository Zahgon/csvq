package json

type PathLexer struct {
	PathScanner
	path  PathExpression
	token PathToken
	err   error
}

func (l *PathLexer) Lex(lval *jpSymType) int { _ = "STUB: not implemented"; return 0 }

func (l *PathLexer) Error(_ string) { _ = "STUB: not implemented"; return }

type PathToken struct {
	Token   int
	Literal string
	Column  int
}

type PathSyntaxError struct {
	Column  int
	Message string
}

func (e PathSyntaxError) Error() string { _ = "STUB: not implemented"; return "" }

func NewPathSyntaxError(message string, token PathToken) error {
	_ = "STUB: not implemented"
	return nil
}
