package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/amstee/easy-cut/src/status"
)

func TestSalon(t *testing.T) {
	InitConfig()
	req, err := http.NewRequest("GET", "/status", nil); if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(status.SystemStatus)

	handler.ServeHTTP(w, req)

	if res := w.Code; res != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res, http.StatusOK)
	}
	expected := `{"status":"ok","service":"test","version":"1.0.0"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}