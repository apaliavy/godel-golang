package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	pb "github.com/apaliavy/godel-golang/demo/lecture-grpc/app/auth/api"
)

func authenticatedMiddleware(ac pb.AuthClient) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("received request")
			resp, err := ac.IsAuthenticated(r.Context(), &pb.IsAuthenticatedRequest{
				Token: r.Header.Get("Bearer"),
			})

			if err != nil {
				logrus.New().WithError(err).Error("failed to check user token - error received")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if !resp.Authenticated {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
