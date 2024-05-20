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

func (h *Handler) GetAllCollections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collections, err := h.service.Collection.GetAll()
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	collectionsByte, err := json.Marshal(collections)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("GetAllCollections is ok")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(collectionsByte)
}

func (h *Handler) GetCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["id"]

	id, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	collection, err := h.service.Collection.Get(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	collectionByte, err := json.Marshal(collection)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("GetCollection with id - " + strconv.Itoa(id))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(collectionByte)
}

func (h *Handler) SaveCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collectionDTO DTO.CollectionInput
	if err := json.NewDecoder(r.Body).Decode(&collectionDTO); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()

	if err := validate.Struct(collectionDTO); err != nil {
		errs := err.(validator.ValidationErrors)
		newErrorResponse(w, http.StatusBadRequest, errs.Error())
		return
	}

	// Преобразование DTO в модель
	collection := models.Collection{
		UserId:      collectionDTO.UserId,
		Name:        collectionDTO.Name,
		Description: collectionDTO.Description,
		Rating:      collectionDTO.Rating,
		Public:      collectionDTO.Public,
	}

	id, err := h.service.Collection.Create(collection)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("SaveCollection with id " + strconv.Itoa(id))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"id": %d}`, id)))
}

func (h *Handler) UpdateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input DTO.CollectionUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Collection.Update(id, input); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("UpdateCollection with id - " + strconv.Itoa(id))
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(statusResponse{"ok"})
	_, _ = w.Write(result)
}

func (h *Handler) DeleteCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Collection.Delete(id); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("DeleteCollection with id - " + strconv.Itoa(id))
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(statusResponse{"ok"})
	_, _ = w.Write(result)
}

func (h *Handler) AddBookToCollection(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Получаем параметры из URL
    vars := mux.Vars(r)
    collectionId, err := strconv.Atoi(vars["id_collection"])
    if err != nil {
        newErrorResponse(w, http.StatusBadRequest, "invalid collection ID")
        return
    }
    bookId, err := strconv.Atoi(vars["id_book"])
    if err != nil {
        newErrorResponse(w, http.StatusBadRequest, "invalid book ID")
        return
    }

    // Вызываем сервис для добавления книги в коллекцию
    if err := h.service.Collection.AddBook(collectionId, bookId); err != nil {
        newErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    log.Printf("AddBookToCollection with collection_id - %d and book_id - %d", collectionId, bookId)
    w.WriteHeader(http.StatusOK)
    result, _ := json.Marshal(statusResponse{"ok"})
    _, _ = w.Write(result)
}

func (h *Handler) RemoveBookFromCollection(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Получаем параметры из URL
    vars := mux.Vars(r)
    collectionId, err := strconv.Atoi(vars["id_collection"])
    if err != nil {
        newErrorResponse(w, http.StatusBadRequest, "invalid collection ID")
        return
    }
    bookId, err := strconv.Atoi(vars["id_book"])
    if err != nil {
        newErrorResponse(w, http.StatusBadRequest, "invalid book ID")
        return
    }

    // Вызываем сервис для удаления книги из коллекции
    if err := h.service.Collection.RemoveBook(collectionId, bookId); err != nil {
        newErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    log.Printf("RemoveBookFromCollection with collection_id - %d and book_id - %d", collectionId, bookId)
    w.WriteHeader(http.StatusOK)
    result, _ := json.Marshal(statusResponse{"ok"})
    _, _ = w.Write(result)
}


func (h *Handler) RateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		Rating int `json:"rating" validate:"required,min=0,max=5"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Collection.RateCollection(id, input.Rating); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("RateCollection with id - " + strconv.Itoa(id) + " and rating - " + strconv.Itoa(input.Rating))
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(statusResponse{"ok"})
	_, _ = w.Write(result)
}

func (h *Handler) GetCollectionsByUserId(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId, err := strconv.Atoi(vars["user_id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    collections, err := h.service.Collection.GetAllByUserId(userId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if len(collections) == 0 {
        http.Error(w, "No collections found for the user", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(collections)
}

