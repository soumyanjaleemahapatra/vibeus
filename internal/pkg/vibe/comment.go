package vibe

import (
	"time"
)

type Comment struct {
	ID         int64     `json:"id,primary_key"`
	Body       string    `json:"body"`
	UserHandle string    `json:"user_handle"`
	CreatedAt  time.Time `json:"created_at"`
	Read       bool      `json:"read"`
	VibeId     int64
}
