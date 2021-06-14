package router

import (
	"todo/src/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", controller.GetTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task/{id}", controller.DeleteTask).Methods("DELETE", "OPTIONS")

	return router
}
