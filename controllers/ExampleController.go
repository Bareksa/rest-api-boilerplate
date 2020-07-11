package controllers

import (
	"github.com/Bareksa/rest-api-boilerplate/services"
	"github.com/Bareksa/rest-api-boilerplate/utils"
	"net/http"
)

type ExampleController struct {
	Service services.IExampleService
}

func InitExampleController() *ExampleController{
	return &ExampleController{
		Service: services.InitExampleService(),
	}
}

func (c *ExampleController) GetUsers(w http.ResponseWriter, r *http.Request)  {
	users := c.Service.GetUser()
	utils.Response(w, http.StatusOK, users)
	return
}
