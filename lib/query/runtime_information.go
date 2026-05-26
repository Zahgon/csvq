package query

import (
	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

const (
	UncommittedInformation  = "UNCOMMITTED"
	CreatedInformation      = "CREATED"
	UpdatedInformation      = "UPDATED"
	UpdatedViewsInformation = "UPDATED_VIEWS"
	LoadedTablesInformation = "LOADED_TABLES"
	WorkingDirectory        = "WORKING_DIRECTORY"
	VersionInformation      = "VERSION"
)

var RuntimeInformatinList = []string{
	UncommittedInformation,
	CreatedInformation,
	UpdatedInformation,
	UpdatedViewsInformation,
	LoadedTablesInformation,
	WorkingDirectory,
	VersionInformation,
}

func GetRuntimeInformation(tx *Transaction, expr parser.RuntimeInformation) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}
