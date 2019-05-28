package models

type Score struct {
	// ID     uint64 `json:"id"`
	Place  uint64 `json:"Place"`
	Name   string `json:"Name"`
	Points uint64 `json:"Score"`
}

//easyjson:json
var Scores []Score
