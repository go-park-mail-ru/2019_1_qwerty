package api

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type ScoreTable struct {
	Place string
	Name  string
	Score string
}

var testScore = []ScoreTable{
	ScoreTable{"1", "user1", "999"},
	ScoreTable{"2", "user2", "998"},
	ScoreTable{"3", "user3", "997"},
	ScoreTable{"4", "user4", "996"},
	ScoreTable{"5", "user5", "995"},
}

func GetNextAfter(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	//index := ctx.UserValue("index").(string)

	//score_json, _ := json.Marshal(testScore)
	//userVar2fmt.Println(string(score_json))
	json.NewEncoder(ctx).Encode(testScore)
}
