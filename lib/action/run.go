package action

import (
	"context"
	"time"

	"github.com/mithrandie/csvq/lib/query"
)

func Run(ctx context.Context, proc *query.Processor, input string, sourceFile string, outfile string) error {
	_ = "STUB: not implemented"
	return nil
}

func LaunchInteractiveShell(ctx context.Context, proc *query.Processor) error {
	_ = "STUB: not implemented"
	return nil
}

func showStats(ctx context.Context, proc *query.Processor, start time.Time) {
	_ = "STUB: not implemented"
	return
}
