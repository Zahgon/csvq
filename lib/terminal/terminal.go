package terminal

import (
	"bytes"
	"context"

	"github.com/mithrandie/csvq/lib/excmd"
	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/query"
)

const (
	DefaultPrompt           string = "csvq > "
	DefaultContinuousPrompt string = "     > "
)

type PromptEvaluationError struct {
	Message string
}

func NewPromptEvaluationError(message string) error { _ = "STUB: not implemented"; return nil }

func (e PromptEvaluationError) Error() string { _ = "STUB: not implemented"; return "" }

type PromptElement struct {
	Text string
	Type excmd.ElementType
}

type Prompt struct {
	scope              *query.ReferenceScope
	sequence           []PromptElement
	continuousSequence []PromptElement

	buf bytes.Buffer
}

func NewPrompt(scope *query.ReferenceScope) *Prompt { _ = "STUB: not implemented"; return nil }

func (p *Prompt) LoadConfig() error { _ = "STUB: not implemented"; return nil }

func (p *Prompt) RenderPrompt(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Prompt) RenderContinuousPrompt(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Prompt) Render(ctx context.Context, sequence []PromptElement) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Prompt) evaluate(ctx context.Context, expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Prompt) StripEscapeSequence(s string) string { _ = "STUB: not implemented"; return "" }
