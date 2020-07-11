package routes

import (
	"github.com/Bareksa/rest-api-boilerplate/controllers"
	"github.com/gorilla/mux"
)

type Route struct {}

func (r *Route) Init() *mux.Router{
	healthCheckController := controllers.InitHealthCheckController()
	exampleController := controllers.InitExampleController()
	route := mux.NewRouter().StrictSlash(false)

	v1 := route.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/ping", healthCheckController.Ping).Methods("GET")
	v1.HandleFunc("/users", exampleController.GetUsers).Methods("GET")

	return v1
}