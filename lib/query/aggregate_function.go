package query

import (
	"github.com/mithrandie/csvq/lib/option"

	"github.com/mithrandie/csvq/lib/value"
)

type AggregateFunction func([]value.Primary, *option.Flags) value.Primary

var AggregateFunctions = map[string]AggregateFunction{
	"COUNT":  Count,
	"MAX":    Max,
	"MIN":    Min,
	"SUM":    Sum,
	"AVG":    Avg,
	"STDEV":  StdEV,
	"STDEVP": StdEVP,
	"VAR":    Var,
	"VARP":   VarP,
	"MEDIAN": Median,
}

func Count(list []value.Primary, _ *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func Max(list []value.Primary, flags *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func Min(list []value.Primary, flags *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func Sum(list []value.Primary, _ *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func Avg(list []value.Primary, _ *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func StdEV(list []value.Primary, _ *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func StdEVP(list []value.Primary, _ *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func Var(list []value.Primary, _ *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func VarP(list []value.Primary, _ *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func floatList(list []value.Primary) []float64 { _ = "STUB: not implemented"; return nil }

func sum(list []float64) float64 { _ = "STUB: not implemented"; return 0 }

func average(list []float64) float64 { _ = "STUB: not implemented"; return 0 }

func variance(list []float64, isP bool) float64 { _ = "STUB: not implemented"; return 0 }

func standardDeviation(list []float64, isP bool) float64 { _ = "STUB: not implemented"; return 0 }

func Median(list []value.Primary, flags *option.Flags) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func ListAgg(list []value.Primary, separator string) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}

func JsonAgg(list []value.Primary) value.Primary {
	_ = "STUB: not implemented"
	return *new(value.Primary)
}
