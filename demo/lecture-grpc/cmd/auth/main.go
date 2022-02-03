package main

import (
	"github.com/sirupsen/logrus"

	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/auth"
)

func main() {
	authService := auth.NewAuthService()
	if err := authService.Run(); err != nil {
		logrus.New().WithError(err).Fatal("failed to run auth service")
	}
}
