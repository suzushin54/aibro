package adapters

import (
	"context"
	"io"
	"log/slog"

	aibrov1 "github.com/suzushin54/aibro/gen/aibro/v1"
	"github.com/suzushin54/aibro/gen/aibro/v1/aibrov1connect"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AibroServiceHandler implements the Aibrov1connect.AibroServiceHandler interface.
type AibroServiceHandler struct {
	logger *slog.Logger
	aibrov1connect.UnimplementedAIBroServiceHandler
}

func NewAibroServiceHandler(
	l *slog.Logger,
) *AibroServiceHandler {
	l = l.With("component", "AibroServiceHandler")

	return &AibroServiceHandler{
		logger: l,
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
				return status.Error(codes.OK, "Stream completed successfully")
			}
			s.logger.ErrorContext(ctx, "Failed to receive message", "error", err)
			return status.Error(codes.Internal, "Failed to receive message")
		}

		s.logger.InfoContext(ctx, "Received message", "content", req.Message)

		if err = stream.Send(&aibrov1.ChatStreamResponse{
			Message: "Hello World!",
		}); err != nil {
			s.logger.ErrorContext(ctx, "Failed to send message", "error", err)
			return status.Error(codes.Internal, "Failed to send message")
		}
	}
}
