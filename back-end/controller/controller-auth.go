package controller

import (
	"encoding/json"
	"net/http"

	"fullstack-project/back-end/config"
	"fullstack-project/back-end/model"
	"fullstack-project/back-end/validator"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// config.DB.Create(&user)

	// Validasi email
	if err := validator.ValidateEmail(user.Email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi password
	if err := validator.ValidatePassword(user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi role
	if err := validator.ValidateRole(user.Role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Cek email sudah terdaftar
	var existing model.User
	if err := config.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		http.Error(w, "email already registered", http.StatusConflict)
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)

	// Simpan ke database
	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	// Jangan kirim password kembali ke client
	user.Password = ""

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {

	// JWT akan dibuat pada middleware
}