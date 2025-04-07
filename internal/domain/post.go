package domain

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" binding:"required" gorm:"size:255;not null"`
	Content   string    `json:"content" binding:"required" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

// PostRepository 인터페이스 정의
type PostRepository interface {
	GetAll() ([]Post, error)
	GetByID(id uint) (Post, error)
	Create(post Post) (Post, error)
	Update(post Post) (Post, error)
	Delete(id uint) error
}
