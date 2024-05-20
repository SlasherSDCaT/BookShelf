package service

import (
	"vk/internal/models"
	"vk/internal/models/DTO"
)

type BookService struct {
	repo Book
}

func NewBookService(repo Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAll() ([]DTO.BookDTO, error) {
	return s.repo.FindAll()
}

func (s *BookService) Get(id int) (DTO.BookDTO, error) {
	return s.repo.FindOne(id)
}

func (s *BookService) Create(book models.Book) (int, error) {
	return s.repo.Create(book)
}

func (s *BookService) Update(id int, input DTO.BookUpdate) error {
	return s.repo.Update(id, input)
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *BookService) FindByGenre(genre string) ([]DTO.BookDTO, error) {
	return s.repo.FindByGenre(genre)
}
