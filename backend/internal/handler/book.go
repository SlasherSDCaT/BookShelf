package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"vk/internal/models"
	"vk/internal/models/DTO"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.service.Book.GetAll()
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	booksByte, err := json.Marshal(books)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("GetAllBooks is ok")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(booksByte)
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["id"]

	id, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	book, err := h.service.Book.Get(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	bookByte, err := json.Marshal(book)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("GetBook with id - " + strconv.Itoa(id))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bookByte)
}

func (h *Handler) SaveBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bookDTO DTO.BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()

	if err := validate.Struct(bookDTO); err != nil {
		errs := err.(validator.ValidationErrors)
		newErrorResponse(w, http.StatusBadRequest, errs.Error())
		return
	}

	// Преобразование DTO в модель
	book := models.Book{
		UserId:      bookDTO.UserId,
		Title:       bookDTO.Title,
		Author:      bookDTO.Author,
		Genre:       bookDTO.Genre,
		Description: bookDTO.Description,
		Image:       bookDTO.Image,
		Body:        bookDTO.Body,
	}

	id, err := h.service.Book.Create(book)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("SaveBook with id " + strconv.Itoa(id))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"id": %d}`, id)))
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input DTO.BookUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Book.Update(id, input); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("UpdateBook with id - " + strconv.Itoa(id))
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(statusResponse{"ok"})
	_, _ = w.Write(result)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Book.Delete(id); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("DeleteBook with id - " + strconv.Itoa(id))
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(statusResponse{"ok"})
	_, _ = w.Write(result)
}

func (h *Handler) FindBooksByGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	genre := r.URL.Query().Get("genre")
	if genre == "" {
		newErrorResponse(w, http.StatusBadRequest, "genre query parameter is required")
		return
	}

	books, err := h.service.Book.FindByGenre(genre)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	booksByte, err := json.Marshal(books)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("FindBooksByGenre is ok")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(booksByte)
}
