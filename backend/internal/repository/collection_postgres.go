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

type CollectionPostgres struct {
	db *sql.DB
}

func NewCollectionPostgres(db *sql.DB) *CollectionPostgres {
	return &CollectionPostgres{db: db}
}

// Create adds a new collection to the database
func (r *CollectionPostgres) Create(collection models.Collection) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, name, description, rating, public) VALUES ($1, $2, $3, $4, $5) RETURNING collection_id", db.COLLECTIONS)
	row := r.db.QueryRow(query, collection.UserId, collection.Name, collection.Description, collection.Rating, collection.Public)
	if err := row.Scan(&id); err != nil {
		log.Panic(err)
		return 0, err
	}
	return id, nil
}

// Update modifies an existing collection in the database
func (r *CollectionPostgres) Update(id int, input DTO.CollectionUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *input.UserId)
		argId++
	}
	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *input.Rating)
		argId++
	}
	if input.Public != nil {
		setValues = append(setValues, fmt.Sprintf("public=$%d", argId))
		args = append(args, *input.Public)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE collection_id=$%d", db.COLLECTIONS, setQuery, argId)
	args = append(args, id)
	log.Println("query: " + query)
	if _, err := r.db.Exec(query, args...); err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

// Delete removes a collection from the database
func (r *CollectionPostgres) Delete(id int) error {
    // Удаление записей из таблицы связи CollectionBook
    deleteCollectionBooksQuery := fmt.Sprintf("DELETE FROM %s WHERE collection_id=$1", db.COLLECTION_BOOKS)
    if _, err := r.db.Exec(deleteCollectionBooksQuery, id); err != nil {
        log.Panic(err)
        return err
    }

    // Удаление записи из таблицы Collections
    deleteCollectionQuery := fmt.Sprintf("DELETE FROM %s WHERE collection_id=$1", db.COLLECTIONS)
    if _, err := r.db.Exec(deleteCollectionQuery, id); err != nil {
        log.Panic(err)
        return err
    }

    return nil
}

func (r *CollectionPostgres) GetAll() ([]DTO.CollectionDTO, error) {
    query := fmt.Sprintf("SELECT collection_id, user_id, name, description, rating, public FROM %s", db.COLLECTIONS)
    rows, err := r.db.Query(query)
    if err != nil {
        log.Panic(err)
        return nil, err
    }
    defer rows.Close()

    collectionsDTO := make([]DTO.CollectionDTO, 0)
    for rows.Next() {
        var collection models.Collection
        if err := rows.Scan(&collection.CollectionId, &collection.UserId, &collection.Name, &collection.Description, &collection.Rating, &collection.Public); err != nil {
            log.Fatal(err)
            return nil, err
        }
        collectionDTO := DTO.CollectionDTO{
            CollectionId: collection.CollectionId,
            UserId:       collection.UserId,
            Name:         collection.Name,
            Description:  collection.Description,
            Rating:       collection.Rating,
            Public:       collection.Public,
        }
        collectionsDTO = append(collectionsDTO, collectionDTO)
    }
    return collectionsDTO, nil
}


// FindOne retrieves a single collection from the database by ID
func (r *CollectionPostgres) Get(id int) (DTO.CollectionWithBooksDTO, error) {
	query := fmt.Sprintf("SELECT collection_id, user_id, name, description, rating, public FROM %s WHERE collection_id=$1", db.COLLECTIONS)
	var collection models.Collection
	if err := r.db.QueryRow(query, id).Scan(&collection.CollectionId, &collection.UserId, &collection.Name, &collection.Description, &collection.Rating, &collection.Public); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return DTO.CollectionWithBooksDTO{}, errors.New("empty result")
		} else {
			log.Panic(err)
			return DTO.CollectionWithBooksDTO{}, err
		}
	}

	booksQuery := fmt.Sprintf(`
		SELECT b.book_id, b.user_id, b.title, b.author, b.genre, b.description, b.image, b.body 
		FROM %s AS b
		INNER JOIN %s AS cb ON b.book_id = cb.book_id
		WHERE cb.collection_id=$1`, db.BOOKS, db.COLLECTION_BOOKS)

	rows, err := r.db.Query(booksQuery, id)
	if err != nil {
		log.Panic(err)
		return DTO.CollectionWithBooksDTO{}, err
	}
	defer rows.Close()

	var books []DTO.BookDTO
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookId, &book.UserId, &book.Title, &book.Author, &book.Genre, &book.Description, &book.Image, &book.Body); err != nil {
			log.Fatal(err)
			return DTO.CollectionWithBooksDTO{}, err
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
		books = append(books, bookDTO)
	}

	collectionDTO := DTO.CollectionWithBooksDTO{
		CollectionId: collection.CollectionId,
		UserId:       collection.UserId,
		Name:         collection.Name,
		Description:  collection.Description,
		Rating:       collection.Rating,
		Public:       collection.Public,
		Books:        books,
	}
	return collectionDTO, nil
}

func (r *CollectionPostgres) AddBook(collectionId int, bookId int) error {
	// Check if the collection exists
	var collectionExists bool
	collectionQuery := "SELECT EXISTS(SELECT 1 FROM collections WHERE collection_id=$1)"
	err := r.db.QueryRow(collectionQuery, collectionId).Scan(&collectionExists)
	if err != nil {
		log.Println("Error checking if collection exists:", err)
		return err
	}
	if !collectionExists {
		log.Printf("Collection with ID %d does not exist\n", collectionId)
		return fmt.Errorf("collection with ID %d does not exist", collectionId)
	}

	// Check if the book exists
	var bookExists bool
	bookQuery := "SELECT EXISTS(SELECT 1 FROM books WHERE book_id=$1)"
	err = r.db.QueryRow(bookQuery, bookId).Scan(&bookExists)
	if err != nil {
		log.Println("Error checking if book exists:", err)
		return err
	}
	if !bookExists {
		log.Printf("Book with ID %d does not exist\n", bookId)
		return fmt.Errorf("book with ID %d does not exist", bookId)
	}

	// Insert the record into collections_books
	insertQuery := fmt.Sprintf("INSERT INTO %s (collection_id, book_id) VALUES ($1, $2)", db.COLLECTION_BOOKS)
	if _, err := r.db.Exec(insertQuery, collectionId, bookId); err != nil {
		log.Println("Error inserting into collections_books:", err)
		return err
	}
	return nil
}

func (r *CollectionPostgres) RemoveBook(collectionId int, bookId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE collection_id=$1 AND book_id=$2", db.COLLECTION_BOOKS)
	if _, err := r.db.Exec(query, collectionId, bookId); err != nil {
		log.Println("Error deleting from collections_books:", err)
		return err
	}
	return nil
}

func (r *CollectionPostgres) RateCollection(collectionId int, rating int) error {
	// Добавить новый рейтинг в таблицу collection_ratings
	addRatingQuery := fmt.Sprintf("INSERT INTO %s (collection_id, rating) VALUES ($1, $2)", db.COLLECTION_RATINGS)
	if _, err := r.db.Exec(addRatingQuery, collectionId, rating); err != nil {
		log.Println("Error adding rating to collection_ratings:", err)
		return err
	}

	// Пересчитать средний рейтинг для коллекции
	updateAvgRatingQuery := fmt.Sprintf(`
        UPDATE %s
        SET rating = (
            SELECT AVG(rating)
            FROM %s
            WHERE collection_id = $1
        )
        WHERE collection_id = $1`, db.COLLECTIONS, db.COLLECTION_RATINGS)

	if _, err := r.db.Exec(updateAvgRatingQuery, collectionId); err != nil {
		log.Println("Error updating average rating in collections:", err)
		return err
	}

	return nil
}

// GetAllByUserId retrieves all collections from the database for a specific user
func (r *CollectionPostgres) GetAllByUserId(userId int) ([]DTO.CollectionDTO, error) {
	query := fmt.Sprintf("SELECT collection_id, user_id, name, description, rating, public FROM %s WHERE user_id=$1", db.COLLECTIONS)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()

	collectionsDTO := make([]DTO.CollectionDTO, 0)
	for rows.Next() {
		var collection models.Collection
		if err := rows.Scan(&collection.CollectionId, &collection.UserId, &collection.Name, &collection.Description, &collection.Rating, &collection.Public); err != nil {
			log.Fatal(err)
			return nil, err
		}
		collectionDTO := DTO.CollectionDTO{
			CollectionId: collection.CollectionId,
			UserId:       collection.UserId,
			Name:         collection.Name,
			Description:  collection.Description,
			Rating:       collection.Rating,
			Public:       collection.Public,
		}
		collectionsDTO = append(collectionsDTO, collectionDTO)
	}
	return collectionsDTO, nil
}
