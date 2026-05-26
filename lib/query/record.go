package query

import (
	"sync"

	"github.com/mithrandie/csvq/lib/value"
)

type Cell []value.Primary

func NewCell(val value.Primary) Cell { _ = "STUB: not implemented"; return *new(Cell) }

func NewGroupCell(values []value.Primary) Cell { _ = "STUB: not implemented"; return *new(Cell) }

type RecordSet []Record

func (r RecordSet) Copy() RecordSet { _ = "STUB: not implemented"; return *new(RecordSet) }

func (r RecordSet) Merge(r2 RecordSet) RecordSet { _ = "STUB: not implemented"; return *new(RecordSet) }

type Record []Cell

func NewRecordWithId(internalId int, values []value.Primary) Record {
	_ = "STUB: not implemented"
	return *new(Record)
}

func NewRecord(values []value.Primary) Record { _ = "STUB: not implemented"; return *new(Record) }

func NewEmptyRecord(len int) Record { _ = "STUB: not implemented"; return *new(Record) }

func (r Record) GroupLen() int { _ = "STUB: not implemented"; return 0 }

func (r Record) Copy() Record { _ = "STUB: not implemented"; return *new(Record) }

func (r Record) Merge(r2 Record, pool *sync.Pool) Record {
	_ = "STUB: not implemented"
	return *new(Record)
}

func MergeRecordSetList(list []RecordSet) RecordSet {
	_ = "STUB: not implemented"
	return *new(RecordSet)
}
