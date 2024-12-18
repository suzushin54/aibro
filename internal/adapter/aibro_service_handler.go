package adapter

import (
	"context"
	"io"
	"log/slog"
	"strings"

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
	s.logger.InfoContext(ctx, "ChatStream called")

	for {
		s.logger.InfoContext(ctx, "Waiting for message")

		req, err := stream.Receive()

		if err != nil {
			if err == io.EOF {
				s.logger.InfoContext(ctx, "Client closed the stream")
				// クライアントがストリームを閉じた場合は正常終了扱い
				return nil
			}
			s.logger.ErrorContext(ctx, "Failed to receive message", "error", err)
			return connect.NewError(connect.CodeInternal, err)
		}

		s.logger.InfoContext(ctx, "Received message", "content", req.Message)

		s.core.DefineFlow(func(ctx context.Context, message string) (string, error) {
			return "This is a message from your buddy. Respond in a way that speaks to you. :" + message, nil
		})

		resChan, errChan := s.core.ProcessMessage(ctx, req.Message)
		var buffer string
		for res := range resChan {
			buffer += res
		}

		buffer = strings.TrimSpace(buffer) // ignore trailing newline
		if buffer != "" {
			if err := stream.Send(&aibrov1.ChatStreamResponse{
				Message: buffer,
			}); err != nil {
				s.logger.ErrorContext(ctx, "Failed to send message", "error", err)
				return connect.NewError(connect.CodeInternal, err)
			}
		}

		if err := <-errChan; err != nil {
			s.logger.ErrorContext(ctx, "Failed to query AI", "error", err)
			return connect.NewError(connect.CodeInternal, err)
		}

		if ctx.Err() != nil {
			s.logger.InfoContext(ctx, "Context done")
			return nil
		}
	}
}
