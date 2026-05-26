package query

import (
	"context"

	"github.com/mithrandie/csvq/lib/parser"
)

func FetchCursor(ctx context.Context, scope *ReferenceScope, name parser.Identifier, fetchPosition parser.FetchPosition, vars []parser.Variable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DeclareView(ctx context.Context, scope *ReferenceScope, expr parser.ViewDeclaration) error {
	_ = "STUB: not implemented"
	return nil
}

func Select(ctx context.Context, scope *ReferenceScope, query parser.SelectQuery) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func selectEntity(ctx context.Context, scope *ReferenceScope, expr parser.QueryExpression, forUpdate bool) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func selectSetEntity(ctx context.Context, scope *ReferenceScope, expr parser.QueryExpression, forUpdate bool) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func selectSet(ctx context.Context, scope *ReferenceScope, set parser.SelectSet, forUpdate bool) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func selectSetForRecursion(ctx context.Context, scope *ReferenceScope, view *View, set parser.SelectSet, forUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func Insert(ctx context.Context, scope *ReferenceScope, query parser.InsertQuery) (*FileInfo, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func Update(ctx context.Context, scope *ReferenceScope, query parser.UpdateQuery) ([]*FileInfo, []int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func Replace(ctx context.Context, scope *ReferenceScope, query parser.ReplaceQuery) (*FileInfo, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func Delete(ctx context.Context, scope *ReferenceScope, query parser.DeleteQuery) ([]*FileInfo, []int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTable(ctx context.Context, scope *ReferenceScope, query parser.CreateTable) (*FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddColumns(ctx context.Context, scope *ReferenceScope, query parser.AddColumns) (*FileInfo, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

//parser.AFTER

func DropColumns(ctx context.Context, scope *ReferenceScope, query parser.DropColumns) (*FileInfo, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func RenameColumn(ctx context.Context, scope *ReferenceScope, query parser.RenameColumn) (*FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetTableAttribute(ctx context.Context, scope *ReferenceScope, query parser.SetTableAttribute) (*FileInfo, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
