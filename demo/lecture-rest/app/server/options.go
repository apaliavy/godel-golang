package server

import (
	"github.com/go-chi/chi/v5"

	"github.com/apaliavy/godel-golang/demo/lecture-rest/app/handlers"
)

func WithUsersHandler(uh *handlers.Users) Options {
	return func(s *Server) error {
		s.router.Route("/users", func(r chi.Router) {
			r.Post("/", uh.Create)
			r.Get("/{id}", uh.Get)
			r.Put("/{id}", uh.Update)
			r.Patch("/{id}", uh.Modify)
			r.Delete("/{id}", uh.Delete)
		})
		return nil
	}
}
