package models

type Message struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Processed bool   `json:"processed"`
}
