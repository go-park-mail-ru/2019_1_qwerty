package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	models "../models"
)

func init() {
	models.Scores = []models.Score{}
}

const frameSize = 10

func GetNextAfter(w http.ResponseWriter, r *http.Request) {
	offset := 0
	if val, err := strconv.Atoi(r.FormValue("offset")); err == nil {
		offset = val
	}

	w.Header().Set("Content-Type", "application/json")
	if offset < 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var answ []models.Score
	if offset+frameSize < len(models.Scores) {
		answ = models.Scores[offset : offset+frameSize]
	} else if offset < len(models.Scores) {
		answ = models.Scores[offset:]
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(answ)
	w.WriteHeader(http.StatusOK)
}

func CreateScore(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	points := 0
	if val, err := strconv.Atoi(r.FormValue("points")); err == nil {
		points = val
	}
	models.Scores = append(models.Scores, models.Score{uint64(rand.Intn(100)), name, uint64(points)})
	w.WriteHeader(http.StatusCreated)
}
