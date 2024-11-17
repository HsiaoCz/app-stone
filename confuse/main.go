package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Users struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func WriteJson(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.UserID = uuid.New().String()
	WriteJson(w, http.StatusOK, map[string]any{
		"message": "create user success",
		"status":  http.StatusOK,
		"user":    user,
	})
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("POST /user", HandleCreateUser)
	http.ListenAndServe(":3001", router)
}
