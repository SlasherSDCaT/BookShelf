package service

import (
	"vk/internal/models"
	"vk/internal/models/DTO"
)

type CommentService struct {
	repo Comment
}

func NewCommentService(repo Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(comment models.Comment) (int, error) {
	return s.repo.Create(comment)
}
func (s *CommentService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CommentService) FindByBook(bookId int) ([]DTO.CommentDTO, error) {
	return s.repo.FindByBook(bookId)
}
