package option

import (
	"time"

	"github.com/mithrandie/go-text"
	txjson "github.com/mithrandie/go-text/json"
)

const (
	VariableSign            = "@"
	FlagSign                = "@@"
	EnvironmentVariableSign = "@%"
	RuntimeInformationSign  = "@#"
)
const DelimitAutomatically = "SPACES"

const (
	RepositoryFlag               = "REPOSITORY"
	TimezoneFlag                 = "TIMEZONE"
	DatetimeFormatFlag           = "DATETIME_FORMAT"
	AnsiQuotesFlag               = "ANSI_QUOTES"
	StrictEqualFlag              = "STRICT_EQUAL"
	WaitTimeoutFlag              = "WAIT_TIMEOUT"
	ImportFormatFlag             = "IMPORT_FORMAT"
	DelimiterFlag                = "DELIMITER"
	AllowUnevenFieldsFlag        = "ALLOW_UNEVEN_FIELDS"
	DelimiterPositionsFlag       = "DELIMITER_POSITIONS"
	JsonQueryFlag                = "JSON_QUERY"
	EncodingFlag                 = "ENCODING"
	NoHeaderFlag                 = "NO_HEADER"
	WithoutNullFlag              = "WITHOUT_NULL"
	StripEndingLineBreakFlag     = "STRIP_ENDING_LINE_BREAK"
	FormatFlag                   = "FORMAT"
	ExportEncodingFlag           = "WRITE_ENCODING"
	ExportDelimiterFlag          = "WRITE_DELIMITER"
	ExportDelimiterPositionsFlag = "WRITE_DELIMITER_POSITIONS"
	WithoutHeaderFlag            = "WITHOUT_HEADER"
	LineBreakFlag                = "LINE_BREAK"
	EncloseAllFlag               = "ENCLOSE_ALL"
	JsonEscapeFlag               = "JSON_ESCAPE"
	PrettyPrintFlag              = "PRETTY_PRINT"
	ScientificNotationFlag       = "SCIENTIFIC_NOTATION"
	EastAsianEncodingFlag        = "EAST_ASIAN_ENCODING"
	CountDiacriticalSignFlag     = "COUNT_DIACRITICAL_SIGN"
	CountFormatCodeFlag          = "COUNT_FORMAT_CODE"
	ColorFlag                    = "COLOR"
	QuietFlag                    = "QUIET"
	LimitRecursion               = "LIMIT_RECURSION"
	CPUFlag                      = "CPU"
	StatsFlag                    = "STATS"
)

var FlagList = []string{
	RepositoryFlag,
	TimezoneFlag,
	DatetimeFormatFlag,
	AnsiQuotesFlag,
	StrictEqualFlag,
	WaitTimeoutFlag,
	ImportFormatFlag,
	DelimiterFlag,
	AllowUnevenFieldsFlag,
	DelimiterPositionsFlag,
	JsonQueryFlag,
	EncodingFlag,
	NoHeaderFlag,
	WithoutNullFlag,
	StripEndingLineBreakFlag,
	FormatFlag,
	ExportEncodingFlag,
	ExportDelimiterFlag,
	ExportDelimiterPositionsFlag,
	WithoutHeaderFlag,
	LineBreakFlag,
	EncloseAllFlag,
	JsonEscapeFlag,
	PrettyPrintFlag,
	ScientificNotationFlag,
	EastAsianEncodingFlag,
	CountDiacriticalSignFlag,
	CountFormatCodeFlag,
	ColorFlag,
	QuietFlag,
	LimitRecursion,
	CPUFlag,
	StatsFlag,
}

type Format int

const (
	AutoSelect Format = -1 + iota
	CSV
	TSV
	FIXED
	JSON
	JSONL
	LTSV
	GFM
	ORG
	BOX
	TEXT
)

var FormatLiteral = map[Format]string{
	CSV:   "CSV",
	TSV:   "TSV",
	FIXED: "FIXED",
	JSON:  "JSON",
	JSONL: "JSONL",
	LTSV:  "LTSV",
	GFM:   "GFM",
	ORG:   "ORG",
	BOX:   "BOX",
	TEXT:  "TEXT",
}

func (f Format) String() string { _ = "STUB: not implemented"; return "" }

var ImportFormats = []Format{
	CSV,
	TSV,
	FIXED,
	JSON,
	JSONL,
	LTSV,
}

var JsonEscapeTypeLiteral = map[txjson.EscapeType]string{
	txjson.Backslash:        "BACKSLASH",
	txjson.HexDigits:        "HEX",
	txjson.AllWithHexDigits: "HEXALL",
}

func JsonEscapeTypeToString(escapeType txjson.EscapeType) string {
	_ = "STUB: not implemented"
	return ""
}

const (
	CsvExt      = ".csv"
	TsvExt      = ".tsv"
	JsonExt     = ".json"
	JsonlExt    = ".jsonl"
	LtsvExt     = ".ltsv"
	GfmExt      = ".md"
	OrgExt      = ".org"
	SqlExt      = ".sql"
	CsvqProcExt = ".cql"
	TextExt     = ".txt"
)

type ImportOptions struct {
	Format             Format
	Delimiter          rune
	AllowUnevenFields  bool
	DelimiterPositions []int
	SingleLine         bool
	JsonQuery          string
	Encoding           text.Encoding
	NoHeader           bool
	WithoutNull        bool
}

func (ops ImportOptions) Copy() ImportOptions {
	_ = "STUB: not implemented"
	return *new(ImportOptions)
}

func NewImportOptions() ImportOptions { _ = "STUB: not implemented"; return *new(ImportOptions) }

type ExportOptions struct {
	StripEndingLineBreak bool
	Format               Format
	Encoding             text.Encoding
	Delimiter            rune
	DelimiterPositions   []int
	SingleLine           bool
	WithoutHeader        bool
	LineBreak            text.LineBreak
	EncloseAll           bool
	JsonEscape           txjson.EscapeType
	PrettyPrint          bool
	ScientificNotation   bool

	// For Calculation of String Width
	EastAsianEncoding    bool
	CountDiacriticalSign bool
	CountFormatCode      bool

	Color bool
}

func (ops ExportOptions) Copy() ExportOptions {
	_ = "STUB: not implemented"
	return *new(ExportOptions)
}

func NewExportOptions() ExportOptions { _ = "STUB: not implemented"; return *new(ExportOptions) }

type Flags struct {
	// Common Settings
	Repository     string
	Location       string
	DatetimeFormat []string
	AnsiQuotes     bool
	StrictEqual    bool

	WaitTimeout float64

	// For Import
	ImportOptions ImportOptions

	// For Export
	ExportOptions ExportOptions

	// System Use
	Quiet          bool
	LimitRecursion int64
	CPU            int
	Stats          bool

	defaultTimeLocation *time.Location
}

func GetDefaultNumberOfCPU() int { _ = "STUB: not implemented"; return 0 }

func NewFlags(env *Environment) (*Flags, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Flags) GetTimeLocation() *time.Location { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetRepository(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetLocation(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetDatetimeFormat(s string) { _ = "STUB: not implemented"; return }

func (f *Flags) SetAnsiQuotes(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetStrictEqual(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetWaitTimeout(t float64) { _ = "STUB: not implemented"; return }

func (f *Flags) SetImportFormat(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetDelimiter(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetAllowUnevenFields(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetDelimiterPositions(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetJsonQuery(s string) { _ = "STUB: not implemented"; return }

func (f *Flags) SetEncoding(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetNoHeader(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetWithoutNull(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetFormat(s string, outfile string, canOutputToPipe bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Flags) SetWriteEncoding(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetWriteDelimiter(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetWriteDelimiterPositions(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetWithoutHeader(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetLineBreak(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetJsonEscape(s string) error { _ = "STUB: not implemented"; return nil }

func (f *Flags) SetPrettyPrint(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetScientificNotation(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetStripEndingLineBreak(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetEncloseAll(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetColor(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetEastAsianEncoding(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetCountDiacriticalSign(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetCountFormatCode(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetQuiet(b bool) { _ = "STUB: not implemented"; return }

func (f *Flags) SetLimitRecursion(i int64) { _ = "STUB: not implemented"; return }

func (f *Flags) SetCPU(i int) { _ = "STUB: not implemented"; return }

func (f *Flags) SetStats(b bool) { _ = "STUB: not implemented"; return }
