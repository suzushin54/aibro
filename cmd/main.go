package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/suzushin54/aibro/internal/adapter"
	"github.com/suzushin54/aibro/internal/infra/ai"
	server "github.com/suzushin54/aibro/internal/infra/grpc"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		logger.Error("PROJECT_ID is not set")
		os.Exit(1)
	}

	region := os.Getenv("REGION")
	if region == "" {
		region = "asia-northeast1"
	}

	// cf. https://cloud.google.com/vertex-ai/generative-ai/docs/learn/model-versions
	model := os.Getenv("MODEL_NAME")
	if model == "" {
		model = "gemini-1.5-pro" // auto-updated alias
	}

	config := ai.Config{
		ProjectID: projectID,
		Region:    region,
		ModelName: model,
	}

	ac, err := ai.NewClient(ctx, &config)
	if err != nil {
		logger.Error("failed to create ai client", "error", err)
		os.Exit(1)
	}

	ash := adapter.NewAibroServiceHandler(logger, ac)

	srv := server.NewServer(ash)

	go func() {
		if err := srv.Start(); err != nil {
			logger.Info("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	srv.WaitForSignal()
}
