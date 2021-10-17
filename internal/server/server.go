package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"

	"github.com/sledro/goapp/internal/secrets"
	"github.com/sledro/goapp/internal/store"
	"github.com/sledro/goapp/pkg/logger"
)

// Server type
type server struct {
	r       *mux.Router
	log     *logrus.Logger
	db      *sqlx.DB
	secrets secrets.Secrets
}

// NewServer - Create Server
func NewServer(secretName, region string) server {
	s := server{}
	s.log = logger.NewLogger()
	var err error
	s.secrets, err = secrets.LoadSecrets(secretName, region)
	if err != nil {
		s.log.Error("error:", err)
	}
	s.db = store.NewDatabase(s.secrets.DBUser, s.secrets.DBPass, s.secrets.DBHost, s.secrets.DBPort, s.secrets.DBDatabase)
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
