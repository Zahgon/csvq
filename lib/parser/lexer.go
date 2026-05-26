package parser

type Lexer struct {
	Scanner
	program []Statement
	token   Token
	err     error
}

func (l *Lexer) Lex(lval *yySymType) int { _ = "STUB: not implemented"; return 0 }

func (l *Lexer) Error(e string) { _ = "STUB: not implemented"; return }

type Token struct {
	Token         int
	Literal       string
	Quoted        bool
	HolderOrdinal int
	Line          int
	Char          int
	SourceFile    string
}

func (t Token) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (t Token) String() string { _ = "STUB: not implemented"; return "" }

type SyntaxError struct {
	SourceFile string
	Line       int
	Char       int
	Message    string
}

func (e SyntaxError) Error() string { _ = "STUB: not implemented"; return "" }

func NewSyntaxError(message string, token Token) error { _ = "STUB: not implemented"; return nil }
