package server

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleAuthLogin(t *testing.T) {
	// Test Cases
	cases := []HandlerTestCase{
		{
			Name:        "Test Happy Path",
			Route:       "http://0.0.0.0:8080/v1/auth",
			Method:      "POST",
			ContentType: "application/json",
			Auth:        "",
			Body: `{
				"email":"john@doe.com",
				"password":"12345678"
			}`,
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name:               "Test No JSON Body",
			Route:              "http://0.0.0.0:8080/v1/auth",
			Method:             "POST",
			ContentType:        "application/json",
			Auth:               "",
			Body:               ``,
			ExpectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			Name:        "Test User Not Found",
			Route:       "http://0.0.0.0:8080/v1/auth",
			Method:      "POST",
			ContentType: "application/json",
			Auth:        "",
			Body: `{
				"email":"test@test.com",
				"password":"12345678"
			}`,
			ExpectedStatusCode: http.StatusUnauthorized,
			TestFunc: func(body []byte, t *testing.T) {
				var res map[string]interface{}
				json.Unmarshal(body, &res)
				assert.Equal(t, map[string]interface{}{"error": "sql: no rows in result set"}, res)
			},
		},
	}

	// Run tests
	for _, c := range cases {
		t.Run(c.Name, c.GenericHandlerTestFunc())
	}
}
