package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/murilosolino/challenge-backend-7/internal/controllers"
	"github.com/murilosolino/challenge-backend-7/internal/dto"
)

type DestinationSvcSpy struct {
	GeneratedFor string
	CreatedMap   map[string]any
	ListCalled   bool
	DeletedId    int
	UpdatedId    int
	UpdatedMap   map[string]interface{}
	SearchedName string
}

func (s *DestinationSvcSpy) CreateDestination(m map[string]interface{}) error {
	s.CreatedMap = map[string]any{}
	for k, v := range m {
		s.CreatedMap[k] = v
	}
	return nil
}

func (s *DestinationSvcSpy) ListDestinations() ([]dto.Destination, error) {
	s.ListCalled = true
	return []dto.Destination{{Id: 1, Name: "Paris"}}, nil
}
func (s *DestinationSvcSpy) DeleteDestinationById(id int) error {
	s.DeletedId = id
	return nil
}
func (s *DestinationSvcSpy) UpdateDestination(id int, m map[string]interface{}) error {
	s.UpdatedId = id
	s.UpdatedMap = m
	return nil
}
func (s *DestinationSvcSpy) FindByName(name string) (dto.Destination, error) {
	s.SearchedName = name
	return dto.Destination{Id: 1, Name: name}, nil
}

func (s *DestinationSvcSpy) GenerateIADescriptiveText(destination string) string {
	s.GeneratedFor = destination
	return "descricao gerada (mock)"
}

func TestCreateNewDestination_GeneratesDescriptiveTextWhenMissing(t *testing.T) {
	spy := &DestinationSvcSpy{}
	controller := controllers.NewDestinationController(spy)

	body := map[string]any{"name": "Paris", "price": 1000}
	payload, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/destinations", bytes.NewBuffer(payload))
	rec := httptest.NewRecorder()

	controller.CreateNewDestination(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperado status %d, recebido %d", http.StatusCreated, rec.Code)
	}
	if spy.GeneratedFor != "Paris" {
		t.Fatalf("esperado gerar descricao para %q, recebeu %q", "Paris", spy.GeneratedFor)
	}
	if spy.CreatedMap == nil {
		t.Fatalf("esperado CreateDestination ser chamado")
	}
	if got := spy.CreatedMap["descriptive_text"]; got != "descricao gerada (mock)" {
		t.Fatalf("esperado descriptive_text %q, recebeu %#v", "descricao gerada (mock)", got)
	}
}

func TestListAllDestinations(t *testing.T) {
	spy := &DestinationSvcSpy{}
	controller := controllers.NewDestinationController(spy)

	req := httptest.NewRequest(http.MethodGet, "/destinations", nil)
	rec := httptest.NewRecorder()

	controller.ListAllDestinations(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if !spy.ListCalled {
		t.Fatalf("expected ListDestinations to be called")
	}
}

func TestListAllDestinationsByName(t *testing.T) {
	spy := &DestinationSvcSpy{}
	controller := controllers.NewDestinationController(spy)

	req := httptest.NewRequest(http.MethodGet, "/destinations?name=Paris", nil)
	rec := httptest.NewRecorder()

	controller.ListAllDestinations(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if spy.SearchedName != "Paris" {
		t.Fatalf("expected to search by name Paris, got %q", spy.SearchedName)
	}
}

func TestDeleteDestination(t *testing.T) {
	spy := &DestinationSvcSpy{}
	controller := controllers.NewDestinationController(spy)

	req := httptest.NewRequest(http.MethodDelete, "/destinations/1", nil)
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	controller.DeleteDestination(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
	if spy.DeletedId != 1 {
		t.Fatalf("expected to delete id 1, got %d", spy.DeletedId)
	}
}

func TestUpdateDestination(t *testing.T) {
	spy := &DestinationSvcSpy{}
	controller := controllers.NewDestinationController(spy)

	body := map[string]any{"name": "Updated Paris", "price": 1200}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPut, "/destinations/1", bytes.NewBuffer(payload))
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	controller.UpdateDestination(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
	if spy.UpdatedId != 1 {
		t.Fatalf("expected to update id 1, got %d", spy.UpdatedId)
	}
	if spy.UpdatedMap == nil {
		t.Fatalf("expected UpdateDestination to be called with a map")
	}
}
