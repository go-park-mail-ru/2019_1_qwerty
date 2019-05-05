package models

import (
	"encoding/json"
)

//Logs - websocket logs
type Logs struct {
	Head    string      `json:"head"`
	Content interface{} `json:"content,omitempty"`
}

//IncomeLogs - incoming logs
type IncomeLogs struct {
	Head    string          `json:"head"`
	Content json.RawMessage `json:"content,omitempty"`
}
