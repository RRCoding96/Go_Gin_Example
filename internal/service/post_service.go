package service

import (
	"myapp/internal/domain"
)

type PostService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) GetAllPosts() ([]domain.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) GetPost(id uint) (domain.Post, error) {
	return s.repo.GetByID(id)
}

func (s *PostService) CreatePost(post domain.Post) (domain.Post, error) {
	return s.repo.Create(post)
}

func (s *PostService) UpdatePost(post domain.Post) (domain.Post, error) {
	return s.repo.Update(post)
}

func (s *PostService) DeletePost(id uint) error {
	return s.repo.Delete(id)
}
