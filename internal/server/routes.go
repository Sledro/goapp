package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/sledro/goapp/api"
	middl "github.com/sledro/goapp/internal/middleware"
)

// routes - Setups chi router, middlewares and defines all api endpoints
func (s *server) routes() {
	// Inject routes
	s.r = chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	s.r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Inject chi middleware
	// A good base middleware stack
	// Injects a request ID into the context of each request
	s.r.Use(middleware.RequestID)
	// Sets a http.Request's RemoteAddr to either X-Real-IP or X-Forwarded-For
	s.r.Use(middleware.RealIP)
	// Logs the start and end of each request with the elapsed processing time
	s.r.Use(middleware.Logger)
	// Gracefully absorb panics and prints the stack trace
	s.r.Use(middleware.Recoverer)
	// Sets HTTP response headers as content type JSON
	s.r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.r.Use(middleware.Timeout(60 * time.Second))

	// setup v1 subrouter
	s.r.Route("/v1", func(r chi.Router) {

		// health
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			api.JSON(w, 200, map[string]interface{}{"health_status": "online", "string": "test", "int": 3, "float": 1.32, "bool": true})
		})

		// auth
		r.Route("/auth", func(r chi.Router) {
			r.Post("/", s.handleAuthLogin) // POST /users
		})

		// user
		r.Route("/user", func(r chi.Router) {
			r.Post("/", s.handleUserCreate)
			r.Get("/", s.handleUserGet)
			r.Put("/{id}", middl.Auth(s.handleUserUpdate, s.secrets.JWTSecret))
			r.Delete("/{id}", middl.Auth(s.handleUserDelete, s.secrets.JWTSecret))
			r.Get("/list", middl.Auth(s.handleUserList, s.secrets.JWTSecret))
		})

	})
}
