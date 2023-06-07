package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jonathansoufer/go-postgres-api/models"
)
func Get(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := models.Get(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}