package main

import (
	"github.com/sirupsen/logrus"

	"github.com/apaliavy/godel-golang/demo/lecture-grpc/app/gateway"
)

func main() {
	application := gateway.NewService()
	if err := application.Run(); err != nil {
		logrus.New().WithError(err).Fatal("failed to start service")
	}
}
