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

	"github.com/apaliavy/godel-golang/demo/lecture-elk/app/config"
	"github.com/apaliavy/godel-golang/demo/lecture-elk/app/log"
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
		router: chi.NewRouter(),
		logger: log.DefaultLogger(),
	}
}

func (s *Server) Run() {
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

	if err := gracefulShutdown(context.Background(), s); err != nil {
		s.logger.WithError(err).Error("failed to gracefully handle shutdown")
	}
}

func gracefulShutdown(ctx context.Context, s *Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-quit

	s.logger.Info("shutting down...")

	ctx, shutdown := context.WithTimeout(ctx, s.cfg.API.GracefulTimeout*time.Second)
	defer shutdown()

	//_ = s.DB().Close()

	return s.httpServer.Shutdown(ctx)
}
