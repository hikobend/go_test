package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestPingRouter(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// ...
	assert.Equal(t, w.Body.String(), "{\"msg\":\"pong\"}")
}

func TestRouter(t *testing.T) {
	tests := []struct {
		name string
		body string
		code int
		resp string
	}{
		{
			name: "POST /ps with valid request body",
			body: `{"name":"test"}`,
			code: http.StatusOK,
			resp: `{"msg":{"name":"test"}}`,
		},
		{
			name: "POST /ps with invalid request body",
			body: `{}`,
			code: http.StatusBadRequest,
			resp: `{"msg":"error"}`,
		},
	}

	// create a test server
	ts := httptest.NewServer(router())
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := http.Post(ts.URL+"/ps", "application/json", bytes.NewBuffer([]byte(tt.body)))
			if err != nil {
				t.Errorf("Error making POST request: %v", err)
			}
			if res.StatusCode != tt.code {
				t.Errorf("Expected status %d; got %v", tt.code, res.StatusCode)
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error reading response body: %v", err)
			}
			if string(body) != tt.resp {
				t.Errorf("Expected response body %q; got %q", tt.resp, string(body))
			}
		})
	}
}
