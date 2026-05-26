package json

type QueryLexer struct {
	QueryScanner
	query QueryExpression
	token QueryToken
	err   error
}

func (l *QueryLexer) Lex(lval *jqSymType) int { _ = "STUB: not implemented"; return 0 }

func (l *QueryLexer) Error(e string) { _ = "STUB: not implemented"; return }

type QueryToken struct {
	Token   int
	Literal string
	Column  int
}

type QuerySyntaxError struct {
	Char    int
	Message string
}

func (e QuerySyntaxError) Error() string { _ = "STUB: not implemented"; return "" }

func NewQuerySyntaxError(message string, token QueryToken) error {
	_ = "STUB: not implemented"
	return nil
}
