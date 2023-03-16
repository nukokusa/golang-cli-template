package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/hashicorp/logutils"
)

type CLIOptions struct {
	Hello    *HelloOption     `cmd:"" help:"say hello"`
	LogLevel string           `help:"logging level: DEBUG, INFO, WARN, ERROR" enum:"DEBUG,INFO,WARN,ERROR" name:"loglevel" default:"INFO"`
	Version  kong.VersionFlag `help:"show Version" name:"version" short:"v"`
}

func Run(ctx context.Context, args []string) error {
	var opts CLIOptions
	parser, err := kong.New(&opts, kong.Vars{"version": "app " + Version})
	if err != nil {
		return err
	}
	kctx, err := parser.Parse(args)
	if err != nil {
		return err
	}

	logFilter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(opts.LogLevel),
		Writer:   os.Stderr,
	}
	log.SetOutput(logFilter)
	log.Println("[DEBUG] parsed args")

	app, err := New(ctx)
	if err != nil {
		return err
	}
	cmd := strings.Fields(kctx.Command())[0]
	return app.Dispatch(ctx, cmd, &opts)
}

func (app *App) Dispatch(ctx context.Context, command string, opts *CLIOptions) error {
	switch command {
	case "hello":
		return app.Hello(ctx, opts.Hello)
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}
