package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 메모리에 게시글을 저장할 슬라이스
var Posts []Post
var NextID uint = 1
