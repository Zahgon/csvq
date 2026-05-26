package query

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/mithrandie/csvq/lib/doc"
	"github.com/mithrandie/csvq/lib/file"
	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/go-text/color"
)

type UrlResource struct {
	MimeType string
	Data     []byte
}

func NewUrlResource(res *http.Response) (*UrlResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Transaction struct {
	Session *Session

	Environment *option.Environment
	Palette     *color.Palette
	Flags       *option.Flags

	WaitTimeout   time.Duration
	RetryDelay    time.Duration
	FileContainer *file.Container

	CachedViews      ViewMap
	UncommittedViews UncommittedViews

	UrlCache map[string]*UrlResource

	operationMutex   *sync.Mutex
	viewLoadingMutex *sync.Mutex
	stdinIsLocked    bool

	flagMutex *sync.RWMutex

	PreparedStatements PreparedStatementMap

	SelectedViews []*View
	AffectedRows  int

	AutoCommit bool
}

func NewTransaction(ctx context.Context, defaultWaitTimeout time.Duration, retryDelay time.Duration, session *Session) (*Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Transaction) UpdateWaitTimeout(waitTimeout float64, retryDelay time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (tx *Transaction) UseColor(useColor bool) { _ = "STUB: not implemented"; return }

func (tx *Transaction) Commit(ctx context.Context, scope *ReferenceScope, expr parser.Expression) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) Rollback(scope *ReferenceScope, expr parser.Expression) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) quietForTemporaryViews(expr parser.Expression) bool {
	_ = "STUB: not implemented"
	return false
}

func (tx *Transaction) ReleaseResources() error { _ = "STUB: not implemented"; return nil }

func (tx *Transaction) ReleaseResourcesWithErrors() error { _ = "STUB: not implemented"; return nil }

func (tx *Transaction) LockStdinContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) UnlockStdin() { _ = "STUB: not implemented"; return }

func (tx *Transaction) RLockStdinContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) RUnlockStdin() { _ = "STUB: not implemented"; return }

func (tx *Transaction) ClearUrlCache() { _ = "STUB: not implemented"; return }

func (tx *Transaction) CreateDocumentWriter() *doc.Writer { _ = "STUB: not implemented"; return nil }

func (tx *Transaction) Error(s string) string { _ = "STUB: not implemented"; return "" }

func (tx *Transaction) Warn(s string) string { _ = "STUB: not implemented"; return "" }

func (tx *Transaction) Notice(s string) string { _ = "STUB: not implemented"; return "" }

func (tx *Transaction) Log(log string, quiet bool) { _ = "STUB: not implemented"; return }

func (tx *Transaction) LogNotice(log string, quiet bool) { _ = "STUB: not implemented"; return }

func (tx *Transaction) LogWarn(log string, quiet bool) { _ = "STUB: not implemented"; return }

func (tx *Transaction) LogError(log string) { _ = "STUB: not implemented"; return }

var errNotAllowdFlagFormat = errors.New("not allowed flag format")
var errInvalidFlagName = errors.New("invalid flag name")

func (tx *Transaction) SetFormatFlag(value interface{}, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) SetFlag(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) setFlag(key string, value interface{}, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) GetFlag(key string) (value.Primary, bool) {
	_ = "STUB: not implemented"
	return *new(value.Primary), false
}
