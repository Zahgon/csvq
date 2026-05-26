package file

func ParseError(err error) error { _ = "STUB: not implemented"; return nil }

type IOError struct {
	message string
}

func NewIOError(message string) error { _ = "STUB: not implemented"; return nil }

func (e IOError) Error() string { _ = "STUB: not implemented"; return "" }

type NotExistError struct {
	message string
}

func NewNotExistError(message string) error { _ = "STUB: not implemented"; return nil }

func (e NotExistError) Error() string { _ = "STUB: not implemented"; return "" }

type AlreadyExistError struct {
	message string
}

func NewAlreadyExistError(message string) error { _ = "STUB: not implemented"; return nil }

func (e AlreadyExistError) Error() string { _ = "STUB: not implemented"; return "" }

type LockError struct {
	message string
}

func NewLockError(message string) error { _ = "STUB: not implemented"; return nil }

func (e LockError) Error() string { _ = "STUB: not implemented"; return "" }

type TimeoutError struct {
	message string
}

func NewTimeoutError(path string) error { _ = "STUB: not implemented"; return nil }

func (e TimeoutError) Error() string { _ = "STUB: not implemented"; return "" }

type ContextCanceled struct {
	message string
}

func NewContextCanceled() error { _ = "STUB: not implemented"; return nil }

func (e ContextCanceled) Error() string { _ = "STUB: not implemented"; return "" }

type ContextDone struct {
	message string
}

func NewContextDone(message string) error { _ = "STUB: not implemented"; return nil }

func (e ContextDone) Error() string { _ = "STUB: not implemented"; return "" }

type ForcedUnlockError struct {
	Errors []error
}

func NewForcedUnlockError(errs []error) error { _ = "STUB: not implemented"; return nil }

func (e ForcedUnlockError) Error() string { _ = "STUB: not implemented"; return "" }

type CompositeError struct {
	message string
}

func NewCompositeError(err1 error, err2 error) error { _ = "STUB: not implemented"; return nil }

func (e CompositeError) Error() string { _ = "STUB: not implemented"; return "" }
