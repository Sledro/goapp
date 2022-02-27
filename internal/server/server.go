package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"

	"github.com/sledro/goapp/internal/logger"
	"github.com/sledro/goapp/internal/secrets"
	"github.com/sledro/goapp/internal/services"
	"github.com/sledro/goapp/internal/store"
)

// Server type
type server struct {
	r        chi.Router
	log      *logrus.Logger
	db       *sqlx.DB
	secrets  secrets.Secrets
	services *Services
	stores   *Stores
}

// These are injected into the server
// and can be called by the handlers
type Services struct {
	AuthService services.AuthServiceInterface
	UserService services.UserServiceInterface
}

// These are injected into the server
// and can be called by the handlers
type Stores struct {
	AuthStore store.AuthStoreInterface
	UserStore store.UserStoreInterface
}

// Create Server
func NewServer() (server, error) {
	// Create a new server
	s := server{}
	// Inject logger
	s.log = logger.NewLogger()
	// Inject secrets
	var err error
	s.secrets, err = secrets.LoadSecrets()
	if err != nil {
		return s, err
	}
	// Inject database
	s.db = store.NewDatabase(s.secrets.Database.User, s.secrets.Database.Pass, s.secrets.Database.Host, s.secrets.Database.Port, s.secrets.Database.Name)
	// Inject stores
	store.UserStoreInstance = &store.UserStore{DB: s.db}
	store.AuthStoreInstance = &store.AuthStore{DB: s.db}
	s.stores = &Stores{
		UserStore: store.UserStoreInstance,
		AuthStore: store.AuthStoreInstance,
	}
	// Inject services
	services.UserServiceInstance = &services.UserService{UserStore: s.stores.UserStore}
	services.AuthServiceInstance = &services.AuthService{DB: s.db, JWTSecret: s.secrets.JWTSecret, UserService: services.UserServiceInstance, AuthStore: s.stores.AuthStore}
	s.services = &Services{
		AuthService: services.AuthServiceInstance,
		UserService: services.UserServiceInstance,
	}
	return s, err
}

// Load routes into server and
// starts HTTP server
func (s *server) StartServer() {
	log.Info("📡 Server Started. API Server is now listening on http://localhost:" + s.secrets.AppPort)
	s.routes()
	log.Fatal(http.ListenAndServe(":"+s.secrets.AppPort, s.r))
}

// Turns server into http server
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
