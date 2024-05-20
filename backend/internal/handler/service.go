package handler

import (
	"vk/internal/models"
	"vk/internal/models/DTO"
	"vk/internal/service"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, int, error)
	ParseToken(token string) (int, error)
	GetUser(id int) (models.User, error)
}


type Book interface {
	GetAll() ([]DTO.BookDTO, error)
	Get(id int) (DTO.BookDTO, error)
	Create(book models.Book) (int, error)
	Update(id int, input DTO.BookUpdate) error
	Delete(id int) error
	FindByGenre(genre string) ([]DTO.BookDTO, error)
}

type Collection interface {
	GetAll() ([]DTO.CollectionDTO, error)
	Get(id int) (DTO.CollectionWithBooksDTO, error)
	Create(collection models.Collection) (int, error)
	Update(id int, input DTO.CollectionUpdate) error
	Delete(id int) error
	AddBook(collectionId int, bookId int) error
	RemoveBook(collectionId int, bookId int) error
	RateCollection(id int, rating int) error
	GetAllByUserId(userId int) ([]DTO.CollectionDTO, error)
}

type Comment interface {
	Create(comment models.Comment) (int, error)
	Delete(id int) error
	FindByBook(bookId int) ([]DTO.CommentDTO, error)
}

type Service struct {
	Authorization
	Book
	Collection
	Comment
}

func NewService(repo *service.Repository) *Service {
	return &Service{
		Authorization: service.NewAuthService(repo.Authorization),
		Book:         service.NewBookService(repo.Book),
		Collection:         service.NewCollectionService(repo.Collection),
		Comment:         service.NewCommentService(repo.Comment),
	}
}
