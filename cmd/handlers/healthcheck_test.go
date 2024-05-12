package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	expectedStatus      = "ok"
	expectedType        = "handlers.HealthCheckResponse"
	expectedContentType = "application/json"
)

func setupTest() (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	rr := httptest.NewRecorder()
	return req, rr
}

func TestHealthCheckHandler(t *testing.T) {
	req, rr := setupTest()
	handler := http.HandlerFunc(HealthCheckHandler)

	t.Run("returns ok status", func(t *testing.T) {
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		var response HealthCheckResponse
		_ = json.NewDecoder(rr.Body).Decode(&response)

		if reflect.TypeOf(response).String() != expectedType {
			t.Errorf("handler returned unexpected type: got %v want %v",
				reflect.TypeOf(response).String(), expectedType)
		}

		if response.Status != expectedStatus {
			t.Errorf("handler returned unexpected status: got %v want %v",
				response.Status, expectedStatus)
		}
	})

	t.Run("returns json content type", func(t *testing.T) {
		handler.ServeHTTP(rr, req)

		if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
			t.Errorf("content type header does not match: got %v want %v",
				contentType, expectedContentType)
		}
	})

	t.Run("returns correct response body", func(t *testing.T) {
		handler.ServeHTTP(rr, req)

		var response HealthCheckResponse
		_ = json.NewDecoder(rr.Body).Decode(&response)

		if response.Status != expectedStatus {
			t.Errorf("handler returned unexpected status: got %v want %v",
				response.Status, expectedStatus)
		}
	})
}
