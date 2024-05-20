package service

import (
	"database/sql"
	"vk/internal/models"
	"vk/internal/models/DTO"
	"vk/internal/repository"
)

type Authorization interface {
	Create(user models.User) (int, error)
	GetOne(username, password string) (models.User, error)
	GetOneById(id int) (models.User, error)
}

type Book interface {
	FindAll() ([]DTO.BookDTO, error)
	FindOne(id int) (DTO.BookDTO, error)
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

type Repository struct {
	Authorization
	Book
	Collection
	Comment
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: repository.NewAuthPostgres(db),
		Book:         repository.NewBookPostgres(db),
		Collection:         repository.NewCollectionPostgres(db),
		Comment:         repository.NewCommentPostgres(db),
	}
}
