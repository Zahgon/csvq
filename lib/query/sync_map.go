package query

import (
	"sync"
)

type SyncMap struct {
	m   *sync.Map
	mtx *sync.Mutex
}

func NewSyncMap() *SyncMap { _ = "STUB: not implemented"; return nil }

func (m SyncMap) store(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (m SyncMap) load(key string) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (m SyncMap) delete(key string) { _ = "STUB: not implemented"; return }

func (m SyncMap) exists(name string) bool { _ = "STUB: not implemented"; return false }

func (m SyncMap) lock() { _ = "STUB: not implemented"; return }

func (m SyncMap) unlock() { _ = "STUB: not implemented"; return }

func (m SyncMap) Clear() { _ = "STUB: not implemented"; return }

func (m SyncMap) Range(fn func(key, value interface{}) bool) { _ = "STUB: not implemented"; return }

func (m SyncMap) Keys() []string { _ = "STUB: not implemented"; return nil }

func (m SyncMap) Len() int { _ = "STUB: not implemented"; return 0 }

func (m SyncMap) SortedKeys() []string { _ = "STUB: not implemented"; return nil }
