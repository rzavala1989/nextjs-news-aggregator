package auth

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"news-aggregator/internal/db"
	"news-aggregator/internal/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginHandler invoked")

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("Error decoding login request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	dbUser, err := db.GetUserByUsername(u.Username)
	if err != nil {
		log.Println("Error fetching user by username:", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password))
	if err != nil {
		log.Println("Password mismatch for user:", u.Username)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(u.Username)
	if err != nil {
		log.Println("Error generating JWT:", err)
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	log.Println("User logged in successfully:", u.Username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":       dbUser.ID,
			"username": dbUser.Username,
		},
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RegisterHandler invoked")

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("Error decoding register request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	u.Password = string(hashedPassword)

	err = db.CreateUser(&u)
	if err != nil {
		log.Println("Error registering user:", err)
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	log.Println("User registered successfully:", u.Username)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
