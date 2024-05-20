package models

type Comment struct {
	CommentId int    `json:"comment_id"`
	UserId    int    `json:"user_id"`
	BookId    int    `json:"book_id"`
	Rating    int    `json:"rating" validate:"min=0,max=5"`
	Text      string `json:"text,omitempty"`
}