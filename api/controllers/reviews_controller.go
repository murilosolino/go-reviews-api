package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/murilosolino/challenge-backend-7/api/model"
	"github.com/murilosolino/challenge-backend-7/api/services"
)

type ReviewsController struct {
	svc services.ReviewService
}

func NewReviewController(svc services.ReviewService) *ReviewsController {
	return &ReviewsController{svc: svc}
}

func (c *ReviewsController) ListAllReviews(w http.ResponseWriter, r *http.Request) {
	data, err := c.svc.GetAllReviews()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&data)
	if err != nil {
		slog.Error("[JSON-ENCODER:ERROR][ReviewsController][ListAllReviews()] houve um erro ao realizar o encode dos dados", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
}

func (c *ReviewsController) FindReviewById(w http.ResponseWriter, r *http.Request) {
	var review model.ReviewsRow
	idS := r.PathValue("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		slog.Error("[ConversionError][ReviewController][FindReviewById()] Não foi possível converter id para inteiro",
			"error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	review, err = c.svc.SearchById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	json.NewEncoder(w).Encode(&review)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
}

func (c *ReviewsController) CreateNewReview(w http.ResponseWriter, r *http.Request) {
	var rev model.ReviewsRow
	json.NewDecoder(r.Body).Decode(&rev)

	err := c.svc.CreateReview(rev)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (c *ReviewsController) UpdateReview(w http.ResponseWriter, r *http.Request) {
	var review map[string]interface{}
	idS := r.PathValue("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		slog.Error("[ConversionError][ReviewController][UpdateReview] Não foi possível converter id para inteiro",
			"error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	json.NewDecoder(r.Body).Decode(&review)
	err = c.svc.UpdateReview(id, review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *ReviewsController) ExceludeReview(w http.ResponseWriter, r *http.Request) {
	idS := r.PathValue("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		slog.Error("[ConversionError][ReviewController][UpdateReview] Não foi possível converter id para inteiro",
			"error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	err = c.svc.ExceludeReview(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
