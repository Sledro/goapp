package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
	"github.com/sledro/golang-framework/internal/storage"
	"github.com/sledro/golang-framework/pkg/logger"
	"github.com/sledro/golang-framework/pkg/secrets"
	"gorm.io/gorm"
)

// Server type
type server struct {
	r       *mux.Router
	log     *logrus.Logger
	db      *gorm.DB
	secrets map[string]string
}

// NewServer - Create Server
func NewServer(secretName, region string) server {
	s := server{}
	s.log = logger.NewLogger()
	s.secrets = secrets.LoadSecrets(secretName, region)
	s.db = storage.NewDatabase(s.secrets["DB_USER"], s.secrets["DB_PASS"], s.secrets["DB_HOST"], s.secrets["DB_PORT"], s.secrets["DB_DATABASE"])
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
