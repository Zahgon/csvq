package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/doc"
	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
)

type ObjectStatus int

const (
	ObjectFixed ObjectStatus = iota
	ObjectCreated
	ObjectUpdated
	ReadOnly
)

const IgnoredFlagPrefix = "(ignored) "

const (
	ReloadConfig = "CONFIG"
)

const (
	ShowTables     = "TABLES"
	ShowViews      = "VIEWS"
	ShowCursors    = "CURSORS"
	ShowFunctions  = "FUNCTIONS"
	ShowStatements = "STATEMENTS"
	ShowFlags      = "FLAGS"
	ShowEnv        = "ENV"
	ShowRuninfo    = "RUNINFO"
)

var ShowObjectList = []string{
	ShowTables,
	ShowViews,
	ShowCursors,
	ShowFunctions,
	ShowStatements,
	ShowFlags,
	ShowEnv,
	ShowRuninfo,
}

func Echo(ctx context.Context, scope *ReferenceScope, expr parser.Echo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func Print(ctx context.Context, scope *ReferenceScope, expr parser.Print) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func Printf(ctx context.Context, scope *ReferenceScope, expr parser.Printf) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func Source(ctx context.Context, scope *ReferenceScope, expr parser.Source) ([]parser.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadContentsFromFile(ctx context.Context, tx *Transaction, fpath parser.Identifier) (content string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func LoadStatementsFromFile(ctx context.Context, tx *Transaction, fpath parser.Identifier) (statements []parser.Statement, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseExecuteStatements(ctx context.Context, scope *ReferenceScope, expr parser.Execute) ([]parser.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetFlag(ctx context.Context, scope *ReferenceScope, expr parser.SetFlag) error {
	_ = "STUB: not implemented"
	return nil
}

func AddFlagElement(ctx context.Context, scope *ReferenceScope, expr parser.AddFlagElement) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveFlagElement(ctx context.Context, scope *ReferenceScope, expr parser.RemoveFlagElement) error {
	_ = "STUB: not implemented"
	return nil
}

func ShowFlag(tx *Transaction, expr parser.ShowFlag) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func showFlag(tx *Transaction, flagName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func ShowObjects(scope *ReferenceScope, expr parser.ShowObjects) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func writeTableAttribute(w *doc.Writer, flags *option.Flags, info *FileInfo) {
	_ = "STUB: not implemented"
	return
}

func writeFields(w *doc.Writer, fields []string) { _ = "STUB: not implemented"; return }

func writeFunctions(w *doc.Writer, funcs UserDefinedFunctionMap) { _ = "STUB: not implemented"; return }

func ShowFields(ctx context.Context, scope *ReferenceScope, expr parser.ShowFields) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func writeFieldList(w *doc.Writer, fields []string) { _ = "STUB: not implemented"; return }

func writeQuery(w *doc.Writer, s string) { _ = "STUB: not implemented"; return }

func SetEnvVar(ctx context.Context, scope *ReferenceScope, expr parser.SetEnvVar) error {
	_ = "STUB: not implemented"
	return nil
}

func UnsetEnvVar(expr parser.UnsetEnvVar) error { _ = "STUB: not implemented"; return nil }

func Chdir(ctx context.Context, scope *ReferenceScope, expr parser.Chdir) error {
	_ = "STUB: not implemented"
	return nil
}

func Pwd(expr parser.Pwd) (string, error) { _ = "STUB: not implemented"; return "", nil }

func Reload(ctx context.Context, tx *Transaction, expr parser.Reload) error {
	_ = "STUB: not implemented"
	return nil
}

func Syntax(ctx context.Context, scope *ReferenceScope, expr parser.Syntax) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
