package file

import (
	"context"
	"time"
)

type Container struct {
	m map[string]*Handler
}

func NewContainer() *Container { _ = "STUB: not implemented"; return nil }

func (c *Container) Keys() []string { _ = "STUB: not implemented"; return nil }

func (c *Container) Add(path string, handler *Handler) error { _ = "STUB: not implemented"; return nil }

func (c *Container) Remove(path string) { _ = "STUB: not implemented"; return }

func (c *Container) CreateHandlerWithoutLock(ctx context.Context, path string, defaultWaitTimeout time.Duration, retryDelay time.Duration) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Container) CreateHandlerForRead(ctx context.Context, path string, defaultWaitTimeout time.Duration, retryDelay time.Duration) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Container) CreateHandlerForCreate(path string) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Container) CreateHandlerForUpdate(ctx context.Context, path string, defaultWaitTimeout time.Duration, retryDelay time.Duration) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Container) createHandler(ctx context.Context, path string, defaultWaitTimeout time.Duration, retryDelay time.Duration, fn func(context.Context, string, time.Duration, time.Duration) (*Handler, error)) (*Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Container) Close(h *Handler) error { _ = "STUB: not implemented"; return nil }

func (c *Container) Commit(h *Handler) error { _ = "STUB: not implemented"; return nil }

func (c *Container) CloseWithErrors(h *Handler) (err error) { _ = "STUB: not implemented"; return nil }

func (c *Container) CloseAll() error { _ = "STUB: not implemented"; return nil }

func (c *Container) CloseAllWithErrors() error { _ = "STUB: not implemented"; return nil }
