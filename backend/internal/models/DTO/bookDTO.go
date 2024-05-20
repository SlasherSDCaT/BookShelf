package DTO

// BookDTO represents the data transfer object for a book
type BookDTO struct {
	BookId      int    `json:"book_id"`
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
	Body        string `json:"body,omitempty"`
}

// BookUpdate represents the data transfer object for updating a book
type BookUpdate struct {
	UserId      *int    `json:"user_id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Author      *string `json:"author,omitempty"`
	Genre       *string `json:"genre,omitempty"`
	Description *string `json:"description,omitempty"`
	Image *string `json:"image,omitempty"`
	Body        *string `json:"body,omitempty"`
}
