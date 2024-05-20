package DTO

type CommentDTO struct {
	CommentId int    `json:"comment_id"`
	UserId    int    `json:"user_id"`
	BookId    int    `json:"book_id"`
	Rating    int    `json:"rating"`
	Text      string `json:"text,omitempty"`
}

type CommentUpdate struct {
	UserId *int    `json:"user_id,omitempty"`
	BookId *int    `json:"book_id,omitempty"`
	Rating *int    `json:"rating,omitempty" validate:"min=0,max=5"`
	Text   *string `json:"text,omitempty"`
}
