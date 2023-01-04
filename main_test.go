package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	mockUserResp := `{"message":"hello world"}`
	// テスト用のサーバーを立てる
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()
	// リクエストを送れるか?
	resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))
	// defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	responseData, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockUserResp, string(responseData))
}
