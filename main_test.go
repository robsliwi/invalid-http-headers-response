package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(invalidHeaders)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if invalidHeader := rr.Header().Get("INVALID HEADAH"); invalidHeader != "invalid" {
		t.Errorf("INVALID HEADAH does not match: got %v want %v",
			invalidHeader, "invalid")
	}
}
