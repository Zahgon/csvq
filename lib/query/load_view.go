package query

import (
	"context"
	"io"

	"github.com/mithrandie/csvq/lib/file"
	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"

	"github.com/mithrandie/go-text"
)

const fileLoadingPreparedRecordSetCap = 300
const fileLoadingBuffer = 300

const inlineTablePrefix = "@__io__"

type RecordReader interface {
	Read() ([]text.RawText, error)
}

func isTableObjectAsDataObject(tablePath parser.QueryExpression) bool {
	_ = "STUB: not implemented"
	return false
}

func isTableObjectAsURL(tablePath parser.QueryExpression) bool {
	_ = "STUB: not implemented"
	return false
}

func LoadView(ctx context.Context, scope *ReferenceScope, tables []parser.QueryExpression, forUpdate bool, useInternalId bool) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadViewFromTableIdentifier(ctx context.Context, scope *ReferenceScope, table parser.QueryExpression, forUpdate bool, useInternalId bool) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadView(ctx context.Context, scope *ReferenceScope, tableExpr parser.QueryExpression, forUpdate bool, useInternalId bool) (view *View, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JSON_INLINE, JSON_TABLE

func joinViews(ctx context.Context, scope *ReferenceScope, view *View, joinView *View, join parser.Join) error {
	_ = "STUB: not implemented"
	return nil
}

func loadObjectFromStdin(
	ctx context.Context,
	scope *ReferenceScope,
	stdin parser.Stdin,
	tableName parser.Identifier,
	forUpdate bool,
	useInternalId bool,
	options option.ImportOptions,
) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadObjectFromString(
	ctx context.Context,
	scope *ReferenceScope,
	data string,
	tablePath parser.QueryExpression,
	tableName parser.Identifier,
	options option.ImportOptions,
) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadDataObject(
	ctx context.Context,
	scope *ReferenceScope,
	dataObject DataObject,
	tablePath parser.QueryExpression,
	tableName parser.Identifier,
	options option.ImportOptions,
) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadHttpObject(
	ctx context.Context,
	scope *ReferenceScope,
	httpObject HttpObject,
	tablePath parser.QueryExpression,
	tableName parser.Identifier,
	options option.ImportOptions,
) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadInlineObjectFromFile(
	ctx context.Context,
	scope *ReferenceScope,
	tableIdentifier parser.Identifier,
	tableName parser.Identifier,
	options option.ImportOptions,
) (view *View, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadObjectFromFile(
	ctx context.Context,
	scope *ReferenceScope,
	fileIdentifier parser.Identifier,
	tableName parser.Identifier,
	forUpdate bool,
	useInternalId bool,
	options option.ImportOptions,
) (view *View, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadObject(
	ctx context.Context,
	scope *ReferenceScope,
	tablePath parser.QueryExpression,
	tableName parser.Identifier,
	forUpdate bool,
	useInternalId bool,
	isInlineObject bool,
	options option.ImportOptions,
) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cacheViewFromFile(
	ctx context.Context,
	scope *ReferenceScope,
	fileIdentifier parser.Identifier,
	forUpdate bool,
	options option.ImportOptions,
) (filePath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Uncommitted newly created tables are not yet created as files, so they need to be checked if they are
// registered in the cache with CreateFilePath(), instead of SearchFilePath().

func loadViewFromFile(ctx context.Context, flags *option.Flags, fp io.Reader, fileInfo *FileInfo, options option.ImportOptions, expr parser.QueryExpression) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadViewFromFixedLengthTextFile(ctx context.Context, fp *file.Reader, fileInfo *FileInfo, withoutNull bool, expr parser.QueryExpression) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadViewFromCSVFile(ctx context.Context, fp *file.Reader, fileInfo *FileInfo, allowUnevenFields bool, withoutNull bool, expr parser.QueryExpression) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadViewFromLTSVFile(ctx context.Context, flags *option.Flags, fp *file.Reader, fileInfo *FileInfo, withoutNull bool, expr parser.QueryExpression) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readRecordSet(ctx context.Context, reader RecordReader, fileSize int64) (RecordSet, error) {
	_ = "STUB: not implemented"
	return *new(RecordSet), nil
}

// Row data sent.

func loadViewFromJsonFile(fp *file.Reader, fileInfo *FileInfo, expr parser.QueryExpression) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadViewFromJsonLinesFile(ctx context.Context, flags *option.Flags, fp *file.Reader, fileInfo *FileInfo, expr parser.QueryExpression) (*View, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Row data sent.
