package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func processRequest(method, url string) (int, string) {
	server := &myHandler{}
	request, _ := http.NewRequest(method, url, nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, request)
	return recorder.Code, recorder.Body.String()
}

func TestHelloRoute(t *testing.T) {
	status, body := processRequest("GET", "/hello")
	if status != 200 {
		t.Fatalf("Received non-200 response: %d\n", status)
	}
	expected := "Hello World!"
	if expected != body {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}

func TestOtherRoute(t *testing.T) {
	status, body := processRequest("GET", "/")
	if status != 200 {
		t.Fatalf("Received non-200 response: %d\n", status)
	}
	expected := "Route with URL: /"
	if expected != body {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}
