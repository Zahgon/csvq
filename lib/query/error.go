package query

import (
	"os"

	"github.com/mithrandie/csvq/lib/parser"
)

const ExitMessage = "exit"
const DefaultUserTriggeredErrorMessage = "triggered error"

const (
	ErrorMessageTemplate                 = "[L:%d C:%d] %s"
	ErrorMessageWithFilepathTemplate     = "%s [L:%d C:%d] %s"
	ErrorMessageWithCustomPrefixTemplate = "[%s] %s"

	ErrMsgSignalReceived = "signal received: %s"

	ErrMsgIncorrectCommandUsage                = "incorrect usage: %s"
	ErrMsgInvalidValueExpression               = "%s: cannot evaluate as a value"
	ErrMsgInvalidPath                          = "%s: %s"
	ErrMsgIO                                   = "%s"
	ErrMsgCommit                               = "failed to commit: %s"
	ErrMsgRollback                             = "failed to rollback: %s"
	ErrMsgCannotDetectFileEncoding             = "cannot detect character encoding: %s"
	ErrMsgFieldAmbiguous                       = "field %s is ambiguous"
	ErrMsgFieldNotExist                        = "field %s does not exist"
	ErrMsgFieldNotGroupKey                     = "field %s is not a group key"
	ErrMsgDuplicateFieldName                   = "field name %s is a duplicate"
	ErrMsgNotGroupingRecords                   = "function %s cannot aggregate not grouping records"
	ErrMsgNotAllowedAnalyticFunction           = "analytic function %s is only available in select clause or order by clause"
	ErrMsgUndeclaredVariable                   = "variable %s is undeclared"
	ErrMsgVariableRedeclared                   = "variable %s is redeclared"
	ErrMsgUndefinedConstant                    = "constant %s is not defined"
	ErrMsgInvalidUrl                           = "failed to parse %q as url"
	ErrMsgUnsupportedUrlScheme                 = "url scheme %s is not supported"
	ErrMsgFunctionNotExist                     = "function %s does not exist"
	ErrMsgFunctionArgumentsLength              = "function %s takes %s"
	ErrMsgFunctionInvalidArgument              = "%s for function %s"
	ErrMsgNestedAggregateFunctions             = "aggregate functions are nested at %s"
	ErrMsgFunctionRedeclared                   = "function %s is redeclared"
	ErrMsgBuiltInFunctionDeclared              = "function %s is a built-in function"
	ErrMsgDuplicateParameter                   = "parameter %s is a duplicate"
	ErrMsgSubqueryTooManyRecords               = "subquery returns too many records, should return only one record"
	ErrMsgSubqueryTooManyFields                = "subquery returns too many fields, should return only one field"
	ErrMsgJsonQueryTooManyRecords              = "json query returns too many records, should return only one record"
	ErrMsgLoadJson                             = "json loading error: %s"
	ErrMsgJsonLinesStructure                   = "json lines must be an array of objects"
	ErrMsgIncorrectLateralUsage                = "LATERAL cannot to be used in a RIGHT or FULL outer join"
	ErrMsgEmptyInlineTable                     = "inline table is empty"
	ErrMsgInvalidTableObject                   = "invalid table object: %s"
	ErrMsgTableObjectInvalidDelimiter          = "invalid delimiter: %s"
	ErrMsgTableObjectInvalidDelimiterPositions = "invalid delimiter positions: %s"
	ErrMsgTableObjectInvalidJsonQuery          = "invalid json query: %s"
	ErrMsgTableObjectArgumentsLength           = "table object %s takes at most %d arguments"
	ErrMsgTableObjectJsonArgumentsLength       = "table object %s takes exactly %d arguments"
	ErrMsgTableObjectInvalidArgument           = "invalid argument for %s: %s"
	ErrMsgCursorRedeclared                     = "cursor %s is redeclared"
	ErrMsgUndeclaredCursor                     = "cursor %s is undeclared"
	ErrMsgCursorClosed                         = "cursor %s is closed"
	ErrMsgCursorOpen                           = "cursor %s is already open"
	ErrMsgInvalidCursorStatement               = "invalid cursor statement: %s"
	ErrMsgPseudoCursor                         = "cursor %s is a pseudo cursor"
	ErrMsgCursorFetchLength                    = "fetching from cursor %s returns %s"
	ErrMsgInvalidFetchPosition                 = "fetching position %s is not an integer value"
	ErrMsgInlineTableRedefined                 = "inline table %s is redefined"
	ErrMsgUndefinedInlineTable                 = "inline table %s is undefined"
	ErrMsgInlineTableFieldLength               = "select query should return exactly %s for inline table %s"
	ErrMsgFileNotExist                         = "file %s does not exist"
	ErrMsgFileAlreadyExist                     = "file %s already exists"
	ErrMsgFileUnableToRead                     = "file %s is unable to be read"
	ErrMsgFileLockTimeout                      = "file %s: lock wait timeout period exceeded"
	ErrMsgFileNameAmbiguous                    = "filename %s is ambiguous"
	ErrMsgDataParsing                          = "data parse error in %s: %s"
	ErrMsgDataEncoding                         = "data encode error: %s"
	ErrMsgTableFieldLength                     = "select query should return exactly %s for table %s"
	ErrMsgTemporaryTableRedeclared             = "view %s is redeclared"
	ErrMsgUndeclaredTemporaryTable             = "view %s is undeclared"
	ErrMsgTemporaryTableFieldLength            = "select query should return exactly %s for view %s"
	ErrMsgDuplicateTableName                   = "table name %s is a duplicate"
	ErrMsgTableNotLoaded                       = "table %s is not loaded"
	ErrMsgStdinEmpty                           = "STDIN is empty"
	ErrMsgInlineTableCannotBeUpdated           = "inline table cannot be updated"
	ErrMsgAliasMustBeSpecifiedForUpdate        = "alias to table identification function or URL must be specified for update"
	ErrMsgRowValueLengthInComparison           = "row value should contain exactly %s"
	ErrMsgFieldLengthInComparison              = "select query should return exactly %s"
	ErrMsgInvalidLimitPercentage               = "limit percentage %s is not a float value"
	ErrMsgInvalidLimitNumber                   = "limit number of records %s is not an integer value"
	ErrMsgInvalidOffsetNumber                  = "offset number %s is not an integer value"
	ErrMsgCombinedSetFieldLength               = "result set to be combined should contain exactly %s"
	ErrMsgRecursionExceededLimit               = "iteration of recursive query exceeded the limit %d"
	ErrMsgNestedRecursion                      = "recursive queries are nested"
	ErrMsgInsertRowValueLength                 = "row value should contain exactly %s"
	ErrMsgInsertSelectFieldLength              = "select query should return exactly %s"
	ErrMsgUpdateFieldNotExist                  = "field %s does not exist in the tables to update"
	ErrMsgUpdateValueAmbiguous                 = "value %s to set in the field %s is ambiguous"
	ErrMsgReplaceKeyNotSet                     = "replace Key %s is not set"
	ErrMsgDeleteTableNotSpecified              = "tables to delete records are not specified"
	ErrMsgShowInvalidObjectType                = "object type %s is invalid"
	ErrMsgReplaceValueLength                   = "%s"
	ErrMsgSourceInvalidFilePath                = "%s is a invalid file path"
	ErrMsgInvalidFlagName                      = "%s is an unknown flag"
	ErrMsgFlagValueNowAllowedFormat            = "%s for %s is not allowed"
	ErrMsgInvalidFlagValue                     = "%s"
	ErrMsgAddFlagNotSupportedName              = "add flag element syntax does not support %s"
	ErrMsgRemoveFlagNotSupportedName           = "remove flag element syntax does not support %s"
	ErrMsgInvalidFlagValueToBeRemoved          = "%s is an invalid value for %s to specify the element"
	ErrMsgInvalidRuntimeInformation            = "%s is an unknown runtime information"
	ErrMsgNotTable                             = "table attributes can only be set on files"
	ErrMsgInvalidTableAttributeName            = "table attribute %s does not exist"
	ErrMsgTableAttributeValueNotAllowedFormat  = "%s for %s is not allowed"
	ErrMsgInvalidTableAttributeValue           = "%s"
	ErrMsgInvalidEventName                     = "%s is an unknown event"
	ErrMsgInternalRecordIdNotExist             = "internal record id does not exist"
	ErrMsgInternalRecordIdEmpty                = "internal record id is empty"
	ErrMsgFieldLengthNotMatch                  = "field length does not match"
	ErrMsgRowValueLengthInList                 = "row value length does not match at index %d"
	ErrMsgFormatStringLengthNotMatch           = "number of replace values does not match"
	ErrMsgUnknownFormatPlaceholder             = "%q is an unknown placeholder"
	ErrMsgFormatUnexpectedTermination          = "unexpected termination of format string"
	ErrMsgExternalCommand                      = "external command: %s"
	ErrMsgHttpRequest                          = "failed to get resource from %s: %s"
	ErrMsgInvalidReloadType                    = "%s is an unknown reload type"
	ErrMsgLoadConfiguration                    = "configuration loading error: %s"
	ErrMsgDuplicateStatementName               = "statement %s is a duplicate"
	ErrMsgStatementNotExist                    = "statement %s does not exist"
	ErrMsgStatementReplaceValueNotSpecified    = "replace value for %s is not specified"
	ErrMsgSelectIntoQueryFieldLengthNotMatch   = "select into query should return exactly %s"
	ErrMsgSelectIntoQueryTooManyRecords        = "select into query returns too many records, should return only one record"
	ErrMsgIntegerDevidedByZero                 = "integer divided by zero"
)

type Error interface {
	Error() string
	Message() string
	Code() int
	Number() int
	Line() int
	Char() int
	Source() string
	appendCompositeError(Error)
}

type BaseError struct {
	source        string
	line          int
	char          int
	message       string
	code          int
	number        int
	prefix        string
	compositeErrs []Error
}

func (e *BaseError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *BaseError) err() string { _ = "STUB: not implemented"; return "" }

func (e *BaseError) Message() string { _ = "STUB: not implemented"; return "" }

func (e *BaseError) Code() int { _ = "STUB: not implemented"; return 0 }

func (e *BaseError) Number() int { _ = "STUB: not implemented"; return 0 }

func (e *BaseError) Line() int { _ = "STUB: not implemented"; return 0 }

func (e *BaseError) Char() int { _ = "STUB: not implemented"; return 0 }

func (e *BaseError) Source() string { _ = "STUB: not implemented"; return "" }

func (e *BaseError) appendCompositeError(err Error) { _ = "STUB: not implemented"; return }

func appendCompositeError(e1 error, e2 error) error { _ = "STUB: not implemented"; return nil }

func NewBaseError(expr parser.Expression, message string, code int, number int) *BaseError {
	_ = "STUB: not implemented"
	return nil
}

func NewBaseErrorWithPrefix(prefix string, message string, code int, number int) *BaseError {
	_ = "STUB: not implemented"
	return nil
}

type FatalError struct {
	*BaseError
}

func NewFatalError(panicReport interface{}) error { _ = "STUB: not implemented"; return nil }

type SystemError struct {
	*BaseError
}

func NewSystemError(message string) error { _ = "STUB: not implemented"; return nil }

type ForcedExit struct {
	*BaseError
}

func NewForcedExit(code int) error { _ = "STUB: not implemented"; return nil }

type UserTriggeredError struct {
	*BaseError
}

func NewUserTriggeredError(expr parser.Trigger, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type SignalReceived struct {
	*BaseError
}

func NewSignalReceived(sig os.Signal) error { _ = "STUB: not implemented"; return nil }

type SyntaxError struct {
	*BaseError
}

func NewSyntaxError(err *parser.SyntaxError) error { _ = "STUB: not implemented"; return nil }

type PreparedStatementSyntaxError struct {
	*BaseError
}

func NewPreparedStatementSyntaxError(err *parser.SyntaxError) error {
	_ = "STUB: not implemented"
	return nil
}

type ContextCanceled struct {
	*BaseError
}

func NewContextCanceled(message string) error { _ = "STUB: not implemented"; return nil }

type ContextDone struct {
	*BaseError
}

func NewContextDone(message string) error { _ = "STUB: not implemented"; return nil }

type IncorrectCommandUsageError struct {
	*BaseError
}

func NewIncorrectCommandUsageError(message string) error { _ = "STUB: not implemented"; return nil }

type InvalidValueExpressionError struct {
	*BaseError
}

func NewInvalidValueExpressionError(expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidPathError struct {
	*BaseError
}

func NewInvalidPathError(expr parser.Expression, path string, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type IOError struct {
	*BaseError
}

func NewIOError(expr parser.QueryExpression, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type CommitError struct {
	*BaseError
}

func NewCommitError(expr parser.Expression, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type RollbackError struct {
	*BaseError
}

func NewRollbackError(expr parser.Expression, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type CannotDetectFileEncodingError struct {
	*BaseError
}

func NewCannotDetectFileEncodingError(file parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type FieldAmbiguousError struct {
	*BaseError
}

func NewFieldAmbiguousError(field parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type FieldNotExistError struct {
	*BaseError
}

func NewFieldNotExistError(field parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type FieldNotGroupKeyError struct {
	*BaseError
}

func NewFieldNotGroupKeyError(field parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type DuplicateFieldNameError struct {
	*BaseError
}

func NewDuplicateFieldNameError(fieldName parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type NotGroupingRecordsError struct {
	*BaseError
}

func NewNotGroupingRecordsError(expr parser.QueryExpression, funcname string) error {
	_ = "STUB: not implemented"
	return nil
}

type NotAllowedAnalyticFunctionError struct {
	*BaseError
}

func NewNotAllowedAnalyticFunctionError(expr parser.AnalyticFunction) error {
	_ = "STUB: not implemented"
	return nil
}

type UndeclaredVariableError struct {
	*BaseError
}

func NewUndeclaredVariableError(expr parser.Variable) error { _ = "STUB: not implemented"; return nil }

type VariableRedeclaredError struct {
	*BaseError
}

func NewVariableRedeclaredError(expr parser.Variable) error { _ = "STUB: not implemented"; return nil }

type UndefinedConstantError struct {
	*BaseError
}

func NewUndefinedConstantError(expr parser.Constant) error { _ = "STUB: not implemented"; return nil }

type InvalidUrlError struct {
	*BaseError
}

func NewInvalidUrlError(expr parser.Url) error { _ = "STUB: not implemented"; return nil }

type UnsupportedUrlSchemeError struct {
	*BaseError
}

func NewUnsupportedUrlSchemeError(expr parser.Url, scheme string) error {
	_ = "STUB: not implemented"
	return nil
}

type FunctionNotExistError struct {
	*BaseError
}

func NewFunctionNotExistError(expr parser.QueryExpression, funcname string) error {
	_ = "STUB: not implemented"
	return nil
}

type FunctionArgumentLengthError struct {
	*BaseError
}

func NewFunctionArgumentLengthError(expr parser.QueryExpression, funcname string, argslen []int) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFunctionArgumentLengthErrorWithCustomArgs(expr parser.QueryExpression, funcname string, argstr string) error {
	_ = "STUB: not implemented"
	return nil
}

type FunctionInvalidArgumentError struct {
	*BaseError
}

func NewFunctionInvalidArgumentError(function parser.QueryExpression, funcname string, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type NestedAggregateFunctionsError struct {
	*BaseError
}

func NewNestedAggregateFunctionsError(expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type FunctionRedeclaredError struct {
	*BaseError
}

func NewFunctionRedeclaredError(expr parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type BuiltInFunctionDeclaredError struct {
	*BaseError
}

func NewBuiltInFunctionDeclaredError(expr parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type DuplicateParameterError struct {
	*BaseError
}

func NewDuplicateParameterError(expr parser.Variable) error { _ = "STUB: not implemented"; return nil }

type SubqueryTooManyRecordsError struct {
	*BaseError
}

func NewSubqueryTooManyRecordsError(expr parser.Subquery) error {
	_ = "STUB: not implemented"
	return nil
}

type SubqueryTooManyFieldsError struct {
	*BaseError
}

func NewSubqueryTooManyFieldsError(expr parser.Subquery) error {
	_ = "STUB: not implemented"
	return nil
}

type JsonQueryTooManyRecordsError struct {
	*BaseError
}

func NewJsonQueryTooManyRecordsError(expr parser.JsonQuery) error {
	_ = "STUB: not implemented"
	return nil
}

type LoadJsonError struct {
	*BaseError
}

func NewLoadJsonError(expr parser.QueryExpression, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type JsonLinesStructureError struct {
	*BaseError
}

func NewJsonLinesStructureError(expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type IncorrectLateralUsageError struct {
	*BaseError
}

func NewIncorrectLateralUsageError(expr parser.Table) error { _ = "STUB: not implemented"; return nil }

type EmptyInlineTableError struct {
	*BaseError
}

func NewEmptyInlineTableError(expr parser.FormatSpecifiedFunction) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidTableObjectError struct {
	*BaseError
}

func NewInvalidTableObjectError(expr parser.FormatSpecifiedFunction, objectName string) error {
	_ = "STUB: not implemented"
	return nil
}

type TableObjectInvalidDelimiterError struct {
	*BaseError
}

func NewTableObjectInvalidDelimiterError(expr parser.FormatSpecifiedFunction, delimiter string) error {
	_ = "STUB: not implemented"
	return nil
}

type TableObjectInvalidDelimiterPositionsError struct {
	*BaseError
}

func NewTableObjectInvalidDelimiterPositionsError(expr parser.FormatSpecifiedFunction, positions string) error {
	_ = "STUB: not implemented"
	return nil
}

type TableObjectInvalidJsonQueryError struct {
	*BaseError
}

func NewTableObjectInvalidJsonQueryError(expr parser.FormatSpecifiedFunction, jsonQuery string) error {
	_ = "STUB: not implemented"
	return nil
}

type TableObjectArgumentsLengthError struct {
	*BaseError
}

func NewTableObjectArgumentsLengthError(expr parser.FormatSpecifiedFunction, argLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type TableObjectJsonArgumentsLengthError struct {
	*BaseError
}

func NewTableObjectJsonArgumentsLengthError(expr parser.FormatSpecifiedFunction, argLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type TableObjectInvalidArgumentError struct {
	*BaseError
}

func NewTableObjectInvalidArgumentError(expr parser.FormatSpecifiedFunction, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type CursorRedeclaredError struct {
	*BaseError
}

func NewCursorRedeclaredError(cursor parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type UndeclaredCursorError struct {
	*BaseError
}

func NewUndeclaredCursorError(cursor parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type CursorClosedError struct {
	*BaseError
}

func NewCursorClosedError(cursor parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type CursorOpenError struct {
	*BaseError
}

func NewCursorOpenError(cursor parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type InvalidCursorStatementError struct {
	*BaseError
}

func NewInvalidCursorStatementError(statement parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type PseudoCursorError struct {
	*BaseError
}

func NewPseudoCursorError(cursor parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type CursorFetchLengthError struct {
	*BaseError
}

func NewCursorFetchLengthError(cursor parser.Identifier, returnLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidFetchPositionError struct {
	*BaseError
}

func NewInvalidFetchPositionError(position parser.FetchPosition) error {
	_ = "STUB: not implemented"
	return nil
}

type InLineTableRedefinedError struct {
	*BaseError
}

func NewInLineTableRedefinedError(table parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type UndefinedInLineTableError struct {
	*BaseError
}

func NewUndefinedInLineTableError(table parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type InlineTableFieldLengthError struct {
	*BaseError
}

func NewInlineTableFieldLengthError(query parser.SelectQuery, table parser.Identifier, fieldLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type FileNotExistError struct {
	*BaseError
}

func NewFileNotExistError(file parser.QueryExpression) error { _ = "STUB: not implemented"; return nil }

type FileAlreadyExistError struct {
	*BaseError
}

func NewFileAlreadyExistError(file parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type FileUnableToReadError struct {
	*BaseError
}

func NewFileUnableToReadError(file parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type FileLockTimeoutError struct {
	*BaseError
}

func NewFileLockTimeoutError(file parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type FileNameAmbiguousError struct {
	*BaseError
}

func NewFileNameAmbiguousError(file parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type DataParsingError struct {
	*BaseError
}

func NewDataParsingError(file parser.QueryExpression, filepath string, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type DataEncodingError struct {
	*BaseError
}

func NewDataEncodingError(message string) error { _ = "STUB: not implemented"; return nil }

type TableFieldLengthError struct {
	*BaseError
}

func NewTableFieldLengthError(query parser.SelectQuery, table parser.Identifier, fieldLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type TemporaryTableRedeclaredError struct {
	*BaseError
}

func NewTemporaryTableRedeclaredError(table parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type UndeclaredTemporaryTableError struct {
	*BaseError
}

func NewUndeclaredTemporaryTableError(table parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type TemporaryTableFieldLengthError struct {
	*BaseError
}

func NewTemporaryTableFieldLengthError(query parser.SelectQuery, table parser.Identifier, fieldLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type DuplicateTableNameError struct {
	*BaseError
}

func NewDuplicateTableNameError(table parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type TableNotLoadedError struct {
	*BaseError
}

func NewTableNotLoadedError(table parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type StdinEmptyError struct {
	*BaseError
}

func NewStdinEmptyError(stdin parser.Stdin) error { _ = "STUB: not implemented"; return nil }

type InlineTableCannotBeUpdatedError struct {
	*BaseError
}

func NewInlineTableCannotBeUpdatedError(expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type AliasMustBeSpecifiedForUpdateError struct {
	*BaseError
}

func NewAliasMustBeSpecifiedForUpdateError(expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type RowValueLengthInComparisonError struct {
	*BaseError
}

func NewRowValueLengthInComparisonError(expr parser.QueryExpression, valueLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type SelectFieldLengthInComparisonError struct {
	*BaseError
}

func NewSelectFieldLengthInComparisonError(query parser.Subquery, valueLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidLimitPercentageError struct {
	*BaseError
}

func NewInvalidLimitPercentageError(clause parser.LimitClause) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidLimitNumberError struct {
	*BaseError
}

func NewInvalidLimitNumberError(clause parser.LimitClause) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidOffsetNumberError struct {
	*BaseError
}

func NewInvalidOffsetNumberError(clause parser.OffsetClause) error {
	_ = "STUB: not implemented"
	return nil
}

type CombinedSetFieldLengthError struct {
	*BaseError
}

func NewCombinedSetFieldLengthError(selectEntity parser.QueryExpression, fieldLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type RecursionExceededLimitError struct {
	*BaseError
}

func NewRecursionExceededLimitError(selectEntity parser.QueryExpression, limit int64) error {
	_ = "STUB: not implemented"
	return nil
}

type NestedRecursionError struct {
	*BaseError
}

func NewNestedRecursionError(expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type InsertRowValueLengthError struct {
	*BaseError
}

func NewInsertRowValueLengthError(rowValue parser.RowValue, valueLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type InsertSelectFieldLengthError struct {
	*BaseError
}

func NewInsertSelectFieldLengthError(query parser.SelectQuery, fieldLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type UpdateFieldNotExistError struct {
	*BaseError
}

func NewUpdateFieldNotExistError(field parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type UpdateValueAmbiguousError struct {
	*BaseError
}

func NewUpdateValueAmbiguousError(field parser.QueryExpression, value parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type ReplaceKeyNotSetError struct {
	*BaseError
}

func NewReplaceKeyNotSetError(key parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type DeleteTableNotSpecifiedError struct {
	*BaseError
}

func NewDeleteTableNotSpecifiedError(query parser.DeleteQuery) error {
	_ = "STUB: not implemented"
	return nil
}

type ShowInvalidObjectTypeError struct {
	*BaseError
}

func NewShowInvalidObjectTypeError(expr parser.Expression, objectType string) error {
	_ = "STUB: not implemented"
	return nil
}

type ReplaceValueLengthError struct {
	*BaseError
}

func NewReplaceValueLengthError(expr parser.Expression, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type SourceInvalidFilePathError struct {
	*BaseError
}

func NewSourceInvalidFilePathError(source parser.Source, arg parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidFlagNameError struct {
	*BaseError
}

func NewInvalidFlagNameError(expr parser.Flag) error { _ = "STUB: not implemented"; return nil }

type InvalidRuntimeInformationError struct {
	*BaseError
}

func NewInvalidRuntimeInformationError(expr parser.RuntimeInformation) error {
	_ = "STUB: not implemented"
	return nil
}

type FlagValueNotAllowedFormatError struct {
	*BaseError
}

func NewFlagValueNotAllowedFormatError(setFlag parser.SetFlag) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidFlagValueError struct {
	*BaseError
}

func NewInvalidFlagValueError(expr parser.SetFlag, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type AddFlagNotSupportedNameError struct {
	*BaseError
}

func NewAddFlagNotSupportedNameError(expr parser.AddFlagElement) error {
	_ = "STUB: not implemented"
	return nil
}

type RemoveFlagNotSupportedNameError struct {
	*BaseError
}

func NewRemoveFlagNotSupportedNameError(expr parser.RemoveFlagElement) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidFlagValueToBeRemoveError struct {
	*BaseError
}

func NewInvalidFlagValueToBeRemovedError(unsetFlag parser.RemoveFlagElement) error {
	_ = "STUB: not implemented"
	return nil
}

type NotTableError struct {
	*BaseError
}

func NewNotTableError(expr parser.QueryExpression) error { _ = "STUB: not implemented"; return nil }

type InvalidTableAttributeNameError struct {
	*BaseError
}

func NewInvalidTableAttributeNameError(expr parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type TableAttributeValueNotAllowedFormatError struct {
	*BaseError
}

func NewTableAttributeValueNotAllowedFormatError(expr parser.SetTableAttribute) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidTableAttributeValueError struct {
	*BaseError
}

func NewInvalidTableAttributeValueError(expr parser.SetTableAttribute, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidEventNameError struct {
	*BaseError
}

func NewInvalidEventNameError(expr parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type InternalRecordIdNotExistError struct {
	*BaseError
}

func NewInternalRecordIdNotExistError() error { _ = "STUB: not implemented"; return nil }

type InternalRecordIdEmptyError struct {
	*BaseError
}

func NewInternalRecordIdEmptyError() error { _ = "STUB: not implemented"; return nil }

type FieldLengthNotMatchError struct {
	*BaseError
}

func NewFieldLengthNotMatchError(expr parser.QueryExpression) error {
	_ = "STUB: not implemented"
	return nil
}

type RowValueLengthInListError struct {
	*BaseError
	Index int
}

func NewRowValueLengthInListError(i int) error { _ = "STUB: not implemented"; return nil }

type FormatStringLengthNotMatchError struct {
	*BaseError
}

func NewFormatStringLengthNotMatchError() error { _ = "STUB: not implemented"; return nil }

type UnknownFormatPlaceholderError struct {
	*BaseError
}

func NewUnknownFormatPlaceholderError(placeholder rune) error {
	_ = "STUB: not implemented"
	return nil
}

type FormatUnexpectedTerminationError struct {
	*BaseError
}

func NewFormatUnexpectedTerminationError() error { _ = "STUB: not implemented"; return nil }

type ExternalCommandError struct {
	*BaseError
}

func NewExternalCommandError(expr parser.Expression, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type HttpRequestError struct {
	*BaseError
}

func NewHttpRequestError(expr parser.Expression, url string, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type InvalidReloadTypeError struct {
	*BaseError
}

func NewInvalidReloadTypeError(expr parser.Reload, name string) error {
	_ = "STUB: not implemented"
	return nil
}

type LoadConfigurationError struct {
	*BaseError
}

func NewLoadConfigurationError(expr parser.Expression, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type DuplicateStatementNameError struct {
	*BaseError
}

func NewDuplicateStatementNameError(name parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

type StatementNotExistError struct {
	*BaseError
}

func NewStatementNotExistError(name parser.Identifier) error { _ = "STUB: not implemented"; return nil }

type StatementReplaceValueNotSpecifiedError struct {
	*BaseError
}

func NewStatementReplaceValueNotSpecifiedError(placeholder parser.Placeholder) error {
	_ = "STUB: not implemented"
	return nil
}

type SelectIntoQueryFieldLengthNotMatchError struct {
	*BaseError
}

func NewSelectIntoQueryFieldLengthNotMatchError(query parser.SelectQuery, fieldLen int) error {
	_ = "STUB: not implemented"
	return nil
}

type SelectIntoQueryTooManyRecordsError struct {
	*BaseError
}

func NewSelectIntoQueryTooManyRecordsError(query parser.SelectQuery) error {
	_ = "STUB: not implemented"
	return nil
}

type IntegerDevidedByZeroError struct {
	*BaseError
}

func NewIntegerDevidedByZeroError(expr parser.Arithmetic) error {
	_ = "STUB: not implemented"
	return nil
}

func searchSelectClause(query parser.SelectQuery) parser.SelectClause {
	_ = "STUB: not implemented"
	return *new(parser.SelectClause)
}

func searchSelectClauseInSelectEntity(selectEntity parser.QueryExpression) parser.SelectClause {
	_ = "STUB: not implemented"
	return *new(parser.SelectClause)
}

func searchSelectClauseInSelectSetEntity(selectSetEntity parser.QueryExpression) parser.SelectClause {
	_ = "STUB: not implemented"
	return *new(parser.SelectClause)
}

func ConvertFileHandlerError(err error, ident parser.Identifier) error {
	_ = "STUB: not implemented"
	return nil
}

func ConvertLoadConfigurationError(err error) error { _ = "STUB: not implemented"; return nil }

func ConvertContextError(err error) error { _ = "STUB: not implemented"; return nil }
