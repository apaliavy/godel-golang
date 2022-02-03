package users

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	pb "github.com/apaliavy/godel-golang/demo/lecture-grpc/app/users/api"
	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/users/config"
	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/users/handlers"
)

type Service struct {
	grpcServer *grpc.Server
	cfg        *config.AppConfig
}

func NewUsersService() *Service {
	return &Service{
		grpcServer: grpc.NewServer(grpcOpts()...),
		cfg:        config.Load(),
	}
}

func (s *Service) Run() error {
	pb.RegisterUsersServer(s.grpcServer, &handlers.Users{})

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
	if err != nil {
		return errors.Wrap(err, "failed to setup gRPC listener")
	}

	return s.grpcServer.Serve(listener)
}

func grpcOpts() []grpc.ServerOption {
	return make([]grpc.ServerOption, 0)
}
