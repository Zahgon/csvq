//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows

package terminal

import (
	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/query"

	"github.com/mithrandie/readline-csvq"
)

const (
	dummySubquery    = "____subquery____"
	dummyTableObject = "____table_object____"
	dummyTable       = "____table____"
)

var statementPrefix = []string{
	"WITH",
	"SELECT",
	"INSERT",
	"UPDATE",
	"REPLACE",
	"DELETE",
	"CREATE",
	"ALTER",
	"DECLARE",
	"PREPARE",
	"VAR",
	"SET",
	"UNSET",
	"ADD",
	"REMOVE",
	"ECHO",
	"PRINT",
	"PRINTF",
	"CHDIR",
	"EXECUTE",
	"SHOW",
	"SOURCE",
	"SYNTAX",
	"RELOAD",
}

var singleCommandStatement = []string{
	"COMMIT",
	"ROLLBACK",
	"EXIT",
	"PWD",
}

var delimiterCandidates = []string{
	"','",
	"'\\t'",
}

var delimiterPositionsCandidates = []string{
	"'SPACES'",
	"'S[]'",
	"'[]'",
}

var joinCandidates = []string{
	"JOIN",
	"CROSS",
	"INNER",
	"FULL",
	"LEFT",
	"RIGHT",
	"NATURAL",
}

var tableObjectCandidates = []string{
	"CSV()",
	"FIXED()",
	"JSON()",
	"JSONL()",
	"LTSV()",
}

var exportEncodingsCandidates = []string{
	"SJIS",
	"UTF16",
	"UTF16BE",
	"UTF16BEM",
	"UTF16LE",
	"UTF16LEM",
	"UTF8",
	"UTF8M",
}

type ReadlineListener struct {
	scanner parser.Scanner
}

func skipInputtingEnclosure(line []rune, pos int) []rune { _ = "STUB: not implemented"; return nil }

func completeEnclosure(line []rune, pos int, rightEnclosure rune) []rune {
	_ = "STUB: not implemented"
	return nil
}

func (l ReadlineListener) OnChange(line []rune, pos int, key rune) ([]rune, int, bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

type Completer struct {
	completer *readline.PrefixCompleter
	scope     *query.ReferenceScope

	flagList      []string
	runinfoList   []string
	funcs         []string
	aggFuncs      []string
	analyticFuncs []string

	constants []string

	statementList    []string
	userFuncs        []string
	userAggFuncs     []string
	userFuncList     []string
	viewList         []string
	cursorList       []string
	funcList         []string
	aggFuncList      []string
	analyticFuncList []string
	varList          []string
	envList          []string
	enclosedEnvList  []string
	allColumns       []string
	tableColumns     map[string][]string

	tokens            []parser.Token
	lastIdx           int
	selectIntoEnabled bool

	isInAndAfterSelect bool
}

func NewCompleter(scope *query.ReferenceScope) *Completer { _ = "STUB: not implemented"; return nil }

func (c *Completer) Do(line []rune, pos int, index int) (readline.CandidateList, int) {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList), 0
}

func (c *Completer) Update() { _ = "STUB: not implemented"; return }

func (c *Completer) updateStatements() { _ = "STUB: not implemented"; return }

func (c *Completer) updateViews() { _ = "STUB: not implemented"; return }

func (c *Completer) updateCursors() { _ = "STUB: not implemented"; return }

func (c *Completer) updateFunctions() { _ = "STUB: not implemented"; return }

func (c *Completer) updateVariables() { _ = "STUB: not implemented"; return }

func (c *Completer) updateEnvironmentVariables() { _ = "STUB: not implemented"; return }

func (c *Completer) updateAllColumns() { _ = "STUB: not implemented"; return }

func (c *Completer) GetStatementPrefix(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) Statements(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) TableObjectArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) FunctionArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) substringArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

//Do nothing

func (c *Completer) functionArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) WithArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) combineTableAlias(fromIdx int) { _ = "STUB: not implemented"; return }

func (c *Completer) allTableCandidates(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) allTableCandidatesForUpdate(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) allTableCandidatesWithSpaceForUpdate(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) fromClause(i int, line string, origLine string, index int) (tables readline.CandidateList, customList readline.CandidateList, restrict bool) {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList), *new(readline.CandidateList), false
}

//OK

func (c *Completer) whereClause(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) SelectArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

//Do nothing

func (c *Completer) InsertArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) UpdateArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) ReplaceArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) DeleteArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) CreateArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) AlterArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

//Column

//Column

func (c *Completer) DeclareArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) PrepareArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) FetchArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) SetArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) UsingArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) AddFlagArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) RemoveFlagArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) DisposeArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) ShowArgs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) SearchAllTablesWithSpace(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) SearchAllTables(line string, _ string, _ int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) SearchExecutableFiles(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) SearchDirs(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) SearchValuesWithSpace(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) SearchValues(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) CursorStatus(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) caseExpressionIsNotEnclosed() bool { _ = "STUB: not implemented"; return false }

func (c *Completer) CaseExpression(line string, origLine string, index int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) EncloseQuotation(line string, origLine string, _ int) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) ListFiles(path string, includeExt []string, repository string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *Completer) AllColumnList() []string { _ = "STUB: not implemented"; return nil }

func (c *Completer) ColumnList(tableName string, repository string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (*Completer) columnList(view *query.View) []string { _ = "STUB: not implemented"; return nil }

func (c *Completer) completeArgs(
	line string,
	origLine string,
	index int,
	fn func(i int) (keywords []string, customList readline.CandidateList, breakLoop bool),
) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) UpdateTokens(line string, origLine string) { _ = "STUB: not implemented"; return }

func (c *Completer) SetLastIndex(line string) { _ = "STUB: not implemented"; return }

func (c *Completer) setCursorIsInAndAfterSelect() { _ = "STUB: not implemented"; return }

func (c *Completer) searchStartIndex() int { _ = "STUB: not implemented"; return 0 }

func (c *Completer) combineSubqueryTokens() { _ = "STUB: not implemented"; return }

func (c *Completer) combineTableObject() { _ = "STUB: not implemented"; return }

func (c *Completer) combineFunction() { _ = "STUB: not implemented"; return }

func (c *Completer) isTableObject(token parser.Token) bool { _ = "STUB: not implemented"; return false }

func (c *Completer) isFunction(token parser.Token) bool { _ = "STUB: not implemented"; return false }

func (c *Completer) BracketIsEnclosed() bool { _ = "STUB: not implemented"; return false }

func (c *Completer) candidateList(list []string, appendSpace bool) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) identifierList(list []string, appendSpace bool) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) filteredCandidateList(line string, list []string, appendSpace bool) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) candidate(candidate string, appendSpace bool) readline.Candidate {
	_ = "STUB: not implemented"
	return *new(readline.Candidate)
}

func (c *Completer) identifier(candidate string, appendSpace bool) readline.Candidate {
	_ = "STUB: not implemented"
	return *new(readline.Candidate)
}

func (c *Completer) aggregateFunctionCandidateList(line string) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) analyticFunctionCandidateList(line string) readline.CandidateList {
	_ = "STUB: not implemented"
	return *new(readline.CandidateList)
}

func (c *Completer) environmentVariableList(line string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *Completer) tableFormatList() []string { _ = "STUB: not implemented"; return nil }

func (c *Completer) importFormatList() []string { _ = "STUB: not implemented"; return nil }

func (c *Completer) encodingList() []string { _ = "STUB: not implemented"; return nil }

func (c *Completer) lineBreakList() []string { _ = "STUB: not implemented"; return nil }

func (c *Completer) jsonEscapeTypeList() []string { _ = "STUB: not implemented"; return nil }
