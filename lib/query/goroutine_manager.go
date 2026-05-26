package query

import (
	"context"
	"sync"
)

var (
	gm    *GoroutineManager
	getGm sync.Once
)

const MinimumRequiredPerCPUCore = 80

func GetGoroutineManager() *GoroutineManager { _ = "STUB: not implemented"; return nil }

type GoroutineManager struct {
	Count                  int
	CountMutex             *sync.Mutex
	MinimumRequiredPerCore int
}

func (m *GoroutineManager) AssignRoutineNumber(recordLen int, minimumRequiredPerCore int, cpuNum int) int {
	_ = "STUB: not implemented"
	return 0
}

func (m *GoroutineManager) Release() { _ = "STUB: not implemented"; return }

type GoroutineTaskManager struct {
	Number int

	grTaskMutex *sync.Mutex
	grCount     int
	recordLen   int
	waitGroup   sync.WaitGroup
	err         error
}

func NewGoroutineTaskManager(recordLen int, minimumRequiredPerCore int, cpuNum int) *GoroutineTaskManager {
	_ = "STUB: not implemented"
	return nil
}

func (m *GoroutineTaskManager) HasError() bool { _ = "STUB: not implemented"; return false }

func (m *GoroutineTaskManager) SetError(e error) { _ = "STUB: not implemented"; return }

func (m *GoroutineTaskManager) Err() error { _ = "STUB: not implemented"; return nil }

func (m *GoroutineTaskManager) RecordRange(routineIndex int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (m *GoroutineTaskManager) Add() { _ = "STUB: not implemented"; return }

func (m *GoroutineTaskManager) Done() { _ = "STUB: not implemented"; return }

func (m *GoroutineTaskManager) Wait() { _ = "STUB: not implemented"; return }

func (m *GoroutineTaskManager) run(ctx context.Context, fn func(int) error, thIdx int) {
	_ = "STUB: not implemented"
	return
}

func (m *GoroutineTaskManager) Run(ctx context.Context, fn func(int) error) error {
	_ = "STUB: not implemented"
	return nil
}
