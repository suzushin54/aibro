package main

import (
	"log/slog"
	"os"

	adapters "github.com/suzushin54/aibro/internal/adapter"
	server "github.com/suzushin54/aibro/internal/infra"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ash := adapters.NewAibroServiceHandler(logger)

	srv := server.NewServer(ash)

	go func() {
		if err := srv.Start(); err != nil {
			logger.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	srv.WaitForSignal()
}
