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

type ReviewSvcSpy struct {
	CreatedMap map[string]interface{}
	UpdatedMap map[string]interface{}
	DeletedId  int
	FoundId    int
}

func (s *ReviewSvcSpy) GetAllReviews() ([]dto.Review, error) {
	return []dto.Review{
		{Id: 1, AuthorName: "User1", Review: "Review1", Url_photo: "pic1.jpg"},
	}, nil
}

func (s *ReviewSvcSpy) Get3RandomReviews() ([]dto.Review, error) {
	return []dto.Review{
		{Id: 1, AuthorName: "User1", Review: "Review1", Url_photo: "pic1.jpg"},
		{Id: 2, AuthorName: "User2", Review: "Review2", Url_photo: "pic2.jpg"},
		{Id: 3, AuthorName: "User3", Review: "Review3", Url_photo: "pic3.jpg"},
	}, nil
}

func (s *ReviewSvcSpy) CreateReview(rev map[string]interface{}) error {
	s.CreatedMap = rev
	return nil
}

func (s *ReviewSvcSpy) UpdateReview(id int, r map[string]interface{}) error {
	s.UpdatedMap = r
	return nil
}

func (s *ReviewSvcSpy) ExceludeReview(id int) error {
	s.DeletedId = id
	return nil
}

func (s *ReviewSvcSpy) SearchById(id int) (dto.Review, error) {
	s.FoundId = id
	return dto.Review{Id: int64(id), AuthorName: "User1", Review: "Review1", Url_photo: "pic1.jpg"}, nil
}

func TestListAllReviews(t *testing.T) {
	spy := &ReviewSvcSpy{}
	controller := controllers.NewReviewController(spy)

	req := httptest.NewRequest(http.MethodGet, "/reviews", nil)
	rec := httptest.NewRecorder()

	controller.ListAllReviews(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestFindReviewById(t *testing.T) {
	spy := &ReviewSvcSpy{}
	controller := controllers.NewReviewController(spy)

	req := httptest.NewRequest(http.MethodGet, "/reviews/1", nil)
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	controller.FindReviewById(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if spy.FoundId != 1 {
		t.Fatalf("expected to search by id 1, got %d", spy.FoundId)
	}
}

func TestGetRandomReviews(t *testing.T) {
	spy := &ReviewSvcSpy{}
	controller := controllers.NewReviewController(spy)

	req := httptest.NewRequest(http.MethodGet, "/reviews/random", nil)
	rec := httptest.NewRecorder()

	controller.GetRandomReviews(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestCreateNewReview(t *testing.T) {
	spy := &ReviewSvcSpy{}
	controller := controllers.NewReviewController(spy)

	body := map[string]any{"author": "Valid Name", "review": "Valid Depoiment", "url_photo": "https://url.jpeg"}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/reviews", bytes.NewBuffer(payload))
	rec := httptest.NewRecorder()

	controller.CreateNewReview(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
	if spy.CreatedMap == nil {
		t.Fatalf("expected CreateReview to be called")
	}
}

func TestUpdateReview(t *testing.T) {
	spy := &ReviewSvcSpy{}
	controller := controllers.NewReviewController(spy)

	body := map[string]any{"author": "Updated Name", "review": "Updated Depoiment", "url_photo": "https://url.jpeg"}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPut, "/reviews/1", bytes.NewBuffer(payload))
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	controller.UpdateReview(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
	if spy.UpdatedMap == nil {
		t.Fatalf("expected UpdateReview to be called")
	}
}

func TestExceludeReview(t *testing.T) {
	spy := &ReviewSvcSpy{}
	controller := controllers.NewReviewController(spy)

	req := httptest.NewRequest(http.MethodDelete, "/reviews/1", nil)
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	controller.ExceludeReview(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
	if spy.DeletedId != 1 {
		t.Fatalf("expected to delete id 1, got %d", spy.DeletedId)
	}
}
