package core

import (
	"context"

	"github.com/suzushin54/aibro/internal/infra/ai"
)

type AIBroCore struct {
	aiClient ai.Client
}

func NewAIBroCore(aiClient ai.Client) *AIBroCore {
	return &AIBroCore{
		aiClient: aiClient,
	}
}

func (a *AIBroCore) ProcessMessage(ctx context.Context, message string) (<-chan string, <-chan error) {
	// TODO: DefineFlow. analyzing the message, updating conversation context...

	resChan, errChan := a.aiClient.Query(ctx, message, "")

	return resChan, errChan
}
