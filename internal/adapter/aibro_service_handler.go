package adapters

import (
	"context"
	"io"
	"log/slog"

	aibrov1 "github.com/suzushin54/aibro/gen/aibro/v1"
	"github.com/suzushin54/aibro/gen/aibro/v1/aibrov1connect"
	"github.com/suzushin54/aibro/pkg/ai"

	"connectrpc.com/connect"
)

// AibroServiceHandler implements the Aibrov1connect.AibroServiceHandler interface.
type AibroServiceHandler struct {
	logger *slog.Logger
	client ai.Client
	aibrov1connect.UnimplementedAIBroServiceHandler
}

func NewAibroServiceHandler(l *slog.Logger, c ai.Client) *AibroServiceHandler {
	l = l.With("component", "AibroServiceHandler")

	return &AibroServiceHandler{
		logger: l,
		client: c,
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

		res, err := s.client.Query(ctx, req.Message, "")
		if err != nil {
			s.logger.ErrorContext(ctx, "Failed to query AI", "error", err)
			return connect.NewError(connect.CodeInternal, err)
		}

		if err = stream.Send(&aibrov1.ChatStreamResponse{
			Message: res,
		}); err != nil {
			s.logger.ErrorContext(ctx, "Failed to send message", "error", err)
			return connect.NewError(connect.CodeInternal, err)
		}
	}
}
