package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jonathansoufer/go-postgres-api/models"
)

func List(w http.ResponseWriter, r *http.Request){
	todos, err := models.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
	
}