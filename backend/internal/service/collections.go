package service

import (
	"vk/internal/models"
	"vk/internal/models/DTO"
)

type CollectionService struct {
	repo Collection
}

func NewCollectionService(repo Collection) *CollectionService {
	return &CollectionService{repo: repo}
}

func (s *CollectionService) GetAll() ([]DTO.CollectionDTO, error) {
	return s.repo.GetAll()
}

func (s *CollectionService) Get(id int) (DTO.CollectionWithBooksDTO, error) {
	return s.repo.Get(id)
}

func (s *CollectionService) Create(collection models.Collection) (int, error) {
	return s.repo.Create(collection)
}

func (s *CollectionService) Update(id int, input DTO.CollectionUpdate) error {
	return s.repo.Update(id, input)
}

func (s *CollectionService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CollectionService) AddBook(collectionId int, bookId int) error {
	return s.repo.AddBook(collectionId, bookId)
}

func (s *CollectionService) RemoveBook(collectionId int, bookId int) error {
	return s.repo.RemoveBook(collectionId, bookId)
}

func (s *CollectionService) RateCollection(id int, rating int) error {
	return s.repo.RateCollection(id, rating)
}
func (s *CollectionService) GetAllByUserId(userId int) ([]DTO.CollectionDTO, error){
	return s.repo.GetAllByUserId(userId)
}
