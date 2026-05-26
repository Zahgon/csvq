package option

import (
	"github.com/mithrandie/go-text/color"
)

const (
	NoEffect         = ""
	LableEffect      = "label"
	NumberEffect     = "number"
	StringEffect     = "string"
	BooleanEffect    = "boolean"
	TernaryEffect    = "ternary"
	DatetimeEffect   = "datetime"
	NullEffect       = "null"
	ObjectEffect     = "object"
	AttributeEffect  = "attribute"
	IdentifierEffect = "identifier"
	ValueEffect      = "value"
	EmphasisEffect   = "emphasis"
	PromptEffect     = "prompt"
	ErrorEffect      = "error"
	WarnEffect       = "warn"
	NoticeEffect     = "notice"
)

func NewPalette(env *Environment) (*color.Palette, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
