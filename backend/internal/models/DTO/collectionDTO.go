package DTO

type CollectionDTO struct {
    CollectionId int    `json:"collection_id"`
    UserId       int    `json:"user_id"`
    Name         string `json:"name"`
    Description  string `json:"description,omitempty"`
    Rating       int    `json:"rating"`
    Public       bool   `json:"public"` 
}

type CollectionInput struct {
    UserId      int    `json:"user_id" validate:"required"`
    Name        string `json:"name" validate:"required,min=1,max=100"`
    Description string `json:"description,omitempty"`
    Rating      int    `json:"rating" validate:"min=0,max=5"`
    Public      bool   `json:"public"`  
}

type CollectionUpdate struct {
    UserId      *int    `json:"user_id,omitempty"`
    Name        *string `json:"name,omitempty" validate:"omitempty,min=1,max=100"`
    Description *string `json:"description,omitempty"`
    Rating      *int    `json:"rating,omitempty" validate:"omitempty,min=0,max=5"`
    Public      *bool   `json:"public,omitempty"`  
}

type CollectionWithBooksDTO struct {
    CollectionId int         `json:"collection_id"`
    UserId       int         `json:"user_id"`
    Name         string      `json:"name"`
    Description  string      `json:"description,omitempty"`
    Rating       int         `json:"rating"`
    Public       bool        `json:"public"`  
    Books        []BookDTO   `json:"books"`
}
