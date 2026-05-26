package json

import (
	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/go-text/json"
)

func LoadValue(queryString string, jsontext string) (value.Primary, error) {
	_ = "STUB: not implemented"
	return *new(value.Primary), nil
}

func LoadArray(queryString string, jsontext string) ([]value.Primary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadTable(queryString string, jsontext string) ([]string, [][]value.Primary, json.EscapeType, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(json.EscapeType), nil
}

func load(queryString string, jsontext string) (json.Structure, json.EscapeType, error) {
	_ = "STUB: not implemented"
	return *new(json.Structure), *new(json.EscapeType), nil
}

func Extract(query QueryExpression, data json.Structure) (json.Structure, error) {
	_ = "STUB: not implemented"
	return *new(json.Structure), nil
}

func existsKeyInFields(key string, list []FieldExpr) bool { _ = "STUB: not implemented"; return false }
