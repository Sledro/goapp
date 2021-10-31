package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/sledro/goapp/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestHandleAuthLogin(t *testing.T) {
	// Test Cases
	cases := []HandlerTestCase{
		{
			Name:        "Test Happy Path",
			Route:       "/v1/auth",
			Method:      http.MethodPost,
			ContentType: "application/json",
			Auth:        "",
			Body: `{
				"email":"test@test.com",
				"password":"12345678"
			}`,
			ExpectedStatusCode: http.StatusOK,
			MockFunc: func(c *HandlerTestCase) {
				c.MockDB.ExpectQuery(store.LoginQuery).WithArgs("test@test.com").
					WillReturnRows(c.MockDB.NewRows([]string{"email", "username"}).AddRow("test@test.com", "username"))
			},
			TestFunc: func(body []byte, t *testing.T) {
				var res map[string]interface{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, err, nil)
			},
		},
		{
			Name:               "Test Could not decode JSON body",
			Route:              "/v1/auth",
			Method:             http.MethodPost,
			ContentType:        "application/json",
			Auth:               "",
			Body:               `-}`,
			ExpectedStatusCode: http.StatusUnprocessableEntity,
			TestFunc: func(body []byte, t *testing.T) {
				var res map[string]interface{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, err, nil)
				assert.Equal(t, map[string]interface{}{"error": "could not decode JSON body"}, res)
			},
		},
		{
			Name:        "Test User Not Found",
			Route:       "/v1/auth",
			Method:      http.MethodPost,
			ContentType: "application/json",
			Auth:        "",
			Body: `{
				"email":"test@test.com",
				"password":"12345678"
			}`,
			ExpectedStatusCode: http.StatusUnauthorized,
			MockFunc: func(c *HandlerTestCase) {
				c.MockDB.ExpectQuery(store.LoginQuery).WithArgs("test@test.com").
					WillReturnError(errors.New("user does not exist"))
			},
			TestFunc: func(body []byte, t *testing.T) {
				var res map[string]interface{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, err, nil)
				assert.Equal(t, map[string]interface{}{"error": "user does not exist"}, res)
			},
		},
	}

	// Run tests
	for _, c := range cases {
		t.Run(c.Name, c.GenericHandlerTestFunc())
	}
}
