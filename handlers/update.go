package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Jonathansoufer/go-postgres-api/models"
)
func Update(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := models.Update(int64(id), todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Too many rows were update by mistake, %d", rows)
	} else if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}