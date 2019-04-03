package api

import (
	"encoding/json"
	"net/http"
	"os"

	"2019_1_qwerty/models"
)

func Version(w http.ResponseWriter, r *http.Request) {
	status := models.Status{
		Commit: os.Getenv("COMMIT"),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
