package query

import (
	"sync"
)

type UncommittedViews struct {
	mtx     *sync.RWMutex
	Created map[string]*FileInfo
	Updated map[string]*FileInfo
}

func NewUncommittedViews() UncommittedViews {
	_ = "STUB: not implemented"
	return *new(UncommittedViews)
}

func (m *UncommittedViews) SetForCreatedView(fileInfo *FileInfo) { _ = "STUB: not implemented"; return }

func (m *UncommittedViews) SetForUpdatedView(fileInfo *FileInfo) { _ = "STUB: not implemented"; return }

func (m *UncommittedViews) Unset(fileInfo *FileInfo) { _ = "STUB: not implemented"; return }

func (m *UncommittedViews) Clean() { _ = "STUB: not implemented"; return }

func (m *UncommittedViews) UncommittedFiles() (map[string]*FileInfo, map[string]*FileInfo) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *UncommittedViews) UncommittedTempViews() map[string]*FileInfo {
	_ = "STUB: not implemented"
	return nil
}

func (m *UncommittedViews) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (m *UncommittedViews) CountCreatedTables() int { _ = "STUB: not implemented"; return 0 }

func (m *UncommittedViews) CountUpdatedTables() int { _ = "STUB: not implemented"; return 0 }

func (m *UncommittedViews) CountUpdatedViews() int { _ = "STUB: not implemented"; return 0 }
