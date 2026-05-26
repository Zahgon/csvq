package json

import (
	"context"

	"github.com/mithrandie/go-text/json"

	"github.com/mithrandie/csvq/lib/value"
)

func ConvertToValue(structure json.Structure) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func ConvertToArray(array json.Array) []value.Primary { _ = "STUB: not implemented"; return nil }

func ConvertToTableValue(array json.Array) ([]string, [][]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ConvertTableValueToJsonStructure(ctx context.Context, fields []string, rows [][]value.Primary) (json.Structure, error) {
	_ = "STUB: not implemented"
	return *new(json.Structure), nil
}

func ParsePathes(fields []string) ([]PathExpression, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConvertRecordValueToJsonStructure(pathes []PathExpression, row []value.Primary) (json.Structure, error) {
	_ = "STUB: not implemented"
	return *new(json.Structure), nil
}

func addPathValueToRowStructure(parent json.Structure, path ObjectPath, val value.Primary, fieldLen int) json.Structure {
	_ = "STUB: not implemented"
	return *new(json.Structure)
}

func ParseValueToStructure(val value.Primary) json.Structure {
	_ = "STUB: not implemented"
	return *new(json.Structure)
}
