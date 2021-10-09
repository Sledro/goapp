package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/util/log"
	"github.com/rs/cors"
)

// Server type
type Server struct {
	router *mux.Router
}

// NewServer - Create Server
func NewServer() Server {
	s := Server{}
	s.router = mux.NewRouter().StrictSlash(true)
	return s
}

// StartServer - Start API
func (s *Server) StartServer() {
	log.Info("📡 Server Started. API Server is now listening on port http://localhost:8080")

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(s.router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

// ServeHTTP - Turns server into http server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
