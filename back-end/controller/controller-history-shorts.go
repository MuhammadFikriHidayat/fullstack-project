package controller

import (
	"encoding/json"
	"net/http"

	"fullstack-project/back-end/config"
	"fullstack-project/back-end/model"
)

func ReceiveShort(w http.ResponseWriter, r *http.Request) {

	var history model.History

	err := json.NewDecoder(r.Body).Decode(&history)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config.DB.Create(&history)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(history)
}

func GetHistory(w http.ResponseWriter, r *http.Request) {

	var histories []model.History

	config.DB.Find(&histories)

	json.NewEncoder(w).Encode(histories)
}

func GetHistoryByID(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var history model.History

	config.DB.First(&history, id)

	json.NewEncoder(w).Encode(history)
}

