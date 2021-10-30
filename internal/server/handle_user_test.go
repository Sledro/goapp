package server

import (
	"encoding/json"
	"testing"

	"github.com/sledro/goapp/api"
	"github.com/stretchr/testify/assert"
)

func TestHandleUserCreate(t *testing.T) {
	// Test Cases
	cases := []HandlerTestCase{
		{
			Name:        "Test Happy Path 200 Success",
			Route:       "http://127.0.0.1:8080/v1/user",
			Method:      "POST",
			ContentType: "application/json",
			Auth:        "",
			Body: `{
				"email":"test@test.com",
				"password":"12345678"
			}`,
			ExpectedStatusCode: 200,
			MockFunc: func(c *HandlerTestCase) {
				//c.MockDB.ExpectQuery()
			},
			TestFunc: func(body []byte, t *testing.T) {
				var res api.Response
				err := json.Unmarshal(body, &res)
				assert.Equal(t, err, nil)
				//assert.NotContains(t, "res", res.Response)
			},
		},
	}

	// Run tests
	for _, c := range cases {
		t.Run(c.Name, c.GenericHandlerTestFunc())
	}
}
