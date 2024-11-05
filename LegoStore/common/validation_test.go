package common

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Test that a valid UUID allows the request to proceed
func TestValidateParams_ValidUUID(t *testing.T) {
	// Arrange
	validUUID := uuid.New().String()

	// Setup request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/uuid/"+validUUID, nil)
	rr := httptest.NewRecorder()

	// Setup chi router and middleware
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", validUUID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Define a handler to be called if middleware succeeds
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	})

	// Apply middleware
	middleware := ValidateParams(nextHandler)
	middleware.ServeHTTP(rr, req)

	// Assert
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}
	expectedBody := "Success"
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, rr.Body.String())
	}
}

// Test that an invalid UUID returns a 400 error
func TestValidateParams_InvalidUUID(t *testing.T) {
	// Arrange
	invalidUUID := "invalid-uuid"

	// Setup request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/uuid/"+invalidUUID, nil)
	rr := httptest.NewRecorder()

	// Setup chi router and middleware
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", invalidUUID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Define a handler that should not be called if middleware fails
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This should not be called"))
	})

	// Apply middleware
	middleware := ValidateParams(nextHandler)
	middleware.ServeHTTP(rr, req)

	// Assert
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
	expectedBody := "Invalid UUID"
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, rr.Body.String())
	}
}