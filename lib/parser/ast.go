package parser

import (
	"time"

	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/ternary"
)

const TokenUndefined = 0

type Statement interface{}

type Expression interface {
	GetBaseExpr() *BaseExpr
	ClearBaseExpr()
	HasParseInfo() bool
	Line() int
	Char() int
	SourceFile() string
}

type QueryExpression interface {
	String() string

	GetBaseExpr() *BaseExpr
	ClearBaseExpr()
	HasParseInfo() bool
	Line() int
	Char() int
	SourceFile() string
}

type BaseExpr struct {
	line       int
	char       int
	sourceFile string
}

func (e *BaseExpr) Line() int { _ = "STUB: not implemented"; return 0 }

func (e *BaseExpr) Char() int { _ = "STUB: not implemented"; return 0 }

func (e *BaseExpr) SourceFile() string { _ = "STUB: not implemented"; return "" }

func (e *BaseExpr) HasParseInfo() bool { _ = "STUB: not implemented"; return false }

func (e *BaseExpr) GetBaseExpr() *BaseExpr { _ = "STUB: not implemented"; return nil }

func (e *BaseExpr) ClearBaseExpr() { _ = "STUB: not implemented"; return }

func NewBaseExpr(token Token) *BaseExpr { _ = "STUB: not implemented"; return nil }

type PrimitiveType struct {
	*BaseExpr
	Literal string
	Value   value.Primary
}

func NewStringValue(s string) PrimitiveType { _ = "STUB: not implemented"; return *new(PrimitiveType) }

func NewIntegerValueFromString(s string) PrimitiveType {
	_ = "STUB: not implemented"
	return *new(PrimitiveType)
}

func NewIntegerValue(i int64) PrimitiveType { _ = "STUB: not implemented"; return *new(PrimitiveType) }

func NewFloatValueFromString(s string) PrimitiveType {
	_ = "STUB: not implemented"
	return *new(PrimitiveType)
}

func NewFloatValue(f float64) PrimitiveType { _ = "STUB: not implemented"; return *new(PrimitiveType) }

func NewTernaryValueFromString(s string) PrimitiveType {
	_ = "STUB: not implemented"
	return *new(PrimitiveType)
}

func NewTernaryValue(t ternary.Value) PrimitiveType {
	_ = "STUB: not implemented"
	return *new(PrimitiveType)
}

func NewDatetimeValueFromString(s string, formats []string, location *time.Location) PrimitiveType {
	_ = "STUB: not implemented"
	return *new(PrimitiveType)
}

func NewDatetimeValue(t time.Time) PrimitiveType {
	_ = "STUB: not implemented"
	return *new(PrimitiveType)
}

func NewNullValue() PrimitiveType { _ = "STUB: not implemented"; return *new(PrimitiveType) }

func (e PrimitiveType) String() string { _ = "STUB: not implemented"; return "" }

func (e PrimitiveType) IsInteger() bool { _ = "STUB: not implemented"; return false }

type Placeholder struct {
	*BaseExpr
	Literal string
	Ordinal int
	Name    string
}

func (e Placeholder) String() string { _ = "STUB: not implemented"; return "" }

type Identifier struct {
	*BaseExpr
	Literal string
	Quoted  bool
}

func (i Identifier) String() string { _ = "STUB: not implemented"; return "" }

type Constant struct {
	*BaseExpr
	Space string
	Name  string
}

func (e Constant) String() string { _ = "STUB: not implemented"; return "" }

type FieldReference struct {
	*BaseExpr
	View   Identifier
	Column QueryExpression
}

func (e FieldReference) String() string { _ = "STUB: not implemented"; return "" }

type ColumnNumber struct {
	*BaseExpr
	View   Identifier
	Number *value.Integer
}

func (e ColumnNumber) String() string { _ = "STUB: not implemented"; return "" }

type Parentheses struct {
	*BaseExpr
	Expr QueryExpression
}

func (p Parentheses) String() string { _ = "STUB: not implemented"; return "" }

type RowValue struct {
	*BaseExpr
	Value QueryExpression
}

func (e RowValue) String() string { _ = "STUB: not implemented"; return "" }

type ValueList struct {
	*BaseExpr
	Values []QueryExpression
}

func (e ValueList) String() string { _ = "STUB: not implemented"; return "" }

type RowValueList struct {
	*BaseExpr
	RowValues []QueryExpression
}

func (e RowValueList) String() string { _ = "STUB: not implemented"; return "" }

type SelectQuery struct {
	*BaseExpr
	WithClause    QueryExpression
	SelectEntity  QueryExpression
	OrderByClause QueryExpression
	LimitClause   QueryExpression
	Context       Token
}

func (e SelectQuery) IsForUpdate() bool { _ = "STUB: not implemented"; return false }

func (e SelectQuery) String() string { _ = "STUB: not implemented"; return "" }

type SelectSet struct {
	*BaseExpr
	LHS      QueryExpression
	Operator Token
	All      Token
	RHS      QueryExpression
}

func (e SelectSet) String() string { _ = "STUB: not implemented"; return "" }

type SelectEntity struct {
	*BaseExpr
	SelectClause  QueryExpression
	IntoClause    QueryExpression
	FromClause    QueryExpression
	WhereClause   QueryExpression
	GroupByClause QueryExpression
	HavingClause  QueryExpression
}

func (e SelectEntity) String() string { _ = "STUB: not implemented"; return "" }

type SelectClause struct {
	*BaseExpr
	Distinct Token
	Fields   []QueryExpression
}

func (sc SelectClause) IsDistinct() bool { _ = "STUB: not implemented"; return false }

func (sc SelectClause) String() string { _ = "STUB: not implemented"; return "" }

type IntoClause struct {
	*BaseExpr
	Variables []Variable
}

func (e IntoClause) String() string { _ = "STUB: not implemented"; return "" }

type FromClause struct {
	*BaseExpr
	Tables []QueryExpression
}

func (f FromClause) String() string { _ = "STUB: not implemented"; return "" }

type WhereClause struct {
	*BaseExpr
	Filter QueryExpression
}

func (w WhereClause) String() string { _ = "STUB: not implemented"; return "" }

type GroupByClause struct {
	*BaseExpr
	Items []QueryExpression
}

func (gb GroupByClause) String() string { _ = "STUB: not implemented"; return "" }

type HavingClause struct {
	*BaseExpr
	Filter QueryExpression
}

func (h HavingClause) String() string { _ = "STUB: not implemented"; return "" }

type OrderByClause struct {
	*BaseExpr
	Items []QueryExpression
}

func (ob OrderByClause) String() string { _ = "STUB: not implemented"; return "" }

type LimitClause struct {
	*BaseExpr
	Type         Token
	Position     Token
	Value        QueryExpression
	Unit         Token
	Restriction  Token
	OffsetClause QueryExpression
}

func (e LimitClause) restrictionString() []string { _ = "STUB: not implemented"; return nil }

func (e LimitClause) String() string { _ = "STUB: not implemented"; return "" }

func (e LimitClause) Percentage() bool { _ = "STUB: not implemented"; return false }

func (e LimitClause) WithTies() bool { _ = "STUB: not implemented"; return false }

type OffsetClause struct {
	*BaseExpr
	Value QueryExpression
	Unit  Token
}

func (e OffsetClause) String() string { _ = "STUB: not implemented"; return "" }

type WithClause struct {
	*BaseExpr
	InlineTables []QueryExpression
}

func (e WithClause) String() string { _ = "STUB: not implemented"; return "" }

type InlineTable struct {
	*BaseExpr
	Recursive Token
	Name      Identifier
	Fields    []QueryExpression
	Query     SelectQuery
}

func (e InlineTable) String() string { _ = "STUB: not implemented"; return "" }

func (e InlineTable) IsRecursive() bool { _ = "STUB: not implemented"; return false }

type Subquery struct {
	*BaseExpr
	Query SelectQuery
}

func (e Subquery) String() string { _ = "STUB: not implemented"; return "" }

type Url struct {
	*BaseExpr
	Raw string
}

func (e Url) String() string { _ = "STUB: not implemented"; return "" }

type TableFunction struct {
	*BaseExpr
	Name string
	Args []QueryExpression
}

func (e TableFunction) String() string { _ = "STUB: not implemented"; return "" }

type FormatSpecifiedFunction struct {
	*BaseExpr
	Type          Token
	FormatElement QueryExpression
	Path          QueryExpression
	Args          []QueryExpression
}

func (e FormatSpecifiedFunction) String() string { _ = "STUB: not implemented"; return "" }

type JsonQuery struct {
	*BaseExpr
	JsonQuery Token
	Query     QueryExpression
	JsonText  QueryExpression
}

func (e JsonQuery) String() string { _ = "STUB: not implemented"; return "" }

type Comparison struct {
	*BaseExpr
	LHS      QueryExpression
	Operator Token
	RHS      QueryExpression
}

func (c Comparison) String() string { _ = "STUB: not implemented"; return "" }

type Is struct {
	*BaseExpr
	LHS      QueryExpression
	RHS      QueryExpression
	Negation Token
}

func (i Is) IsNegated() bool { _ = "STUB: not implemented"; return false }

func (i Is) String() string { _ = "STUB: not implemented"; return "" }

type Between struct {
	*BaseExpr
	LHS      QueryExpression
	Low      QueryExpression
	High     QueryExpression
	Negation Token
}

func (b Between) IsNegated() bool { _ = "STUB: not implemented"; return false }

func (b Between) String() string { _ = "STUB: not implemented"; return "" }

type In struct {
	*BaseExpr
	LHS      QueryExpression
	Values   QueryExpression
	Negation Token
}

func (i In) IsNegated() bool { _ = "STUB: not implemented"; return false }

func (i In) String() string { _ = "STUB: not implemented"; return "" }

type All struct {
	*BaseExpr
	LHS      QueryExpression
	Operator Token
	Values   QueryExpression
}

func (a All) String() string { _ = "STUB: not implemented"; return "" }

type Any struct {
	*BaseExpr
	LHS      QueryExpression
	Operator Token
	Values   QueryExpression
}

func (a Any) String() string { _ = "STUB: not implemented"; return "" }

type Like struct {
	*BaseExpr
	LHS      QueryExpression
	Pattern  QueryExpression
	Negation Token
}

func (l Like) IsNegated() bool { _ = "STUB: not implemented"; return false }

func (l Like) String() string { _ = "STUB: not implemented"; return "" }

type Exists struct {
	*BaseExpr
	Query Subquery
}

func (e Exists) String() string { _ = "STUB: not implemented"; return "" }

type Arithmetic struct {
	*BaseExpr
	LHS      QueryExpression
	Operator Token
	RHS      QueryExpression
}

func (a Arithmetic) String() string { _ = "STUB: not implemented"; return "" }

type UnaryArithmetic struct {
	*BaseExpr
	Operand  QueryExpression
	Operator Token
}

func (e UnaryArithmetic) String() string { _ = "STUB: not implemented"; return "" }

type Logic struct {
	*BaseExpr
	LHS      QueryExpression
	Operator Token
	RHS      QueryExpression
}

func (l Logic) String() string { _ = "STUB: not implemented"; return "" }

type UnaryLogic struct {
	*BaseExpr
	Operand  QueryExpression
	Operator Token
}

func (e UnaryLogic) String() string { _ = "STUB: not implemented"; return "" }

type Concat struct {
	*BaseExpr
	Items []QueryExpression
}

func (c Concat) String() string { _ = "STUB: not implemented"; return "" }

type Function struct {
	*BaseExpr
	Name string
	Args []QueryExpression
	From Token
	For  Token
}

func (e Function) String() string { _ = "STUB: not implemented"; return "" }

type AggregateFunction struct {
	*BaseExpr
	Name     string
	Distinct Token
	Args     []QueryExpression
}

func (e AggregateFunction) String() string { _ = "STUB: not implemented"; return "" }

func (e AggregateFunction) IsDistinct() bool { _ = "STUB: not implemented"; return false }

type Table struct {
	*BaseExpr
	Lateral Token
	Object  QueryExpression
	As      Token
	Alias   QueryExpression
}

func (e Table) String() string { _ = "STUB: not implemented"; return "" }

type Join struct {
	*BaseExpr
	Table     QueryExpression
	JoinTable QueryExpression
	Natural   Token
	JoinType  Token
	Direction Token
	Condition QueryExpression
}

func (j Join) String() string { _ = "STUB: not implemented"; return "" }

type JoinCondition struct {
	*BaseExpr
	On    QueryExpression
	Using []QueryExpression
}

func (jc JoinCondition) String() string { _ = "STUB: not implemented"; return "" }

type Field struct {
	*BaseExpr
	Object QueryExpression
	As     Token
	Alias  QueryExpression
}

func (f Field) String() string { _ = "STUB: not implemented"; return "" }

func (f Field) Name() string { _ = "STUB: not implemented"; return "" }

type AllColumns struct {
	*BaseExpr
}

func (ac AllColumns) String() string { _ = "STUB: not implemented"; return "" }

type Dual struct {
	*BaseExpr
}

func (d Dual) String() string { _ = "STUB: not implemented"; return "" }

type Stdin struct {
	*BaseExpr
}

func (si Stdin) String() string { _ = "STUB: not implemented"; return "" }

type OrderItem struct {
	*BaseExpr
	Value         QueryExpression
	Direction     Token
	NullsPosition Token
}

func (e OrderItem) String() string { _ = "STUB: not implemented"; return "" }

type CaseExpr struct {
	*BaseExpr
	Value QueryExpression
	When  []QueryExpression
	Else  QueryExpression
}

func (e CaseExpr) String() string { _ = "STUB: not implemented"; return "" }

type CaseExprWhen struct {
	*BaseExpr
	Condition QueryExpression
	Result    QueryExpression
}

func (e CaseExprWhen) String() string { _ = "STUB: not implemented"; return "" }

type CaseExprElse struct {
	*BaseExpr
	Result QueryExpression
}

func (e CaseExprElse) String() string { _ = "STUB: not implemented"; return "" }

type ListFunction struct {
	*BaseExpr
	Name     string
	Distinct Token
	Args     []QueryExpression
	OrderBy  QueryExpression
}

func (e ListFunction) String() string { _ = "STUB: not implemented"; return "" }

func (e ListFunction) IsDistinct() bool { _ = "STUB: not implemented"; return false }

type AnalyticFunction struct {
	*BaseExpr
	Name           string
	Distinct       Token
	Args           []QueryExpression
	IgnoreType     Token
	AnalyticClause AnalyticClause
}

func (e AnalyticFunction) String() string { _ = "STUB: not implemented"; return "" }

func (e AnalyticFunction) IsDistinct() bool { _ = "STUB: not implemented"; return false }

func (e AnalyticFunction) IgnoreNulls() bool { _ = "STUB: not implemented"; return false }

type AnalyticClause struct {
	*BaseExpr
	PartitionClause QueryExpression
	OrderByClause   QueryExpression
	WindowingClause QueryExpression
}

func (e AnalyticClause) String() string { _ = "STUB: not implemented"; return "" }

func (e AnalyticClause) PartitionValues() []QueryExpression { _ = "STUB: not implemented"; return nil }

type PartitionClause struct {
	*BaseExpr
	Values []QueryExpression
}

func (e PartitionClause) String() string { _ = "STUB: not implemented"; return "" }

type WindowingClause struct {
	*BaseExpr
	FrameLow  QueryExpression
	FrameHigh QueryExpression
}

func (e WindowingClause) String() string { _ = "STUB: not implemented"; return "" }

type WindowFramePosition struct {
	*BaseExpr
	Direction Token
	Unbounded Token
	Offset    int
}

func (e WindowFramePosition) String() string { _ = "STUB: not implemented"; return "" }

type Variable struct {
	*BaseExpr
	Name string
}

func (v Variable) String() string { _ = "STUB: not implemented"; return "" }

type VariableSubstitution struct {
	*BaseExpr
	Variable Variable
	Value    QueryExpression
}

func (vs VariableSubstitution) String() string { _ = "STUB: not implemented"; return "" }

type VariableAssignment struct {
	*BaseExpr
	Variable Variable
	Value    QueryExpression
}

type VariableDeclaration struct {
	*BaseExpr
	Assignments []VariableAssignment
}

type DisposeVariable struct {
	*BaseExpr
	Variable Variable
}

type EnvironmentVariable struct {
	*BaseExpr
	Name   string
	Quoted bool
}

func (e EnvironmentVariable) String() string { _ = "STUB: not implemented"; return "" }

type RuntimeInformation struct {
	*BaseExpr
	Name string
}

func (e RuntimeInformation) String() string { _ = "STUB: not implemented"; return "" }

type Flag struct {
	*BaseExpr
	Name string
}

func (e Flag) String() string { _ = "STUB: not implemented"; return "" }

type SetEnvVar struct {
	*BaseExpr
	EnvVar EnvironmentVariable
	Value  QueryExpression
}

type UnsetEnvVar struct {
	*BaseExpr
	EnvVar EnvironmentVariable
}

type InsertQuery struct {
	*BaseExpr
	WithClause QueryExpression
	Table      Table
	Fields     []QueryExpression
	ValuesList []QueryExpression
	Query      QueryExpression
}

type UpdateQuery struct {
	*BaseExpr
	WithClause  QueryExpression
	Tables      []QueryExpression
	SetList     []UpdateSet
	FromClause  QueryExpression
	WhereClause QueryExpression
}

type UpdateSet struct {
	*BaseExpr
	Field QueryExpression
	Value QueryExpression
}

type ReplaceQuery struct {
	*BaseExpr
	WithClause QueryExpression
	Table      Table
	Fields     []QueryExpression
	Keys       []QueryExpression
	ValuesList []QueryExpression
	Query      QueryExpression
}

type DeleteQuery struct {
	*BaseExpr
	WithClause  QueryExpression
	Tables      []QueryExpression
	FromClause  FromClause
	WhereClause QueryExpression
}

type CreateTable struct {
	*BaseExpr
	Table       Identifier
	Fields      []QueryExpression
	Query       QueryExpression
	IfNotExists bool
}

type AddColumns struct {
	*BaseExpr
	Table    QueryExpression
	Columns  []ColumnDefault
	Position Expression
}

type ColumnDefault struct {
	*BaseExpr
	Column Identifier
	Value  QueryExpression
}

type ColumnPosition struct {
	*BaseExpr
	Position Token
	Column   QueryExpression
}

type DropColumns struct {
	*BaseExpr
	Table   QueryExpression
	Columns []QueryExpression
}

type RenameColumn struct {
	*BaseExpr
	Table QueryExpression
	Old   QueryExpression
	New   Identifier
}

type SetTableAttribute struct {
	*BaseExpr
	Table     QueryExpression
	Attribute Identifier
	Value     QueryExpression
}

type FunctionDeclaration struct {
	*BaseExpr
	Name       Identifier
	Parameters []VariableAssignment
	Statements []Statement
}

type AggregateDeclaration struct {
	*BaseExpr
	Name       Identifier
	Cursor     Identifier
	Parameters []VariableAssignment
	Statements []Statement
}

type DisposeFunction struct {
	*BaseExpr
	Name Identifier
}

type Return struct {
	*BaseExpr
	Value QueryExpression
}

type Echo struct {
	*BaseExpr
	Value QueryExpression
}

type Print struct {
	*BaseExpr
	Value QueryExpression
}

type Printf struct {
	*BaseExpr
	Format QueryExpression
	Values []QueryExpression
}

type Source struct {
	*BaseExpr
	FilePath QueryExpression
}

type Chdir struct {
	*BaseExpr
	DirPath QueryExpression
}

type Pwd struct {
	*BaseExpr
}

type Reload struct {
	*BaseExpr
	Type Identifier
}

type Execute struct {
	*BaseExpr
	Statements QueryExpression
	Values     []QueryExpression
}

type Syntax struct {
	*BaseExpr
	Keywords []QueryExpression
}

type SetFlag struct {
	*BaseExpr
	Flag  Flag
	Value QueryExpression
}

type AddFlagElement struct {
	*BaseExpr
	Flag  Flag
	Value QueryExpression
}

type RemoveFlagElement struct {
	*BaseExpr
	Flag  Flag
	Value QueryExpression
}

type ShowFlag struct {
	*BaseExpr
	Flag Flag
}

type ShowObjects struct {
	*BaseExpr
	Type Identifier
}

type ShowFields struct {
	*BaseExpr
	Type  Identifier
	Table QueryExpression
}

type If struct {
	*BaseExpr
	Condition  QueryExpression
	Statements []Statement
	ElseIf     []ElseIf
	Else       Else
}

type ElseIf struct {
	*BaseExpr
	Condition  QueryExpression
	Statements []Statement
}

type Else struct {
	*BaseExpr
	Statements []Statement
}

type Case struct {
	*BaseExpr
	Value QueryExpression
	When  []CaseWhen
	Else  CaseElse
}

type CaseWhen struct {
	*BaseExpr
	Condition  QueryExpression
	Statements []Statement
}

type CaseElse struct {
	*BaseExpr
	Statements []Statement
}

type While struct {
	*BaseExpr
	Condition  QueryExpression
	Statements []Statement
}

type WhileInCursor struct {
	*BaseExpr
	WithDeclaration bool
	Variables       []Variable
	Cursor          Identifier
	Statements      []Statement
}

type CursorDeclaration struct {
	*BaseExpr
	Cursor    Identifier
	Query     SelectQuery
	Statement Identifier
}

type OpenCursor struct {
	*BaseExpr
	Cursor Identifier
	Values []ReplaceValue
}

type CloseCursor struct {
	*BaseExpr
	Cursor Identifier
}

type DisposeCursor struct {
	*BaseExpr
	Cursor Identifier
}

type FetchCursor struct {
	*BaseExpr
	Position  FetchPosition
	Cursor    Identifier
	Variables []Variable
}

type FetchPosition struct {
	*BaseExpr
	Position Token
	Number   QueryExpression
}

type CursorStatus struct {
	*BaseExpr
	Cursor   Identifier
	Negation Token
	Type     Token
}

func (e CursorStatus) String() string { _ = "STUB: not implemented"; return "" }

type CursorAttrebute struct {
	*BaseExpr
	Cursor    Identifier
	Attrebute Token
}

func (e CursorAttrebute) String() string { _ = "STUB: not implemented"; return "" }

type ViewDeclaration struct {
	*BaseExpr
	View   Identifier
	Fields []QueryExpression
	Query  QueryExpression
}

type DisposeView struct {
	*BaseExpr
	View QueryExpression
}

type StatementPreparation struct {
	*BaseExpr
	Name      Identifier
	Statement *value.String
}

type ReplaceValue struct {
	*BaseExpr
	Value QueryExpression
	Name  Identifier
}

type ExecuteStatement struct {
	*BaseExpr
	Name   Identifier
	Values []ReplaceValue
}

type DisposeStatement struct {
	*BaseExpr
	Name Identifier
}

type TransactionControl struct {
	*BaseExpr
	Token int
}

type FlowControl struct {
	*BaseExpr
	Token int
}

type Trigger struct {
	*BaseExpr
	Event   Identifier
	Message QueryExpression
	Code    value.Primary
}

type Exit struct {
	*BaseExpr
	Code value.Primary
}

type ExternalCommand struct {
	*BaseExpr
	Command string
}

func putParentheses(s string) string { _ = "STUB: not implemented"; return "" }

func joinWithSpace(s []string) string { _ = "STUB: not implemented"; return "" }

func listQueryExpressions(exprs []QueryExpression) string { _ = "STUB: not implemented"; return "" }

func keyword(token int) string { _ = "STUB: not implemented"; return "" }
