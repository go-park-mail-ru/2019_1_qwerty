package handlers

import (
	"net/http"
	"strconv"

	"2019_1_qwerty/database"
	"2019_1_qwerty/models"
	"log"
)

const sqlGetScore = `
	SELECT player, score FROM scores
	ORDER BY score DESC
	LIMIT $1
	OFFSET $2
`

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

	rows, _ := database.Database.Query(sqlGetScore, 10, offset)
	scores := models.Scores{}
	i := 1

	for rows.Next() {
		score := models.Score{}
		score.Place = uint64(offset + i)
		_ = rows.Scan(&score.Name, &score.Points)
		i++

		scores = append(scores, score)
		log.Println(score.Place, score.Name, score.Points)
	}

	result, _ := scores.MarshalJSON()
	w.Write(result)
	ErrorMux(&w, r, http.StatusOK)
}
