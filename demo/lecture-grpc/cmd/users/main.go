package main

import (
	"github.com/sirupsen/logrus"

	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/users"
)

func main() {
	usersService := users.NewUsersService()
	if err := usersService.Run(); err != nil {
		logrus.New().WithError(err).Fatal("failed to run users service")
	}
}
