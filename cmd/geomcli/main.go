package main

import (
	"log/slog"
	"os"
	"petgo/internal/geometryapp/cli"
)

var logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

func main() {
	app := cli.NewApp()
	if err := app.Run(os.Args[1:]); err != nil {
		logger.Error("geomcli failed", "err", err)
		os.Exit(1)
	}
}
