package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"2019_1_qwerty/models"

	"github.com/prometheus/client_golang/prometheus"
)

var FooCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "foo_total",
	Help: "Number of foo successfully processed.",
})

var Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"status", "path"})

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
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var answ []models.Score

	if offset+frameSize < len(models.Scores) {
		answ = models.Scores[offset : offset+frameSize]
	} else if offset < len(models.Scores) {
		answ = models.Scores[offset:]
	} else {
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(answ)
	Hits.WithLabelValues(string(http.StatusOK), r.URL.String()).Inc()
	FooCount.Add(1)
	w.WriteHeader(http.StatusOK)
}

//CreateScore - add player to scoreboard
func CreateScore(w http.ResponseWriter, r *http.Request) {
	name := "test49" // name = getUsernameByCookie()
	points := 0

	if val, err := strconv.Atoi(r.FormValue("points")); err == nil {
		points = val
	}

	models.Scores = append(models.Scores, models.Score{Place: uint64(rand.Intn(100)), Name: name, Points: uint64(points)})
	Hits.WithLabelValues(string(http.StatusCreated), r.URL.String()).Inc()
	FooCount.Add(1)
	w.WriteHeader(http.StatusCreated)
}
