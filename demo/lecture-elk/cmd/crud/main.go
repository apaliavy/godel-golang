package main

import (
	"github.com/apaliavy/godel-golang/demo/lecture-elk/app/handlers"
	"github.com/apaliavy/godel-golang/demo/lecture-elk/app/server"
	"github.com/apaliavy/godel-golang/demo/lecture-elk/app/users"
)

func main() {
	s := server.New(
		server.WithUsersHandler(handlers.NewUsersHandler(&users.Repository{})),
	)
	s.Run()
}
