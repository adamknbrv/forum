package entity

import "time"

type Message struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	DiscussionID int       `json:"discussion_id"`
	Content      string    `json:"content"`
	CreateAt     time.Time `json:"create_at"`
}
