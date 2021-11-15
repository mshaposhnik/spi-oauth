// main_test.go
package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
	"time"
)

var client = &http.Client{
	Timeout: 1 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func TestMain(m *testing.M) {
	go start()
	m.Run()
}

func TestBadRequestUrl(t *testing.T) {

	r, _ := http.NewRequest("GET", "http://localhost:8000/abcd", nil)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGitHubRedirect(t *testing.T) {

	r, _ := http.NewRequest("GET", "http://localhost:8000/github/authenticate", nil)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.True(t, strings.HasPrefix(resp.Header.Get("Location"), "https://github.com/login/oauth/authorize"))
}

func TestQuayRedirect(t *testing.T) {

	r, _ := http.NewRequest("GET", "http://localhost:8000/quay/authenticate", nil)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.True(t, strings.HasPrefix(resp.Header.Get("Location"), "https://quay.io/oauth/authorize"))
}
