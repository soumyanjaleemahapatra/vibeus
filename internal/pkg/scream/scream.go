package scream

import "time"

type Scream struct {
	ID        int       `json:"id,primary_key"`
	Body      string    `json:"body"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LikeCount time.Time `json:"like_count"`
}
