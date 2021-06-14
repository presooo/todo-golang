package controller

import (
	"encoding/json"
	"net/http"
	"todo/src/models"
	"todo/src/service"

	"github.com/gorilla/mux"
)

func GetTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	task := service.GetTask(params["id"])

	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Tasks
	_ = json.NewDecoder(r.Body).Decode(&task)

	taskId := service.CreateTask(task)

	json.NewEncoder(w).Encode(taskId)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	result := service.DeleteTask(params["id"])

	if result > 0 {
		json.NewEncoder(w).Encode("Successfully deleted the task with id: " + params["id"])
	} else {
		json.NewEncoder(w).Encode("A task with id ' " + params["id"] + "' not found...")
	}
}
