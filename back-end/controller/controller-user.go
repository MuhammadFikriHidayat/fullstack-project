package controller

import (
	"encoding/json"
	"net/http"
	//"strconv"

	"fullstack-project/back-end/config"
	"fullstack-project/back-end/model"
)

// get all
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User

	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// get by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var user model.User

	config.DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}

// add user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var user model.User

	config.DB.First(&user, id)
	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var user model.User

	config.DB.Delete(&user, id)

	w.WriteHeader(http.StatusOK)
}
