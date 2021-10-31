package server

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sledro/goapp/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestHandleUserCreate(t *testing.T) {
	// Test Cases
	cases := []HandlerTestCase{
		{
			Name:               "Test Could not decode JSON body",
			Route:              "/v1/user",
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
			Name:        "Test User already exists",
			Route:       "/v1/user",
			Method:      http.MethodPost,
			ContentType: "application/json",
			Auth:        "",
			Body: `{
				"firstname": "Jon",
				"lastname": "Doe",
				"username": "Superman",
				"password": "12345678",
				"email": "jon@doe.com"
			}`,
			ExpectedStatusCode: http.StatusConflict,
			MockFunc: func(c *HandlerTestCase) {
				c.MockDB.ExpectQuery(store.GetUserQuery).WithArgs(3, "Superman", "jon@doe.com").
					WillReturnRows(c.MockDB.NewRows([]string{"id", "firstname", "lastname", "username", "email"}).
						AddRow(3, "Jon", "Doe", "Superman", "jon@doe.com"))
			},
			TestFunc: func(body []byte, t *testing.T) {
				var res map[string]interface{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, err, nil)
				assert.Equal(t, map[string]interface{}{"error": "this user already exists"}, res)
			},
		},
	}

	// Run tests
	for _, c := range cases {
		t.Run(c.Name, c.GenericHandlerTestFunc())
	}
}
