package models

type User struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Score    uint64 `json:"score"`
	Avatar   string `json:"avatar"`
}

//easyjson:json
var Users map[string]User
