package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/event/eventserver"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver/server"
)

func main() {

	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	config, err := env.LoadEnvConfig()
	if err != nil {
		log.Fatal(err)
	}

	container, err := container.New(config)
	if err != nil {
		log.Fatal(err)
	}

	errStart := container.Start(ctx)
	if errStart != nil {
		log.Fatal(err)
	}

	httpServer := server.New(container, config)
	eventServer := eventserver.NewEventServer(container, config)

	slog.InfoContext(ctx, "Starting Event Server [OutputMessages]...")
	go func(ctx context.Context) {
		for {
			eventServer.ConsumeOutput(ctx)
		}
	}(ctx)

	httpServer.Listen()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sc

	ctx, shutdown := context.WithTimeout(context.Background(), 2*time.Second)
	defer shutdown()

	slog.InfoContext(ctx, "Shutting down services...")

	return nil
}
