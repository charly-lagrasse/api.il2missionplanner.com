package handlers

import (
	"net/http"
	"testing"
)

func TestServersHandlerHandler(t *testing.T) {
	expectedBody := `{"data":["testserver"],"error":""}`
	testHandler("GET", "/servers", http.StatusOK, expectedBody, ServersHandler, t)
}

func TestServerStateHandlerHandler(t *testing.T) {
	// TODO Mock downstream http request
}
