package handlers

import (
	"encoding/json"
	"net/http"
)

type TestHandler struct{}

func (t *TestHandler) HandleTestConnect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"addr":    r.RemoteAddr,
		"agent":   r.UserAgent(),
		"message": "connect success",
	})
}
