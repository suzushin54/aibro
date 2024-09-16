package core

import (
	"context"

	"github.com/suzushin54/aibro/internal/infra/ai"
)

type FlowFunc func(ctx context.Context, message string) (string, error)

type AIBroCore struct {
	aiClient ai.Client
	flow     FlowFunc
}

func NewAIBroCore(aiClient ai.Client) *AIBroCore {
	return &AIBroCore{
		aiClient: aiClient,
		flow:     nil,
	}
}

// DefineFlow sets the flow function for message processing
func (a *AIBroCore) DefineFlow(flow FlowFunc) {
	a.flow = flow
}

func (a *AIBroCore) ProcessMessage(ctx context.Context, message string) (<-chan string, <-chan error) {
	// TODO: DefineFlow. analyzing the message, updating conversation context...

	if a.flow != nil {
		var err error
		message, err = a.flow(ctx, message)
		if err != nil {
			errChan := make(chan error, 1)
			errChan <- err
			close(errChan)
			return nil, errChan
		}
	}

	resChan, errChan := a.aiClient.Query(ctx, message, "")

	return resChan, errChan
}
