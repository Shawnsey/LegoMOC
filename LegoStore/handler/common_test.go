package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSONResponse_Success(t *testing.T) {
	// Arrange
	data := map[string]string{"message": "Hello, World!"}
	expectedStatus := http.StatusOK

	// Setup response recorder
	rr := httptest.NewRecorder()

	// Act
	writeJSONResponse(rr, data, expectedStatus)

	// Assert
	if rr.Code != expectedStatus {
		t.Errorf("Expected status %d, got %d", expectedStatus, rr.Code)
	}

	expectedContentType := "application/json"
	if rr.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, rr.Header().Get("Content-Type"))
	}

	// Check response body
	var responseData map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseData)
	if err != nil || responseData["message"] != "Hello, World!" {
		t.Errorf("Expected JSON response with message 'Hello, World!', got %v", rr.Body.String())
	}
}

func TestWriteJSONResponse_MarshalError(t *testing.T) {
	// Arrange
	data := make(chan int) // Channels cannot be marshaled to JSON
	expectedStatus := http.StatusInternalServerError

	// Setup response recorder
	rr := httptest.NewRecorder()

	// Act
	writeJSONResponse(rr, data, http.StatusOK)

	// Assert
	if rr.Code != expectedStatus {
		t.Errorf("Expected status %d, got %d", expectedStatus, rr.Code)
	}
}

func TestWriteJSONResponse_WriteError(t *testing.T) {
	// Arrange
	data := map[string]string{"message": "Goodbye, World!"}
	expectedStatus := http.StatusCreated

	// Setup response recorder
	rr := httptest.NewRecorder()

	// Act
	writeJSONResponse(rr, data, expectedStatus)

	// Assert
	if rr.Code != expectedStatus {
		t.Errorf("Expected status %d, got %d", expectedStatus, rr.Code)
	}
}
