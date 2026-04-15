package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/murilosolino/challenge-backend-7/api/apperrors"
	"github.com/murilosolino/challenge-backend-7/api/helper"
	"github.com/murilosolino/challenge-backend-7/api/model"
)

type IDestinationSvc interface {
	CreateDestination(d model.DestinationRow) error
	ListDestinations() ([]model.DestinationRow, error)
}

type DestinationController struct {
	svc IDestinationSvc
}

func NewDestinationController(svc IDestinationSvc) *DestinationController {
	return &DestinationController{svc: svc}
}

func (c DestinationController) CreateNewDestination(w http.ResponseWriter, r *http.Request) {
	var d model.DestinationRow
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		slog.Error("[DestinationController][CreateNewDestination]"+apperrors.APP_ERR_BODY_DECODE, "error", err)
		helper.ToJson(w, http.StatusUnprocessableEntity, apperrors.APP_ERR_BODY_DECODE, nil)
		return
	}

	err = c.svc.CreateDestination(d)
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}
	helper.ToJson(w, http.StatusCreated, "", nil)
}

func (c DestinationController) ListAllDestinations(w http.ResponseWriter, r *http.Request) {
	data, err := c.svc.ListDestinations()
	if err != nil {
		helper.ToJson(w, http.StatusInternalServerError, apperrors.APP_ERR_UNPROCESSABLE_REQ, nil)
		return
	}

	helper.ToJson(w, http.StatusOK, "ok", data)
}
