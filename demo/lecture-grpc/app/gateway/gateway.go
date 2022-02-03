package gateway

import (
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	pbAuth "github.com/apaliavy/godel-golang/demo/lecture-grpc/app/auth/api"
	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/gateway/config"
	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/gateway/handlers"
	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/gateway/server"
	pbUsers "github.com/apaliavy/godel-golang/demo/lecture-grpc/app/users/api"
)

type Service struct {
	cfg          *config.AppConfig
	server       *server.Server
	proxyHandler *handlers.UsersProxy
}

func NewService() *Service {
	return &Service{
		cfg: config.Load(),
	}
}

func (s *Service) Run() error {
	usersClientConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.cfg.UsersServiceRPC.Host, s.cfg.UsersServiceRPC.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return errors.Wrap(err, "failure establishing auth gRPC client conn")
	}

	proxy := handlers.NewUsersProxy(pbUsers.NewUsersClient(usersClientConn))

	authClientConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.cfg.AuthServiceRPC.Host, s.cfg.AuthServiceRPC.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return errors.Wrap(err, "failure establishing auth gRPC client conn")
	}

	srv := server.New(
		server.WithUsersProxy(pbAuth.NewAuthClient(authClientConn), proxy),
	)
	return srv.Run()
}
