package controllers

import (
	"github.com/Bareksa/rest-api-boilerplate/utils"
	"net/http"
)

func InitHealthCheckController() *HealthCheckController{
	return &HealthCheckController{}
}

type HealthCheckController struct {
}

func (c *HealthCheckController) Ping(w http.ResponseWriter, r *http.Request)  {
	utils.Response(w, http.StatusOK, "ok")
	return
}
