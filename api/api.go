package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Read - Reads request body and parses JSON to struct
func Read(w http.ResponseWriter, r *http.Request) interface{} {
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return nil
	}
	var i interface{}
	// Parse JSON
	err = json.Unmarshal(body, &i)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return nil
	}
	return i
}
