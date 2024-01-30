package tests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"bytes"
	"encoding/json"
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

func TestPutStateHandlerInit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/state", PutStateHandlerSuccess)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/state", strings.NewReader("INIT"))
	req.Header.Set("Content-Type", "text/plain")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "text/plain; charset=utf-8", contentType)
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(w.Body)
	if err != nil {
		t.Errorf("Error reading request body")
	}
	msgBody := buf.String()
	if msgBody != "Successfully Updated State" {
		t.Errorf("Expected Response %s, got %s", "Successfully Updated State", msgBody)
	}
}

func TestPutStateHandlerPaused(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/state", PutStateHandlerSuccess)

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

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(w.Body)
	if err != nil {
		t.Errorf("Error reading request body")
	}
	msgBody := buf.String()
	if msgBody != "Successfully Updated State" {
		t.Errorf("Expected Response %s, got %s", "Successfully Updated State", msgBody)
	}
}

func TestPutStateHandlerRunning(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/state", PutStateHandlerSuccess)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/state", strings.NewReader("RUNNING"))
	req.Header.Set("Content-Type", "text/plain")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "text/plain; charset=utf-8", contentType)
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(w.Body)
	if err != nil {
		t.Errorf("Error reading request body")
	}
	msgBody := buf.String()
	if msgBody != "Successfully Updated State" {
		t.Errorf("Expected Response %s, got %s", "Successfully Updated State", msgBody)
	}
}

func TestPutStateHandlerShutdown(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/state", PutStateHandlerSuccess)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/state", strings.NewReader("SHUTDOWN"))
	req.Header.Set("Content-Type", "text/plain")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "text/plain; charset=utf-8", contentType)
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(w.Body)
	if err != nil {
		t.Errorf("Error reading request body")
	}
	msgBody := buf.String()
	if msgBody != "Successfully Updated State" {
		t.Errorf("Expected Response %s, got %s", "Successfully Updated State", msgBody)
	}
}

func TestPutStateHandlerInvalid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/state", PutStateHandlerFailure)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/state", strings.NewReader("INVALID"))
	req.Header.Set("Content-Type", "text/plain")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "text/plain; charset=utf-8", contentType)
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(w.Body)
	if err != nil {
		t.Errorf("Error reading request body")
	}
	msgBody := buf.String()
	if msgBody != "Invalid State Value" {
		t.Errorf("Expected Response %s, got %s", "Invalid State Value", msgBody)
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

func TestGetMQStatisticHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/mqstatistic", GetMQStatisticHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mqstatistic", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("Expected Content-Type %s, got %s", "application/json; charset=utf-8", contentType)
	}

	var actualResponse StatResponse
	_ = json.NewDecoder(w.Body).Decode(&actualResponse)

	if actualResponse.OverallStats.ClusterName != "cluster-1" {
		t.Errorf("Expected Cluster-Name %s, got %s", "cluster-1", actualResponse.OverallStats.ClusterName)
	}
}


