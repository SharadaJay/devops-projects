package tests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetMessagesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/messages", GetMessagesHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "text/plain; charset=utf-8", contentType)
	}
}

func TestPutStateHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/state", PutStateHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/state", strings.NewReader("PAUSED"))
	req.Header.Set("Content-Type", "text/plain")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "text/plain; charset=utf-8", contentType)
	}
}

func TestGetStateHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/state", GetMessagesHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/state", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "text/plain; charset=utf-8", contentType)
	}
}

func TestGetRunLogHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/run-log", GetRunLogHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/run-log", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "text/plain; charset=utf-8", contentType)
	}
}


