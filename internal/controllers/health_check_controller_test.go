package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/murilosolino/challenge-backend-7/internal/controllers"
)

func TestHealthCheck(t *testing.T) {
	controller := controllers.NewHealthCheck()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	controller.HealthCheck(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}
