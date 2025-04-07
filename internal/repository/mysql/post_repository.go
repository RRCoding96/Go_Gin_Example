package mysql

import (
	"myapp/internal/domain"
	"time"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) GetAll() ([]domain.Post, error) {
	var posts []domain.Post
	result := r.db.Find(&posts)
	return posts, result.Error
}

func (r *postRepository) GetByID(id uint) (domain.Post, error) {
	var post domain.Post
	result := r.db.First(&post, id)
	return post, result.Error
}

func (r *postRepository) Create(post domain.Post) (domain.Post, error) {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	result := r.db.Create(&post)
	return post, result.Error
}

func (r *postRepository) Update(post domain.Post) (domain.Post, error) {
	post.UpdatedAt = time.Now()
	result := r.db.Save(&post)
	return post, result.Error
}

func (r *postRepository) Delete(id uint) error {
	result := r.db.Delete(&domain.Post{}, id)
	return result.Error
}
