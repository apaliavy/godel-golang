package log

import (
	"net"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func DefaultLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}

	log := logrus.New()
	conn, err := net.Dial("tcp", ":8089")
	if err != nil {
		log.Fatal(err)
	}

	log.Hooks.Add(logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{})))

	logger = log
	return logger
}
