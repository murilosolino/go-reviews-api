package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/murilosolino/challenge-backend-7/api/controllers"
	"github.com/murilosolino/challenge-backend-7/api/model"
)

type MockService struct{}

func (ms MockService) GetAllReviews() ([]model.ReviewsRow, error) {
	return nil, nil
}

func (ms MockService) Get3RandomReviews() ([]model.ReviewsRow, error) {
	return nil, nil
}

func (ms MockService) CreateReview(rev map[string]interface{}) error {
	return nil
}

func (ms MockService) UpdateReview(id int, r map[string]interface{}) error {
	return nil
}

func (ms MockService) ExceludeReview(id int) error {
	return nil
}

func (ms MockService) SearchById(id int) (model.ReviewsRow, error) {
	return model.ReviewsRow{}, nil
}

func (ms MockService) CreateDestination(m map[string]interface{}) error {
	return nil
}

func (ms MockService) ListDestinations() ([]model.DestinationRow, error) {
	return nil, nil
}

func (ms MockService) DeleteDestinationById(id int) error {
	return nil
}

func (ms MockService) UpdateDestination(id int, m map[string]interface{}) error {
	return nil
}

func (ms MockService) FindByName(name string) (model.DestinationRow, error) {
	return model.DestinationRow{}, nil
}

var flagtests = []struct {
	name       string
	endpoint   string
	httpMethod string
	body       map[string]any
	out        int
}{
	{"GET /", "/", "GET", nil, http.StatusOK},
	{"GET /reviews", "/reviews", "GET", nil, http.StatusOK},
	{"GET /reviews/{id}", "/reviews/1", "GET", nil, http.StatusOK},
	{"GET /reviews-home", "/reviews-home", "GET", nil, http.StatusOK},
	{"POST /reviews", "/reviews", "POST", map[string]any{"review": "otimo", "author": "murilo", "url_photo": "http://localhost/photo.jpg"}, http.StatusCreated},
	{"PUT /reviews/{id}", "/reviews/1", "PUT", map[string]any{"review": "atualizado"}, http.StatusNoContent},
	{"PATCH /reviews/{id}", "/reviews/1", "PATCH", map[string]any{"author_name": "novo nome"}, http.StatusNoContent},
	{"DELETE /review/{id}", "/review/1", "DELETE", nil, http.StatusNoContent},

	{"GET /destinations", "/destinations", "GET", nil, http.StatusOK},
	{"POST /reviews", "/reviews", "POST", map[string]any{"img": nil, "name": "São Paulo", "price": 1000}, http.StatusCreated},
	{"PUT /destinations/{id}", "/destinations/1", "PUT", map[string]any{"name": "Paris"}, http.StatusNoContent},
	{"DELETE /destinations/{id}", "/destinations/1", "DELETE", nil, http.StatusNoContent},
}

func setupRouter() *http.ServeMux {
	reviewController := controllers.NewReviewController(MockService{})
	hcController := controllers.NewHealthCheck()
	controllerDest := controllers.NewDestinationController(MockService{})

	r := http.NewServeMux()
	r.HandleFunc("GET /", hcController.HealthCheck)
	r.HandleFunc("GET /reviews", reviewController.ListAllReviews)
	r.HandleFunc("GET /reviews/{id}", reviewController.FindReviewById)
	r.HandleFunc("GET /reviews-home", reviewController.GetRandomReviews)
	r.HandleFunc("POST /reviews", reviewController.CreateNewReview)
	r.HandleFunc("PUT /reviews/{id}", reviewController.UpdateReview)
	r.HandleFunc("PATCH /reviews/{id}", reviewController.UpdateReview)
	r.HandleFunc("DELETE /review/{id}", reviewController.ExceludeReview)

	r.HandleFunc("POST /destinations", controllerDest.CreateNewDestination)
	r.HandleFunc("GET /destinations", controllerDest.ListAllDestinations)
	r.HandleFunc("DELETE /destinations/{id}", controllerDest.DeleteDestination)
	r.HandleFunc("PUT /destinations/{id}", controllerDest.UpdateDestination)

	return r
}

func TestHttpStatusCodeForReviewsEndpoint(t *testing.T) {
	r := setupRouter()

	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			var payload io.Reader
			if tt.body != nil {
				body, _ := json.Marshal(tt.body)
				payload = bytes.NewBuffer(body)
			}

			req, err := http.NewRequest(tt.httpMethod, tt.endpoint, payload)
			if err != nil {
				t.Fatalf("erro ao montar requisicao %s", err.Error())
			}

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			if rec.Code != tt.out {
				t.Errorf("esperado %d, recebido %d", tt.out, rec.Code)
			}
		})
	}
}
