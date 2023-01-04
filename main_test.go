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
	// create a test server
	ts := httptest.NewServer(router())
	defer ts.Close()

	// test the POST /ps route
	res, err := http.Post(ts.URL+"/ps", "application/json", bytes.NewBuffer([]byte(`{"name":"test"}`)))
	if err != nil {
		t.Errorf("Error making POST request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	expected := `{"msg":{"name":"test"}}`
	if string(body) != expected {
		t.Errorf("Expected response body %q; got %q", expected, string(body))
	}

	// test the POST /ps route with a request that should fail validation
	res, err = http.Post(ts.URL+"/ps", "application/json", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Errorf("Error making POST request: %v", err)
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	expected = `{"msg":"error"}`
	if string(body) != expected {
		t.Errorf("Expected response body %q; got %q", expected, string(body))
	}
}
