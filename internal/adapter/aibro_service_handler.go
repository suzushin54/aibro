package adapter

import (
	"context"
	"io"
	"log/slog"

	"connectrpc.com/connect"
	aibrov1 "github.com/suzushin54/aibro/gen/aibro/v1"
	"github.com/suzushin54/aibro/gen/aibro/v1/aibrov1connect"
	"github.com/suzushin54/aibro/internal/core"
)

// AibroServiceHandler implements the Aibrov1connect.AibroServiceHandler interface.
type AibroServiceHandler struct {
	logger *slog.Logger
	core   *core.AIBroCore
	aibrov1connect.UnimplementedAIBroServiceHandler
}

func NewAibroServiceHandler(l *slog.Logger, c *core.AIBroCore) *AibroServiceHandler {
	l = l.With("component", "AibroServiceHandler")

	return &AibroServiceHandler{
		logger: l,
		core:   c,
	}
}

// ChatStream implements the ChatStream RPC method
func (s *AibroServiceHandler) ChatStream(
	ctx context.Context,
	stream *connect.BidiStream[aibrov1.ChatStreamRequest, aibrov1.ChatStreamResponse],
) error {
	for {
		req, err := stream.Receive()
		if err != nil {
			if err == io.EOF {
				// クライアントがストリームを閉じた場合は正常終了扱い
				return nil
			}
			s.logger.ErrorContext(ctx, "Failed to receive message", "error", err)
			return connect.NewError(connect.CodeInternal, err)
		}

		s.logger.InfoContext(ctx, "Received message", "content", req.Message)

		resChan, errChan := s.core.ProcessMessage(ctx, req.Message)

		for {
			select {
			case res, ok := <-resChan:
				if !ok {
					// チャネルが閉じられた場合は終了
					return nil
				}
				if err := stream.Send(&aibrov1.ChatStreamResponse{
					Message: res,
				}); err != nil {
					s.logger.ErrorContext(ctx, "Failed to send message", "error", err)
					return connect.NewError(connect.CodeInternal, err)
				}
			case err := <-errChan:
				if err != nil {
					s.logger.ErrorContext(ctx, "Failed to query AI", "error", err)
					return connect.NewError(connect.CodeInternal, err)
				}
			case <-ctx.Done():
				return nil
			}
		}
	}
}
