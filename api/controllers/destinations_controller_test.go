package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/murilosolino/challenge-backend-7/api/controllers"
	"github.com/murilosolino/challenge-backend-7/api/model"
)

type DestinationSvcSpy struct {
	GeneratedFor string
	CreatedMap   map[string]any
}

func (s *DestinationSvcSpy) CreateDestination(m map[string]interface{}) error {
	s.CreatedMap = map[string]any{}
	for k, v := range m {
		s.CreatedMap[k] = v
	}
	return nil
}

func (s *DestinationSvcSpy) ListDestinations() ([]model.DestinationRow, error) { return nil, nil }
func (s *DestinationSvcSpy) DeleteDestinationById(id int) error                { return nil }
func (s *DestinationSvcSpy) UpdateDestination(id int, m map[string]interface{}) error {
	return nil
}
func (s *DestinationSvcSpy) FindByName(name string) (model.DestinationRow, error) {
	return model.DestinationRow{}, nil
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
