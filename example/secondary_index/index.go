package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yandex-cloud/ydb-go-sdk/v2/connect"
	"github.com/yandex-cloud/ydb-go-sdk/v2/example/internal/cli"
	"github.com/yandex-cloud/ydb-go-sdk/v2/table"
)

const (
	TableSeries         = "series"
	TableSeriesRevViews = "series_rev_views"
)

var actions = map[string]func(context.Context, *table.SessionPool, string, ...string) error{
	"create":    doCreate,
	"generate":  doGenerate,
	"update":    doUpdate,
	"list":      doList,
	"listviews": doListViews,
	"delete":    doDelete,
	"drop":      doDrop,
}

type Command struct {
}

func (cmd *Command) ExportFlags(_ context.Context, flagSet *flag.FlagSet) {
	flagSet.Usage = func() {
		out := flagSet.Output()
		_, _ = fmt.Fprintf(out, "Usage:\n%s command [options]\n", os.Args[0])
		_, _ = fmt.Fprintf(out, "\nOptions:\n")
		flagSet.PrintDefaults()
		_, _ = fmt.Fprintf(out, "\nCommands:\n")
		for c := range actions {
			_, _ = fmt.Fprintf(out, "  - %s\n", c)
		}
	}
}

func (cmd *Command) Run(ctx context.Context, params cli.Parameters) error {
	if len(params.Args) < 1 {
		fmt.Printf("no command specified.\n\n")
		return cli.ErrPrintUsage
	}
	var (
		name   = strings.ToLower(params.Args[0])
		action = actions[name]
	)
	if action == nil {
		fmt.Printf("unexpected command: %q\n\n", params.Args[0])
		return cli.ErrPrintUsage
	}

	connectCtx, cancel := context.WithTimeout(ctx, params.ConnectTimeout)
	defer cancel()
	db, err := connect.New(connectCtx, params.ConnectParams)
	if err != nil {
		return fmt.Errorf("connect error: %w", err)
	}
	defer db.Close()

	return action(ctx, db.Table().Pool(), params.Prefix(), params.Args[1:]...)
}
