package handlers

import (
	"context"
	"strings"

	pb "github.com/apaliavy/godel-golang/demo/lecture-grpc/app/auth/api"
)

// Auth implements pb.AuthServer
type Auth struct {
	pb.UnimplementedAuthServer
}

// IsAuthenticated returns true if token non-empty, false otherwise
// so secure very wow
func (a *Auth) IsAuthenticated(
	ctx context.Context,
	request *pb.IsAuthenticatedRequest,
) (*pb.IsAuthenticatedResponse, error) {
	return &pb.IsAuthenticatedResponse{
		Authenticated: strings.TrimSpace(request.Token) != "",
	}, nil
}
