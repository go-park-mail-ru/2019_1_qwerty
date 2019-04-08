package models

type User struct {
	//ID uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
	Score    uint64 `json:"score"`
	Avatar   string `json:"avatar"`
}

var Users map[string]User
