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
	log.Println("here1")
	offset := 0

	if val, err := strconv.Atoi(r.FormValue("offset")); err == nil {
		log.Println("here2")
		offset = val
		log.Println("here3", offset)
	}

	w.Header().Set("Content-Type", "application/json")

	if offset < 0 {
		log.Println("here4", offset)
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}
	log.Println("here5")
	rows, err := database.Database.Query(sqlGetScore, 10, offset)
	log.Println("here6")
	scores := models.Scores{}
	i := 1

	if err != nil {
		log.Println("here7", err)
	}

	log.Println("here8")

	for rows.Next() {
		score := models.Score{}
		log.Println("here9")
		score.Place = uint64(offset + i)
		log.Println(score.Place)
		_ = rows.Scan(&score.Name, &score.Points)
		i++

		scores = append(scores, score)
		log.Println(score.Place, score.Name, score.Points)
	}

	log.Println("here10", scores)

	result, _ := scores.MarshalJSON()

	log.Println("here11", result)
	w.Write(result)
	ErrorMux(&w, r, http.StatusOK)
}
