package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
	"github.com/mackee/template-tanukistack/client"
	"github.com/mackee/template-tanukistack/handler"
)

func main() {
	var cfg client.Config
	kong.Parse(&cfg)
	if err := run(cfg); err != nil {
		slog.Error("occur error", slog.Any("error", err))
	}
}

func run(cfg client.Config) error {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	cli, err := client.New(cfg)
	if err != nil {
		return err
	}

	handler := handler.NewHandler(cli)
	return handler.ListenAndServe(ctx, cfg.Addr)
}
