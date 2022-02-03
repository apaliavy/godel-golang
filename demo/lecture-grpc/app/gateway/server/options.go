package server

import (
	"github.com/go-chi/chi/v5"

	pb "github.com/apaliavy/godel-golang/demo/lecture-grpc/app/auth/api"
	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/gateway/handlers"
)

func WithUsersProxy(authClient pb.AuthClient, p *handlers.UsersProxy) Options {
	return func(s *Server) error {
		s.router = setupUsersRouter(s.router, authClient, p)
		return nil
	}
}

func setupUsersRouter(r chi.Router, authClient pb.AuthClient, p *handlers.UsersProxy) chi.Router {
	r.Route("/users", func(r chi.Router) {
		r.Use(authenticatedMiddleware(authClient))
		r.Post("/", p.CreateUser)
		r.Get("/", p.ListUsers)
		r.Get("/{id}", p.GetUser)
	})

	return r
}
