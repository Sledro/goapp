package main

import (
	"fmt"

	"github.com/sledro/goapp/internal/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		fmt.Println("could create server:", err)
	} else {
		server.StartServer()
	}
}
