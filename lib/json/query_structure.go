package json

type QueryExpression interface{}

type Element struct {
	Label string
	Child QueryExpression
}

func (e Element) FieldLabel() string { _ = "STUB: not implemented"; return "" }

type ArrayItem struct {
	Index int
	Child QueryExpression
}

func (e ArrayItem) FieldLabel() string { _ = "STUB: not implemented"; return "" }

type RowValueExpr struct {
	Child QueryExpression
}

type TableExpr struct {
	Fields []FieldExpr
}

type FieldExpr struct {
	Element Element
	Alias   string
}

func (e FieldExpr) FieldLabel() string { _ = "STUB: not implemented"; return "" }

func EscapeIdentifier(s string) string { _ = "STUB: not implemented"; return "" }
