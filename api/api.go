package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response - Wrap/format responses
type Response struct {
	StatusCode int         `json:"status_code"`
	Err        bool        `json:"error"`
	Response   interface{} `json:"response"`
}

// JSON - Returns JSON response to the useer. HTTP status code
// and data must be provided
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR - Returns ann error to the user
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
