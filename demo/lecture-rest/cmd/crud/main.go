package main

import (
	"github.com/apaliavy/godel-golang/demo/lecture-rest/app/handlers"
	"github.com/apaliavy/godel-golang/demo/lecture-rest/app/server"
	"github.com/apaliavy/godel-golang/demo/lecture-rest/app/users"
)

func main() {
	s := server.New(
		server.WithUsersHandler(handlers.NewUsersHandler(&users.Repository{})),
	)
	s.Run()
}
