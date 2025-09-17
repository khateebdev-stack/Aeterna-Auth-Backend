package handlers

import (
	"aeterna-auth/internal/models"
	"aeterna-auth/internal/services"
	"encoding/json"
	"net/http"
	"fmt"
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
		if err != nil{
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return



		}

		// we will add jwt logic here later
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Logged In, Welcome to dashboard %s", user.Email)})	}
}
