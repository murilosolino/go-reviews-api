package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/murilosolino/challenge-backend-7/internal/apperrors"
	"github.com/murilosolino/challenge-backend-7/internal/dto"
	"github.com/murilosolino/challenge-backend-7/internal/helper"
	"github.com/murilosolino/challenge-backend-7/internal/validations"
)

type IReviewSvc interface {
	GetAllReviews() ([]dto.Review, error)
	Get3RandomReviews() ([]dto.Review, error)
	CreateReview(rev map[string]interface{}) error
	UpdateReview(id int, r map[string]interface{}) error
	ExceludeReview(id int) error
	SearchById(id int) (dto.Review, error)
}

type ReviewsController struct {
	svc IReviewSvc
}

func NewReviewController(svc IReviewSvc) *ReviewsController {
	return &ReviewsController{svc: svc}
}

func (c *ReviewsController) ListAllReviews(w http.ResponseWriter, r *http.Request) {
	data, err := c.svc.GetAllReviews()
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}

	helper.ToJson(w, http.StatusOK, "ok", data)
}

func (c *ReviewsController) FindReviewById(w http.ResponseWriter, r *http.Request) {
	var review dto.Review
	idS := r.PathValue("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		slog.Error("[ConversionError][ReviewController][FindReviewById()] Não foi possível converter id para inteiro",
			"error", err)

		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}

	review, err = c.svc.SearchById(id)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return

	}

	helper.ToJson(w, http.StatusOK, "ok", review)
}

func (c *ReviewsController) GetRandomReviews(w http.ResponseWriter, r *http.Request) {

	results, err := c.svc.Get3RandomReviews()
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}

	helper.ToJson(w, http.StatusOK, "ok", results)
}

func (c *ReviewsController) CreateNewReview(w http.ResponseWriter, r *http.Request) {
	var rev map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&rev)
	if err != nil {
		helper.ToJson(w, http.StatusUnprocessableEntity, "Não foi possível realizar o parse do corpo da requisição", nil)
		return
	}

	jsonBytes, _ := json.Marshal(rev)
	var reviewDto dto.Review
	_ = json.Unmarshal(jsonBytes, &reviewDto)

	err = validations.ValidateReview(reviewDto)
	if err != nil {
		helper.ToJson(w, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	err = c.svc.CreateReview(rev)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}
	helper.ToJson(w, http.StatusCreated, "ok", nil)

}

func (c *ReviewsController) UpdateReview(w http.ResponseWriter, r *http.Request) {
	var review map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		helper.ToJson(w, http.StatusUnprocessableEntity, "Não foi possível realizar o parse do corpo da requisição", nil)
		return
	}

	jsonBytes, _ := json.Marshal(review)
	var reviewDto dto.Review
	_ = json.Unmarshal(jsonBytes, &reviewDto)

	err = validations.ValidateReview(reviewDto)
	if err != nil {
		helper.ToJson(w, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	idS := r.PathValue("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		slog.Error("[ConversionError][ReviewController][UpdateReview] Não foi possível converter id para inteiro",
			"error", err)
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return

	}

	err = c.svc.UpdateReview(id, review)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helper.ToJson(w, http.StatusNoContent, "ok", nil)
}

func (c *ReviewsController) ExceludeReview(w http.ResponseWriter, r *http.Request) {
	idS := r.PathValue("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		slog.Error("[ConversionError][ReviewController][UpdateReview] Não foi possível converter id para inteiro",
			"error", err)
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return

	}
	err = c.svc.ExceludeReview(id)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}
	helper.ToJson(w, http.StatusNoContent, "ok", nil)
}
