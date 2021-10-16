package main

import (
	"github.com/sledro/golang-framework/internal/server"
)

func main() {
	server := server.NewServer("golang-framework", "eu-west-1")
	server.StartServer()
}
