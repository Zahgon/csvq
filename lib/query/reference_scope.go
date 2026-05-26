package query

import (
	"context"
	"sync"
	"time"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/ternary"
)

const LimitToUseFieldIndexSliceChache = 8

var blockScopePool = sync.Pool{
	New: func() interface{} {
		return NewBlockScope()
	},
}

func GetBlockScope() BlockScope { _ = "STUB: not implemented"; return *new(BlockScope) }

func PutBlockScope(scope BlockScope) { _ = "STUB: not implemented"; return }

var nodeScopePool = sync.Pool{
	New: func() interface{} {
		return NewNodeScope()
	},
}

func GetNodeScope() NodeScope { _ = "STUB: not implemented"; return *new(NodeScope) }

func PutNodeScope(scope NodeScope) { _ = "STUB: not implemented"; return }

type BlockScope struct {
	Variables       VariableMap
	TemporaryTables ViewMap
	Cursors         CursorMap
	Functions       UserDefinedFunctionMap
}

func NewBlockScope() BlockScope { _ = "STUB: not implemented"; return *new(BlockScope) }

func (scope BlockScope) Clear() { _ = "STUB: not implemented"; return }

type NodeScope struct {
	inlineTables InlineTableMap
	aliases      AliasMap
}

func NewNodeScope() NodeScope { _ = "STUB: not implemented"; return *new(NodeScope) }

func (scope NodeScope) Clear() { _ = "STUB: not implemented"; return }

type ReferenceRecord struct {
	view        *View
	recordIndex int

	cache *FieldIndexCache
}

func NewReferenceRecord(view *View, recordIdx int, cacheLen int) ReferenceRecord {
	_ = "STUB: not implemented"
	return *new(ReferenceRecord)
}

func (r *ReferenceRecord) IsInRange() bool { _ = "STUB: not implemented"; return false }

type FieldIndexCache struct {
	limitToUseSlice int
	m               map[parser.QueryExpression]int
	exprs           []parser.QueryExpression
	indices         []int
}

func NewFieldIndexCache(initCap int, limitToUseSlice int) *FieldIndexCache {
	_ = "STUB: not implemented"
	return nil
}

func (c *FieldIndexCache) Get(expr parser.QueryExpression) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (c *FieldIndexCache) Add(expr parser.QueryExpression, idx int) {
	_ = "STUB: not implemented"
	return
}

type ReferenceScope struct {
	Tx *Transaction

	Blocks []BlockScope
	nodes  []NodeScope

	cachedFilePath map[string]string
	now            time.Time

	Records []ReferenceRecord

	RecursiveTable   *parser.InlineTable
	RecursiveTmpView *View
	RecursiveCount   *int64
}

func NewReferenceScope(tx *Transaction) *ReferenceScope { _ = "STUB: not implemented"; return nil }

func NewReferenceScopeWithBlock(tx *Transaction, scope BlockScope) *ReferenceScope {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) CreateScopeForRecordEvaluation(view *View, recordIndex int) *ReferenceScope {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) CreateScopeForSequentialEvaluation(view *View) *ReferenceScope {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) CreateScopeForAnalytics() *ReferenceScope {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) createScope(referenceRecords []ReferenceRecord) *ReferenceScope {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) CreateChild() *ReferenceScope { _ = "STUB: not implemented"; return nil }

func (rs *ReferenceScope) CreateNode() *ReferenceScope { _ = "STUB: not implemented"; return nil }

func (rs *ReferenceScope) Global() BlockScope { _ = "STUB: not implemented"; return *new(BlockScope) }

func (rs *ReferenceScope) CurrentBlock() BlockScope {
	_ = "STUB: not implemented"
	return *new(BlockScope)
}

func (rs *ReferenceScope) ClearCurrentBlock() { _ = "STUB: not implemented"; return }

func (rs *ReferenceScope) CloseCurrentBlock() { _ = "STUB: not implemented"; return }

func (rs *ReferenceScope) CloseCurrentNode() { _ = "STUB: not implemented"; return }

func (rs *ReferenceScope) NextRecord() bool { _ = "STUB: not implemented"; return false }

func (rs *ReferenceScope) FilePathExists(identifier string) bool {
	_ = "STUB: not implemented"
	return false
}

func (rs *ReferenceScope) StoreFilePath(identifier string, fpath string) {
	_ = "STUB: not implemented"
	return
}

func (rs *ReferenceScope) LoadFilePath(identifier string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (rs *ReferenceScope) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (rs *ReferenceScope) DeclareVariable(ctx context.Context, expr parser.VariableDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) DeclareVariableDirectly(variable parser.Variable, val value.Primary) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) GetVariable(expr parser.Variable) (val value.Primary, err error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func (rs *ReferenceScope) SubstituteVariable(ctx context.Context, expr parser.VariableSubstitution) (val value.Primary, err error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func (rs *ReferenceScope) SubstituteVariableDirectly(variable parser.Variable, val value.Primary) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func (rs *ReferenceScope) DisposeVariable(expr parser.Variable) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) AllVariables() VariableMap {
	_ = "STUB: not implemented"
	return *new(VariableMap)
}

func (rs *ReferenceScope) TemporaryTableExists(identifier string) bool {
	_ = "STUB: not implemented"
	return false
}

func (rs *ReferenceScope) GetTemporaryTable(identifier parser.Identifier) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *ReferenceScope) GetTemporaryTableWithInternalId(ctx context.Context, identifier parser.Identifier, flags *option.Flags) (view *View, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *ReferenceScope) SetTemporaryTable(view *View) { _ = "STUB: not implemented"; return }

func (rs *ReferenceScope) ReplaceTemporaryTable(view *View) { _ = "STUB: not implemented"; return }

func (rs *ReferenceScope) DisposeTemporaryTable(name parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) StoreTemporaryTable(session *Session, uncomittedViews map[string]*FileInfo) []string {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) RestoreTemporaryTable(uncomittedViews map[string]*FileInfo) []string {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) AllTemporaryTables() ViewMap {
	_ = "STUB: not implemented"
	return *new(ViewMap)
}

func (rs *ReferenceScope) DeclareCursor(expr parser.CursorDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) AddPseudoCursor(name parser.Identifier, values []value.Primary) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) DisposeCursor(name parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) OpenCursor(ctx context.Context, name parser.Identifier, values []parser.ReplaceValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) CloseCursor(name parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) FetchCursor(name parser.Identifier, position int, number int) ([]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *ReferenceScope) CursorIsOpen(name parser.Identifier) (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}

func (rs *ReferenceScope) CursorIsInRange(name parser.Identifier) (ternary.Value, error) {
	_ = "STUB: not implemented"
	return *new(ternary.Value), nil
}

func (rs *ReferenceScope) CursorCount(name parser.Identifier) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rs *ReferenceScope) AllCursors() CursorMap { _ = "STUB: not implemented"; return *new(CursorMap) }

func (rs *ReferenceScope) DeclareFunction(expr parser.FunctionDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) DeclareAggregateFunction(expr parser.AggregateDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) GetFunction(expr parser.QueryExpression, name string) (*UserDefinedFunction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *ReferenceScope) DisposeFunction(name parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) AllFunctions() (UserDefinedFunctionMap, UserDefinedFunctionMap) {
	_ = "STUB: not implemented"
	return *new(UserDefinedFunctionMap), *new(UserDefinedFunctionMap)
}

func (rs *ReferenceScope) SetInlineTable(ctx context.Context, inlineTable parser.InlineTable) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) GetInlineTable(name parser.Identifier) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *ReferenceScope) StoreInlineTable(name parser.Identifier, view *View) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) InlineTableExists(name parser.Identifier) bool {
	_ = "STUB: not implemented"
	return false
}

func (rs *ReferenceScope) LoadInlineTable(ctx context.Context, clause parser.WithClause) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) AddAlias(alias parser.Identifier, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReferenceScope) GetAlias(alias parser.Identifier) (path string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
