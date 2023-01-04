package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	mockUserResp := `{"message":"hello world"}`
	// テスト用のサーバーを立てる
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()
	// リクエストを送れるか?
	resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()
	// Statusコードは200か?
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	if string(responseData) != mockUserResp {
		t.Fatalf("Expected hello world message, got %v", responseData)
	}
}
