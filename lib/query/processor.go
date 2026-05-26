package query

import (
	"context"
	"time"

	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

type StatementFlow int

const (
	Terminate StatementFlow = iota
	TerminateWithError
	Exit
	Break
	Continue
	Return
)

const StoringResultsContextKey = "sqr"
const StatementReplaceValuesContextKey = "rv"

func ContextForStoringResults(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ContextForPreparedStatement(ctx context.Context, values *ReplaceValues) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type Processor struct {
	Tx             *Transaction
	ReferenceScope *ReferenceScope

	storeResults bool

	returnVal        value.Primary
	measurementStart time.Time
}

func NewProcessor(tx *Transaction) *Processor { _ = "STUB: not implemented"; return nil }

func NewProcessorWithScope(tx *Transaction, scope *ReferenceScope) *Processor {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Processor) NewChildProcessor() *Processor { _ = "STUB: not implemented"; return nil }

func (proc *Processor) Close() { _ = "STUB: not implemented"; return }

func (proc *Processor) Execute(ctx context.Context, statements []parser.Statement) (StatementFlow, error) {
	_ = "STUB: not implemented"
	return *new(StatementFlow), nil
}

func (proc *Processor) execute(ctx context.Context, statements []parser.Statement) (flow StatementFlow, err error) {
	_ = "STUB: not implemented"
	return *new(StatementFlow), nil
}

func (proc *Processor) executeChild(ctx context.Context, statements []parser.Statement) (StatementFlow, error) {
	_ = "STUB: not implemented"
	return *new(StatementFlow), nil
}

func (proc *Processor) ExecuteStatement(ctx context.Context, stmt parser.Statement) (StatementFlow, error) {
	_ = "STUB: not implemented"
	return *new(StatementFlow), nil
}

// Do Nothing

func (proc *Processor) IfStmt(ctx context.Context, stmt parser.If) (StatementFlow, error) {
	_ = "STUB: not implemented"
	return *new(StatementFlow), nil
}

func (proc *Processor) Case(ctx context.Context, stmt parser.Case) (StatementFlow, error) {
	_ = "STUB: not implemented"
	return *new(StatementFlow), nil
}

func (proc *Processor) While(ctx context.Context, stmt parser.While) (StatementFlow, error) {
	_ = "STUB: not implemented"
	return *new(StatementFlow), nil
}

func (proc *Processor) WhileInCursor(ctx context.Context, stmt parser.WhileInCursor) (StatementFlow, error) {
	_ = "STUB: not implemented"
	return *new(StatementFlow), nil
}

func (proc *Processor) ExecExternalCommand(ctx context.Context, stmt parser.ExternalCommand) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Processor) showExecutionTime(ctx context.Context) { _ = "STUB: not implemented"; return }

func (proc *Processor) Log(log string, quiet bool) { _ = "STUB: not implemented"; return }

func (proc *Processor) LogNotice(log string, quiet bool) { _ = "STUB: not implemented"; return }

func (proc *Processor) LogWarn(log string, quiet bool) { _ = "STUB: not implemented"; return }

func (proc *Processor) LogError(log string) { _ = "STUB: not implemented"; return }

func (proc *Processor) AutoCommit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (proc *Processor) Commit(ctx context.Context, expr parser.Expression) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Processor) AutoRollback() error { _ = "STUB: not implemented"; return nil }

func (proc *Processor) Rollback(expr parser.Expression) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Processor) ReleaseResources() error { _ = "STUB: not implemented"; return nil }

func (proc *Processor) ReleaseResourcesWithErrors() error { _ = "STUB: not implemented"; return nil }
