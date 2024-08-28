package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/suzushin54/aibro/gen/aibro/v1/aibrov1connect"
	adapters "github.com/suzushin54/aibro/internal/adapter"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	server *http.Server
}

func NewServer(ash *adapters.AibroServiceHandler) *Server {
	mux := http.NewServeMux()
	path, handler := aibrov1connect.NewAIBroServiceHandler(ash)
	mux.Handle(path, handler)

	return &Server{
		server: &http.Server{
			Addr:    "0.0.0.0:8080",
			Handler: h2c.NewHandler(mux, &http2.Server{}),
		},
	}
}

func (s *Server) Start() error {
	slog.Info("HTTP server starting.", "Address", s.server.Addr)

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Failed to start HTTP server.", "error", err)
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	slog.Info("Shutting down HTTP server.", "Address", s.server.Addr)
	return s.server.Shutdown(context.Background())
}

func (s *Server) WaitForSignal() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	if err := s.Stop(); err != nil {
		slog.Error("Failed to stop server gracefully.", "error", err)
	}
}
