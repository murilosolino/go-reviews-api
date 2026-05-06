package controllers

import "net/http"

type HealthCheckController struct {
}

func NewHealthCheck() *HealthCheckController {
	return &HealthCheckController{}
}

func (c *HealthCheckController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
