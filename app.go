package app

import (
	"context"
	"fmt"
)

var Version string

type App struct{}

func New(ctx context.Context) (*App, error) {
	app := &App{}
	return app, nil
}

type HelloOption struct {
	Name string `help:"Name" name:"name" short:"n" default:"fuga"`
}

func (app *App) Hello(ctx context.Context, opt *HelloOption) error {
	fmt.Printf("Hello %s.\n", opt.Name)
	return nil
}
