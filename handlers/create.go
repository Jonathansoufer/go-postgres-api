package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jonathansoufer/go-postgres-api/models"
)

func Create(w http.ResponseWriter, r *http.Request){
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := models.Insert(todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}