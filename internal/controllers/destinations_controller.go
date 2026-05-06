package controllers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/murilosolino/challenge-backend-7/internal/apperrors"
	"github.com/murilosolino/challenge-backend-7/internal/dto"
	"github.com/murilosolino/challenge-backend-7/internal/helper"
	"github.com/murilosolino/challenge-backend-7/internal/validations"
)

type IDestinationSvc interface {
	CreateDestination(m map[string]interface{}) error
	ListDestinations() ([]dto.Destination, error)
	DeleteDestinationById(id int) error
	UpdateDestination(id int, m map[string]interface{}) error
	FindByName(name string) (dto.Destination, error)
	GenerateIADescriptiveText(destination string) string
}

type DestinationController struct {
	svc IDestinationSvc
}

func NewDestinationController(svc IDestinationSvc) *DestinationController {
	return &DestinationController{svc: svc}
}

func (c DestinationController) CreateNewDestination(w http.ResponseWriter, r *http.Request) {
	m := map[string]any{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		slog.Error("[DestinationController][CreateNewDestination()]"+apperrors.APP_ERR_BODY_DECODE, "error", err)
		helper.ToJson(w, http.StatusUnprocessableEntity, apperrors.APP_ERR_BODY_DECODE, nil)
		return
	}

	jsonBytes, _ := json.Marshal(m)
	var destination dto.Destination
	_ = json.Unmarshal(jsonBytes, &destination)

	err = validations.ValidateDestination(destination)
	if err != nil {
		helper.ToJson(w, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	if m["descriptive_text"] == nil {
		destination := fmt.Sprintf("%v", m["name"])
		m["descriptive_text"] = c.svc.GenerateIADescriptiveText(destination)
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
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		helper.ToJson(w, http.StatusUnprocessableEntity, apperrors.APP_ERR_BODY_DECODE, nil)
		return
	}

	jsonBytes, _ := json.Marshal(m)
	var destination dto.Destination
	_ = json.Unmarshal(jsonBytes, &destination)

	err = validations.ValidateDestination(destination)
	if err != nil {
		helper.ToJson(w, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("[DestinationController][DeleteDestination()]", "error", err)
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}

	err = c.svc.UpdateDestination(id, m)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_DELETE_REGISTER, nil)
		return
	}

	helper.ToJson(w, http.StatusNoContent, "", nil)
}
