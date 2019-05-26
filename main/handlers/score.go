package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"2019_1_qwerty/models"
)

func init() {
	models.Scores = []models.Score{}
}

const frameSize = 10

//GetNextAfter - get 10 players from scoreboard
func GetNextAfter(w http.ResponseWriter, r *http.Request) {
	offset := 0

	if val, err := strconv.Atoi(r.FormValue("offset")); err == nil {
		offset = val
	}

	w.Header().Set("Content-Type", "application/json")

	if offset < 0 {
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}

	var answ []models.Score

	if offset+frameSize < len(models.Scores) {
		answ = models.Scores[offset : offset+frameSize]
	} else if offset < len(models.Scores) {
		answ = models.Scores[offset:]
	} else {
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(answ)
	ErrorMux(&w, r, http.StatusOK)
}

//CreateScore - add player to scoreboard
func CreateScore(w http.ResponseWriter, r *http.Request) {
	name := "test49" // name = getUsernameByCookie()
	points := 0

	if val, err := strconv.Atoi(r.FormValue("points")); err == nil {
		points = val
	}

	models.Scores = append(models.Scores, models.Score{Place: uint64(rand.Intn(100)), Name: name, Points: uint64(points)})
	ErrorMux(&w, r, http.StatusCreated)
}
