package models

import (
	"encoding/json"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Completed   bool   `json:"completed"`
}

func (t Task) String() string {
	jsonData, _ := json.Marshal(t)
	return string(jsonData)
}
