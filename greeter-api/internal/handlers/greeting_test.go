package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetingWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Ada", nil)
	w := httptest.NewRecorder()

	Greeting(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got greetingResponse
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Message != "Hello, Ada!" {
		t.Fatalf("expected %q, got %q", "Hello, Ada!", got.Message)
	}
}

func TestGreetingMissingName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	Greeting(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
	var got errorResponse
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Error == "" {
		t.Fatal("expected a non-empty error message")
	}
}

func TestGreetingEmptyName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=", nil)
	w := httptest.NewRecorder()

	Greeting(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
