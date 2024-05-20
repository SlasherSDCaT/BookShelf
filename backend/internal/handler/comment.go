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

// GetCommentsByBook handles retrieving comments for a specific book
func (h *Handler) GetCommentsByBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["bookId"]
	bookId, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	comments, err := h.service.Comment.FindByBook(bookId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	commentsByte, err := json.Marshal(comments)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("GetCommentsByBook for bookId - %d", bookId)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(commentsByte)
}

// SaveComment handles creating a new comment
func (h *Handler) SaveComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var commentDTO DTO.CommentDTO
	if err := json.NewDecoder(r.Body).Decode(&commentDTO); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	validate := validator.New()
	if err := validate.Struct(commentDTO); err != nil {
		errs := err.(validator.ValidationErrors)
		newErrorResponse(w, http.StatusBadRequest, errs.Error())
		return
	}

	comment := models.Comment{
		UserId: commentDTO.UserId,
		BookId: commentDTO.BookId,
		Rating: commentDTO.Rating,
		Text:   commentDTO.Text,
	}

	id, err := h.service.Comment.Create(comment)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("SaveComment with id %d", id)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"id": %d}`, id)))
}

// DeleteComment handles deleting a comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	if err := h.service.Comment.Delete(id); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("DeleteComment with id - %d", id)
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(statusResponse{"ok"})
	_, _ = w.Write(result)
}
