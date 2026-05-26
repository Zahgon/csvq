package syntax

type Store struct {
	Syntax Syntax
}

func NewStore() *Store { _ = "STUB: not implemented"; return nil }

func (s Store) Search(keys []string) []Expression { _ = "STUB: not implemented"; return nil }

func (s Store) searchSyntax(keys []string, syntax Syntax) []Expression {
	_ = "STUB: not implemented"
	return nil
}

func (s Store) search(keys []string, exp Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

func (s Store) contains(keys []string, content string, name string) bool {
	_ = "STUB: not implemented"
	return false
}
