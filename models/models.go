package models

import (
	"time"

	"github.com/idirall22/comment/models"
)

// Post structure
type Post struct {
	ID        int64             `json:"id"`
	Content   string            `json:"content"`
	MediaURLs []string          `json:"media_urls"`
	UserID    int64             `json:"user_id"`
	GroupID   int64             `json:"group_id"`
	Comments  []*models.Comment `json:"comments"`
	CreatedAt time.Time         `json:"created_at"`
	DeletedAt *time.Time        `json:"deleted_at"`
}

// ClientStream structure
type ClientStream struct {
	Post    chan *Post
	UserID  int64
	GroupID int64
}
