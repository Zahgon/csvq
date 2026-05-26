package query

import (
	"context"
	"errors"
	"io"

	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/value"

	"github.com/mithrandie/go-text"
	"github.com/mithrandie/go-text/color"
	txjson "github.com/mithrandie/go-text/json"
)

var EmptyResultSetError = errors.New("empty result set")
var DataEmpty = errors.New("data empty")

func EncodeView(ctx context.Context, fp io.Writer, view *View, options option.ExportOptions, palette *color.Palette) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// option.CSV

func encodeCSV(ctx context.Context, fp io.Writer, view *View, options option.ExportOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeFixedLengthFormat(ctx context.Context, fp io.Writer, view *View, options option.ExportOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeJson(ctx context.Context, fp io.Writer, view *View, options option.ExportOptions, palette *color.Palette) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeJsonLines(ctx context.Context, fp io.Writer, view *View, options option.ExportOptions, palette *color.Palette) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeText(ctx context.Context, fp io.Writer, view *View, options option.ExportOptions, palette *color.Palette) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func encodeLTSV(ctx context.Context, fp io.Writer, view *View, options option.ExportOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func jsonFloatFormat(useScientificNotation bool) txjson.FloatFormat {
	_ = "STUB: not implemented"
	return *new(txjson.FloatFormat)
}

func ConvertFieldContents(val value.Primary, forTextTable bool, useScientificNotation bool) (string, string, text.FieldAlignment) {
	_ = "STUB: not implemented"
	return "", "", *new(text.FieldAlignment)
}
