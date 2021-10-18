package server

import (
	"encoding/json"
	"testing"

	"github.com/sledro/goapp/api"
	"github.com/stretchr/testify/assert"
)

func TestHandleAuthLogin(t *testing.T) {
	// Test Cases
	cases := []HandlerTestCase{
		{
			Name:        "Test Happy Path 200 Success",
			Route:       "http://0.0.0.0:8080/v1/auth",
			Method:      "POST",
			ContentType: "application/json",
			Auth:        "",
			Body: `{
				"email":"test@dans6s.com",
				"password":"12345678"
			}`,
			ExpectedStatusCode: 200,
			MockFunc: func(c *HandlerTestCase) {
				//c.MockDB.ExpectQuery()
			},
			TestFunc: func(body []byte, t *testing.T) {
				var res api.Response
				json.Unmarshal(body, &res)
				t.Fatalf("error statudds code %v", res.Response)
				assert.NotContains(t, "res", res.Response)
			},
		},
	}

	// Run tests
	for _, c := range cases {
		t.Run(c.Name, c.GenericHandlerTestFunc())
	}
}
