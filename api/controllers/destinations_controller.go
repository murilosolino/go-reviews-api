package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/murilosolino/challenge-backend-7/api/apperrors"
	"github.com/murilosolino/challenge-backend-7/api/helper"
	"github.com/murilosolino/challenge-backend-7/api/model"
)

type IDestinationSvc interface {
	CreateDestination(m map[string]interface{}) error
	ListDestinations() ([]model.DestinationRow, error)
	DeleteDestinationById(id int) error
	UpdateDestination(id int, m map[string]interface{}) error
	FindByName(name string) (model.DestinationRow, error)
}

type DestinationController struct {
	svc IDestinationSvc
}

func NewDestinationController(svc IDestinationSvc) *DestinationController {
	return &DestinationController{svc: svc}
}

func (c DestinationController) CreateNewDestination(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		slog.Error("[DestinationController][CreateNewDestination()]"+apperrors.APP_ERR_BODY_DECODE, "error", err)
		helper.ToJson(w, http.StatusUnprocessableEntity, apperrors.APP_ERR_BODY_DECODE, nil)
		return
	}

	err = c.svc.CreateDestination(m)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}
	helper.ToJson(w, http.StatusCreated, "", nil)
}

func (c DestinationController) ListAllDestinations(w http.ResponseWriter, r *http.Request) {
	destName := r.URL.Query().Get("name")
	if destName != "" {
		data, err := c.svc.FindByName(destName)
		if err != nil {
			helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
			return
		}

		helper.ToJson(w, http.StatusOK, "ok", data)
		return
	}

	data, err := c.svc.ListDestinations()
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}

	helper.ToJson(w, http.StatusOK, "ok", data)
}

func (c DestinationController) DeleteDestination(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("[DestinationController][DeleteDestination()]", "error", err)
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}

	err = c.svc.DeleteDestinationById(id)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_DELETE_REGISTER, nil)
		return
	}

	helper.ToJson(w, http.StatusNoContent, "", nil)
}

func (c DestinationController) UpdateDestination(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("[DestinationController][DeleteDestination()]", "error", err)
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}

	json.NewDecoder(r.Body).Decode(&m)
	err = c.svc.UpdateDestination(id, m)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_DELETE_REGISTER, nil)
		return
	}

	helper.ToJson(w, http.StatusNoContent, "", nil)
}
