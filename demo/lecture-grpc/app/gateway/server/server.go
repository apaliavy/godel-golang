package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/gateway/config"
)

type Server struct {
	cfg        *config.AppConfig
	logger     *logrus.Logger
	httpServer *http.Server
	router     chi.Router
}

type Options func(s *Server) error

func New(opts ...Options) *Server {
	s := defaultServer()

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			s.logger.WithError(err).Fatal("failed to apply server option")
		}
	}

	return s
}

func defaultServer() *Server {
	return &Server{
		cfg:    config.Load(),
		logger: logrus.New(),
		router: chi.NewRouter(),
	}
}

func (s *Server) Run() error {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.API.Host, s.cfg.API.Port),
		Handler: s.router,
	}

	go func() {
		s.logger.Infof("serving at %s:%d", s.cfg.API.Host, s.cfg.API.Port)
		if err := s.httpServer.ListenAndServe(); err != nil {
			s.logger.WithError(err).Fatal("failed to start HTTP server")
		}
	}()

	return gracefulShutdown(context.Background(), s)
}

func gracefulShutdown(ctx context.Context, s *Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-quit

	s.logger.Info("shutting down...")

	ctx, shutdown := context.WithTimeout(ctx, s.cfg.API.GracefulTimeout*time.Second)
	defer shutdown()

	return s.httpServer.Shutdown(ctx)
}
