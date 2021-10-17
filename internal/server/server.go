package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"

	"github.com/sledro/goapp/internal/secrets"
	"github.com/sledro/goapp/internal/services"
	"github.com/sledro/goapp/internal/store"
	"github.com/sledro/goapp/pkg/logger"
)

// Server type
type server struct {
	r        *mux.Router
	log      *logrus.Logger
	db       *sqlx.DB
	secrets  secrets.Secrets
	services *Services
	stores   *Stores
}

// Services - These are injected into the server
// and can be called by the handlers
type Services struct {
	AuthService services.AuthServiceInterface
	UserService services.UserServiceInterface
}

// Stores - These are injected into the server
// and can be called by the handlers
type Stores struct {
	AuthStore store.AuthStoreInterface
	UserStore store.UserStoreInterface
}

// NewServer - Create Server
func NewServer(secretName, region string) server {
	// Create a new server
	s := server{}
	// Inject logger
	s.log = logger.NewLogger()
	// Inject secrets
	var err error
	s.secrets, err = secrets.LoadSecrets(secretName, region)
	if err != nil {
		s.log.Error("error:", err)
	}
	// Inject database
	s.db = store.NewDatabase(s.secrets.DBUser, s.secrets.DBPass, s.secrets.DBHost, s.secrets.DBPort, s.secrets.DBDatabase)
	// Inject stores
	store.UserStoreInstance = &store.UserStore{DB: s.db}
	store.AuthStoreInstance = &store.AuthStore{DB: s.db}
	s.stores = &Stores{
		UserStore: store.UserStoreInstance,
		AuthStore: store.AuthStoreInstance,
	}
	// Inject services
	services.UserServiceInstance = &services.UserService{UserStore: s.stores.UserStore}
	services.AuthServiceInstance = &services.AuthService{DB: s.db, JWTSecret: s.secrets.JWTSecret, UserService: services.UserServiceInstance}

	s.services = &Services{
		AuthService: services.AuthServiceInstance,
		UserService: services.UserServiceInstance,
	}
	// Inject routes
	s.r = mux.NewRouter().StrictSlash(true)
	return s
}

// StartServer - Load routes into server and
// starts HTTP server
func (s *server) StartServer() {
	log.Info("📡 Server Started. API Server is now listening on http://localhost:8080")
	s.routes()
	log.Fatal(http.ListenAndServe(":8080", s.r))
}

// ServeHTTP - Turns server into http server
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
