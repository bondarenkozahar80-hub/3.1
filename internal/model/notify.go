package model

import "time"

type Notification struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	SendAt    int       `json:"send_at"`
	CreatedAt time.Time `json:"created_at"`
}
