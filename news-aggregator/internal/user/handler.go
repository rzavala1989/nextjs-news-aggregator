package user

import (
	"encoding/json"
	"log"
	"net/http"
	"news-aggregator/internal/db"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UserHandler invoked")
	username := r.Context().Value("username").(string)
	user, err := db.GetUserByUsername(username)
	if err != nil {
		log.Println("Error fetching user details:", err)
		http.Error(w, "Error fetching user details", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
