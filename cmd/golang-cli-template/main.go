package main

import (
	"context"
	"log"
	"os"

	app "github.com/nukokusa/golang-cli-template"
)

var Version string

func init() {
	app.Version = Version
}

func main() {
	ctx := context.Background()
	if err := app.Run(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
