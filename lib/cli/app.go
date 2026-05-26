package cli

import (
	"context"

	"github.com/mithrandie/csvq/lib/query"

	"github.com/urfave/cli/v2"
)

func Run() { _ = "STUB: not implemented"; return }

func onUsageError(c *cli.Context, err error, isSubcommand bool) error {
	_ = "STUB: not implemented"
	return nil
}

func commandAction(fn func(ctx context.Context, c *cli.Context, proc *query.Processor) error) func(c *cli.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle signals

// Run preload commands

// Overwrite Flags with Command Options

func overwriteFlags(c *cli.Context, tx *query.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func runPreloadCommands(ctx context.Context, proc *query.Processor) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func readQuery(ctx context.Context, c *cli.Context, tx *query.Transaction) (queryString string, path string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Launch interactive shell

func Exit(err error, tx *query.Transaction) error { _ = "STUB: not implemented"; return nil }
