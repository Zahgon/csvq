package syntax

import (
	"github.com/mithrandie/go-text/color"
)

const (
	NameEffect     = "syntax_name"
	KeywordEffect  = "syntax_keyword"
	LinkEffect     = "syntax_link"
	VariableEffect = "syntax_variable"
	FlagEffect     = "syntax_flag"
	ItalicEffect   = "syntax_italic"
	TypeEffect     = "syntax_type"
)

type Definition struct {
	Name        Name
	Group       []Grammar
	Description Description
}

type Element interface {
	Format(p *color.Palette) string
}

type Name string

func (e Name) String() string { _ = "STUB: not implemented"; return "" }

func (e Name) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Grammar []Element

func (e Grammar) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Description struct {
	Template string
	Values   []Element
}

func (e Description) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Keyword string

func (e Keyword) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Link string

func (e Link) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Option []Element

func (e Option) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type FollowingContinuousOption []Element

func (e FollowingContinuousOption) Format(p *color.Palette) string {
	_ = "STUB: not implemented"
	return ""
}

type ContinuousOption []Element

func (e ContinuousOption) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type AnyOne []Element

func (e AnyOne) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Parentheses []Element

func (e Parentheses) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type PlainGroup []Element

func (e PlainGroup) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type ConnectedGroup []Element

func (e ConnectedGroup) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Function struct {
	Name       string
	Args       []Element
	AfterArgs  []Element
	CustomArgs []Element
	Return     Element
}

func (e Function) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

func formatArg(arg Element, p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type ArgWithDefValue struct {
	Arg     Element
	Default Element
}

func (e ArgWithDefValue) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type String string

func (e String) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Integer string

func (e Integer) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Float string

func (e Float) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Identifier string

func (e Identifier) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Datetime string

func (e Datetime) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Boolean string

func (e Boolean) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Ternary string

func (e Ternary) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Null string

func (e Null) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Variable string

func (e Variable) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Flag string

func (e Flag) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Token string

func (e Token) Format(_ *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Italic string

func (e Italic) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }

type Return string

func (e Return) Format(p *color.Palette) string { _ = "STUB: not implemented"; return "" }
