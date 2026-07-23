package controller

import (
	"encoding/json"
	"net/http"

	"fullstack-project/back-end/config"
	"fullstack-project/back-end/model"
)

func CreateScrapJob(w http.ResponseWriter, r *http.Request) {

	var job model.Scrap

	json.NewDecoder(r.Body).Decode(&job)

	config.DB.Create(&job)

	json.NewEncoder(w).Encode(job)
}

func GetJobs(w http.ResponseWriter, r *http.Request) {

	var jobs []model.Scrap

	config.DB.Find(&jobs)

	json.NewEncoder(w).Encode(jobs)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var job model.Scrap

	config.DB.Delete(&job, id)

	w.WriteHeader(http.StatusOK)
}