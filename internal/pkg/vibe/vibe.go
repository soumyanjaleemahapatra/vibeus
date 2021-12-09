package vibe

import (
	"time"
)

type Vibe struct {
	ID         int64     `json:"id" gorm:"primary_key;auto_increment"`
	Body       string    `json:"body"`
	UserHandle string    `json:"user_handle"`
	CreatedAt  time.Time `json:"created_at"`
	LikeCount  int       `json:"like_count"`
	Comments   []Comment `json:"comments"`
}
