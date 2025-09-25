package handlers

import (
	"aeterna-auth/internal/models"
	"aeterna-auth/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterUser(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body

		var req models.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// 2. call the service to register the user
		if err := userService.RegisterUser(&req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 3. send a succesful response
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User Registered succesfully"})
	}
}

func LoginUser(userService *services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var creds models.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := userService.LoginUser(creds.Email, creds.Password)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return

		}

		// we will add jwt logic here later
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Logged In, Welcome to dashboard %s", user.Email)})
	}
}

func DeleteUser(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email string `json:"email"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload ", http.StatusBadRequest)
			return
		}

		if err := userService.DeleteUser(req.Email); err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User Delete successfully"})
	}
}

func UpdatePassword(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email       string `json:"email"`
			NewPassword string `json:"new_password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return

		}
		if req.Email == "" || req.NewPassword == "" {
			http.Error(w, "Email and new password are required", http.StatusBadRequest)
			return
		}

		if err := userService.UpdateUserPassword(req.Email, req.NewPassword); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Password updated"})
	}
}

func UpdateEmail(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			CurrentEmail string `json:"current_email"`
			NewEmail     string `json:"new_email"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payoad", http.StatusBadRequest)
			return

		}

		if err := userService.UpdateUserEmail(req.CurrentEmail, req.NewEmail); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Email Updated successfully"})
	}
}


func CheckEmail (userServices *services.UserService) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		
	}
}