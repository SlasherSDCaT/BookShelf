package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"vk/internal/db"
	"vk/internal/models"
	"vk/internal/models/DTO"
)

type BookPostgres struct {
	db *sql.DB
}

func NewBookPostgres(db *sql.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

// Create adds a new book to the database
func (r *BookPostgres) Create(book models.Book) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, title, author, genre, description, image, body) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING book_id", db.BOOKS)
	row := r.db.QueryRow(query, book.UserId, book.Title, book.Author, book.Genre, book.Description, book.Image, book.Body)
	if err := row.Scan(&id); err != nil {
		log.Panic(err)
		return 0, err
	}
	return id, nil
}

// Update modifies an existing book in the database
func (r *BookPostgres) Update(id int, input DTO.BookUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *input.UserId)
		argId++
	}
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, *input.Author)
		argId++
	}
	if input.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre=$%d", argId))
		args = append(args, *input.Genre)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("image=$%d", argId))
		args = append(args, *input.Image)
		argId++
	}
	if input.Body != nil {
		setValues = append(setValues, fmt.Sprintf("body=$%d", argId))
		args = append(args, *input.Body)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE book_id=$%d", db.BOOKS, setQuery, argId)
	args = append(args, id)
	log.Println("query: " + query)
	if _, err := r.db.Exec(query, args...); err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

// Delete removes a book from the database
func (r *BookPostgres) Delete(id int) error {
    // Удаление всех комментариев, связанных с книгой
    deleteCommentsQuery := fmt.Sprintf("DELETE FROM %s WHERE book_id=$1", db.COMMENTS)
    if _, err := r.db.Exec(deleteCommentsQuery, id); err != nil {
        log.Panic(err)
        return err
    }

    // Удаление записи из таблицы Books
    deleteBookQuery := fmt.Sprintf("DELETE FROM %s WHERE book_id=$1", db.BOOKS)
    if _, err := r.db.Exec(deleteBookQuery, id); err != nil {
        log.Panic(err)
        return err
    }

    return nil
}

// FindAll retrieves all books from the database
func (r *BookPostgres) FindAll() ([]DTO.BookDTO, error) {
	query := fmt.Sprintf("SELECT book_id, user_id, title, author, genre, description,image, body FROM %s", db.BOOKS)
	rows, err := r.db.Query(query)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()

	booksDTO := make([]DTO.BookDTO, 0)
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookId, &book.UserId, &book.Title, &book.Author, &book.Genre, &book.Description, &book.Image, &book.Body); err != nil {
			log.Fatal(err)
			return nil, err
		}
		bookDTO := DTO.BookDTO{
			BookId:      book.BookId,
			UserId:      book.UserId,
			Title:       book.Title,
			Author:      book.Author,
			Genre:       book.Genre,
			Description: book.Description,
			Image:       book.Image,
			Body:        book.Body,
		}
		booksDTO = append(booksDTO, bookDTO)
	}
	return booksDTO, nil
}

// FindOne retrieves a single book from the database by ID
func (r *BookPostgres) FindOne(id int) (DTO.BookDTO, error) {
	query := fmt.Sprintf("SELECT book_id, user_id, title, author, genre, description, image, body FROM %s WHERE book_id=$1", db.BOOKS)
	var book models.Book
	if err := r.db.QueryRow(query, id).Scan(&book.BookId, &book.UserId, &book.Title, &book.Author, &book.Genre, &book.Description, &book.Image, &book.Body); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return DTO.BookDTO{}, errors.New("empty result")
		} else {
			log.Panic(err)
			return DTO.BookDTO{}, err
		}
	}
	bookDTO := DTO.BookDTO{
		BookId:      book.BookId,
		UserId:      book.UserId,
		Title:       book.Title,
		Author:      book.Author,
		Genre:       book.Genre,
		Description: book.Description,
		Image:       book.Image,
		Body:        book.Body,
	}
	return bookDTO, nil
}

// FindByGenre retrieves books from the database by genre
func (r *BookPostgres) FindByGenre(genre string) ([]DTO.BookDTO, error) {
	query := fmt.Sprintf("SELECT book_id, user_id, title, author, genre, description, image, body FROM %s WHERE genre=$1", db.BOOKS)
	rows, err := r.db.Query(query, genre)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()

	booksDTO := make([]DTO.BookDTO, 0)
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookId, &book.UserId, &book.Title, &book.Author, &book.Genre, &book.Description, &book.Image, &book.Body); err != nil {
			log.Fatal(err)
			return nil, err
		}
		bookDTO := DTO.BookDTO{
			BookId:      book.BookId,
			UserId:      book.UserId,
			Title:       book.Title,
			Author:      book.Author,
			Genre:       book.Genre,
			Description: book.Description,
			Image:       book.Image,
			Body:        book.Body,
		}
		booksDTO = append(booksDTO, bookDTO)
	}
	return booksDTO, nil
}
