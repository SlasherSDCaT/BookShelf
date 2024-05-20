package models

type Book struct {
	BookId      int    `json:"book_id"`
	UserId      int    `json:"user_id"`
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Author      string `json:"author" validate:"required,min=1,max=200"`
	Genre       string `json:"genre,omitempty"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
	Body        string `json:"body,omitempty"`
}