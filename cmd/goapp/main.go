package main

import (
	"github.com/sledro/goapp/internal/server"
)

func main() {
	server := server.NewServer("goapp", "eu-west-1")
	server.StartServer()
}
