package server

import (
	"net/http"

	"github.com/sledro/golang-framework/api"
)

// handleHealthCheck - can be used to check if the server is online
func (s *server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	// Sort the data to be returned
	sortData := map[string]interface{}{
		"health_status": "OK",
		"string":        "test",
		"int":           3,
		"float":         1.36,
		"bool":          true,
	}
	api.JSON(w, 200, sortData)
}
