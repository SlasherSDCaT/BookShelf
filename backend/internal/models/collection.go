package models

type Collection struct {
	CollectionId int    `json:"collection_id"`
	UserId       int    `json:"user_id"`
	Name         string `json:"name" validate:"required,min=1,max=100"`
	Description  string `json:"description,omitempty"`
	Rating       int    `json:"rating" validate:"min=0,max=5"`
	Public       bool   `json:"public"`
}

type CollectionBook struct {
	CollectionId int `json:"collection_id"`
	BookId       int `json:"book_id"`
}
