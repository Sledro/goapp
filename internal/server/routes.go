package server

import (
	"github.com/sledro/goapp/internal/middleware"
)

// routes - defines all api endpoints
func (s *server) routes() {
	// setup v1 subrouter
	v1 := s.r.PathPrefix("/v1").Subrouter()

	// api routes
	v1.HandleFunc("/api/health", middleware.Headers(s.handleHealthCheck)).Methods("GET")

	// auth routes
	v1.HandleFunc("/auth", middleware.Headers(s.handleAuthLogin)).Methods("POST")

	// user routes CRUDL
	v1.HandleFunc("/user", middleware.Headers(s.handleUserCreate)).Methods("POST")
	v1.HandleFunc("/user", middleware.Headers(s.handleUserGet)).Methods("GET")
	v1.HandleFunc("/user/{id}", middleware.Headers(s.handleUserUpdate)).Methods("PUT")
	v1.HandleFunc("/user/{id}", middleware.Headers(middleware.Auth(s.handleUserDelete, s.secrets.JWTSecret))).Methods("DELETE")
	v1.HandleFunc("/user/list", middleware.Headers(s.handleUserList)).Methods("GET")
}
