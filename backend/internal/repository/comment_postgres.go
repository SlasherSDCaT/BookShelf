package repository

import (
	"database/sql"
	"fmt"
	"log"
	"vk/internal/db"
	"vk/internal/models"
	"vk/internal/models/DTO"
)

type CommentPostgres struct {
	db *sql.DB
}

func NewCommentPostgres(db *sql.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

// Create adds a new comment to the database
func (r *CommentPostgres) Create(comment models.Comment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, book_id, rating, text) VALUES ($1, $2, $3, $4) RETURNING comment_id", db.COMMENTS)
	row := r.db.QueryRow(query, comment.UserId, comment.BookId, comment.Rating, comment.Text)
	if err := row.Scan(&id); err != nil {
		log.Panic(err)
		return 0, err
	}
	return id, nil
}

// Delete removes a comment from the database
func (r *CommentPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE comment_id=$1", db.COMMENTS)
	if _, err := r.db.Exec(query, id); err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

// FindByBook retrieves comments from the database by book ID
func (r *CommentPostgres) FindByBook(bookId int) ([]DTO.CommentDTO, error) {
	query := fmt.Sprintf("SELECT comment_id, user_id, book_id, rating, text FROM %s WHERE book_id=$1", db.COMMENTS)
	rows, err := r.db.Query(query, bookId)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()

	commentsDTO := make([]DTO.CommentDTO, 0)
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.CommentId, &comment.UserId, &comment.BookId, &comment.Rating, &comment.Text); err != nil {
			log.Fatal(err)
			return nil, err
		}
		commentDTO := DTO.CommentDTO{
			CommentId: comment.CommentId,
			UserId:    comment.UserId,
			BookId:    comment.BookId,
			Rating:    comment.Rating,
			Text:      comment.Text,
		}
		commentsDTO = append(commentsDTO, commentDTO)
	}
	return commentsDTO, nil
}
