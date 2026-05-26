package query

import (
	"github.com/mithrandie/csvq/lib/file"
	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"

	"github.com/mithrandie/go-text"
	"github.com/mithrandie/go-text/fixedlen"
	"github.com/mithrandie/go-text/json"
)

const (
	TableDelimiter          = "DELIMITER"
	TableDelimiterPositions = "DELIMITER_POSITIONS"
	TableFormat             = "FORMAT"
	TableEncoding           = "ENCODING"
	TableLineBreak          = "LINE_BREAK"
	TableHeader             = "HEADER"
	TableEncloseAll         = "ENCLOSE_ALL"
	TableJsonEscape         = "JSON_ESCAPE"
	TablePrettyPrint        = "PRETTY_PRINT"
)

type ViewType int

const (
	ViewTypeFile ViewType = iota
	ViewTypeTemporaryTable
	ViewTypeStdin
	ViewTypeRemoteObject
	ViewTypeStringObject
	ViewTypeInlineTable
)

var FileAttributeList = []string{
	TableDelimiter,
	TableDelimiterPositions,
	TableFormat,
	TableEncoding,
	TableLineBreak,
	TableHeader,
	TableEncloseAll,
	TableJsonEscape,
	TablePrettyPrint,
}

type TableAttributeUnchangedError struct {
	Path    string
	Message string
}

func NewTableAttributeUnchangedError(fpath string) error { _ = "STUB: not implemented"; return nil }

func (e TableAttributeUnchangedError) Error() string { _ = "STUB: not implemented"; return "" }

type FileInfo struct {
	Path        string
	ArchivePath string

	Format             option.Format
	Delimiter          rune
	DelimiterPositions fixedlen.DelimiterPositions
	JsonQuery          string
	Encoding           text.Encoding
	LineBreak          text.LineBreak
	NoHeader           bool
	EncloseAll         bool
	JsonEscape         json.EscapeType
	PrettyPrint        bool

	SingleLine bool

	Handler *file.Handler

	ForUpdate bool
	ViewType  ViewType

	restorePointHeader    Header
	restorePointRecordSet RecordSet
}

func NewFileInfo(
	filename parser.Identifier,
	repository string,
	options option.ImportOptions,
	defaultFormat option.Format,
) (*FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTemporaryTableFileInfo(name string) *FileInfo { _ = "STUB: not implemented"; return nil }

func NewStdinFileInfo(filePath string, importOptions option.ImportOptions, exportOptions option.ExportOptions) *FileInfo {
	_ = "STUB: not implemented"
	return nil
}

func NewInlineFileInfo(filePath string, importOptions option.ImportOptions, exportOptions option.ExportOptions) *FileInfo {
	_ = "STUB: not implemented"
	return nil
}

func (f *FileInfo) SetAllDefaultFileInfoAttributes(importOptions option.ImportOptions, exportOptions option.ExportOptions) {
	_ = "STUB: not implemented"
	return
}

func (f *FileInfo) SetDefaultFileInfoAttributes(importOptions option.ImportOptions, exportOptions option.ExportOptions) {
	_ = "STUB: not implemented"
	return
}

func (f *FileInfo) IsUpdatable() bool { _ = "STUB: not implemented"; return false }

func (f *FileInfo) SetDelimiter(s string) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) SetDelimiterPositions(s string) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) SetFormat(s string) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) SetEncoding(s string) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) SetLineBreak(s string) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) SetNoHeader(b bool) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) SetEncloseAll(b bool) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) SetJsonEscape(s string) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) SetPrettyPrint(b bool) error { _ = "STUB: not implemented"; return nil }

func (f *FileInfo) IsFile() bool { _ = "STUB: not implemented"; return false }

func (f *FileInfo) IsTemporaryTable() bool { _ = "STUB: not implemented"; return false }

func (f *FileInfo) IsStdin() bool { _ = "STUB: not implemented"; return false }

func (f *FileInfo) IsInMemoryTable() bool { _ = "STUB: not implemented"; return false }

func (f *FileInfo) IsRemoteObject() bool { _ = "STUB: not implemented"; return false }

func (f *FileInfo) IsStringObject() bool { _ = "STUB: not implemented"; return false }

func (f *FileInfo) IsInlineTable() bool { _ = "STUB: not implemented"; return false }

func (f *FileInfo) IdentifiedPath() string { _ = "STUB: not implemented"; return "" }

func (f *FileInfo) ExportOptions(tx *Transaction) option.ExportOptions {
	_ = "STUB: not implemented"
	return *new(option.ExportOptions)
}

func SearchFilePath(filename parser.Identifier, repository string, options option.ImportOptions, defaultFormat option.Format) (string, option.Format, error) {
	_ = "STUB: not implemented"
	return "", *new(option.Format), nil
}

// AutoSelect

func SearchCSVFilePath(filename parser.Identifier, repository string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SearchJsonFilePath(filename parser.Identifier, repository string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SearchJsonlFilePath(filename parser.Identifier, repository string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SearchFixedLengthFilePath(filename parser.Identifier, repository string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SearchLTSVFilePath(filename parser.Identifier, repository string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SearchFilePathFromAllTypes(filename parser.Identifier, repository string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SearchFilePathWithExtType(filename parser.Identifier, repository string, extTypes []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func NewFileInfoForCreate(filename parser.Identifier, repository string, delimiter rune, encoding text.Encoding) (*FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateFilePath(filename parser.Identifier, repository string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
