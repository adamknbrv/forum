package entity

import "time"

type Discussion struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	UserID   int       `json:"user_id"`
	CreateAt time.Time `json:"create_at"`
}
