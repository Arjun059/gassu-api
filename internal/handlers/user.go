package handlers

import (
	"encoding/json"
	"fmt"
	"gassu/internal/models"
	utils "gassu/internal/utils"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Created")
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Get")
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Updated")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Deleted")
}

func (h *UserHandler) SignupUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var body models.User

	// Decode the request body into 'body'
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Body Not Parsed", http.StatusBadRequest)
		return
	}

	// Check if a user with the same email already exists
	if err := h.DB.Where("email = ?", body.Email).First(&user).Error; err == nil {
		http.Error(w, "User Already Exists", http.StatusUnprocessableEntity)
		return
	}

	// Create a new user since no user was found
	if err := h.DB.Create(&body).Error; err != nil {
		log.Printf("Decoded body: %+v\n", body)
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Encode the newly created user in the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

func (h *UserHandler) SigninUser(w http.ResponseWriter, r *http.Request) {

	var body models.User
	json.NewDecoder(r.Body).Decode(&body)

	var user *models.User

	log.Printf("Decoded body: %+v\n", body)

	if e := h.DB.Where("email = ?", body.Email).First(&user).Error; e != nil {
		log.Printf("Decoded body: %+v\n", e)

		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	if user.Email != body.Email || user.Password != body.Password {
		http.Error(w, "Un Auth", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.CreateToken(user.Email, user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	type Response struct {
		Token   string `json:"token"`
		Success bool   `json:"success"`
		Error   bool   `json:"error"`
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Token: tokenString, Success: true, Error: false})

}
