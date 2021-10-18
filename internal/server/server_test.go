package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/sledro/goapp/internal/store"
)

type HandlerTestCase struct {
	Name               string
	Route              string
	Method             string
	ContentType        string
	Auth               string
	Body               string
	ExpectedStatusCode int
	Error              bool
	MockFunc           func(c *HandlerTestCase)
	MockDB             sqlmock.Sqlmock
	TestFunc           func(body []byte, t *testing.T)
}

// NewTestServer - Creates a new test server
func NewTestServer(t *testing.T) (*httptest.Server, sqlmock.Sqlmock, server) {
	// Create a new server
	server := server{}
	// Create new mock db
	db, mock, err := store.NewTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	// Inject mock db
	server.db = db
	// Inject routes
	server.r = mux.NewRouter().StrictSlash(true)
	return httptest.NewServer(server.r), mock, server
}

// GenericHandlerTestFunc -
func (c *HandlerTestCase) GenericHandlerTestFunc() func(t *testing.T) {
	return func(t *testing.T) {
		client := http.Client{}
		// Create a nenw test server
		s, mockDB, _ := NewTestServer(t)
		defer s.Close()

		// Set
		c.MockDB = mockDB

		// Mock any required functions
		if c.MockFunc != nil {
			c.MockFunc(c)
		}

		// Make HTTP request
		req, err := http.NewRequest(c.Method, c.Route, strings.NewReader(c.Body))
		if err != nil {
			t.Fatal("error making http call", err)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Fatal("error making http call", err)
		}
		defer res.Body.Close()

		// Read request body
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal("error reading the response body", err)
			return
		}

		if res.StatusCode != c.ExpectedStatusCode {
			t.Fatalf("error status code %v not expected %v", res.StatusCode, c.ExpectedStatusCode)
		}

		if c.TestFunc != nil {
			c.TestFunc(b, t)
		}
	}
}
