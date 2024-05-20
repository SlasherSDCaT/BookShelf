package handler

import (
	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() *mux.Router {
	r := mux.NewRouter()
	r.Use(Cors);

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sing-up", h.SingUp).Methods("POST", "OPTIONS")
	auth.HandleFunc("/sing-in", h.SingIn).Methods("POST", "OPTIONS")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(h.UserIdentify)

	books := api.PathPrefix("/books").Subrouter()
	books.HandleFunc("", h.GetAllBooks).Methods("GET", "OPTIONS")
	books.HandleFunc("", h.SaveBook).Methods("POST", "OPTIONS")
	books.HandleFunc("/{id}", h.UpdateBook).Methods("PATCH", "OPTIONS")
	books.HandleFunc("/{id}", h.DeleteBook).Methods("DELETE", "OPTIONS")
	books.HandleFunc("/search", h.FindBooksByGenre).Methods("GET", "OPTIONS")
	books.HandleFunc("/{id}", h.GetBook).Methods("GET", "OPTIONS")

	collections := api.PathPrefix("/collections").Subrouter()
	collections.HandleFunc("", h.GetAllCollections).Methods("GET", "OPTIONS")
	collections.HandleFunc("", h.SaveCollection).Methods("POST", "OPTIONS")
	collections.HandleFunc("/{id}", h.UpdateCollection).Methods("PATCH", "OPTIONS")
	collections.HandleFunc("/{id}", h.DeleteCollection).Methods("DELETE", "OPTIONS")
	collections.HandleFunc("/{id}", h.GetCollection).Methods("GET", "OPTIONS")
	collections.HandleFunc("/add-book/{id_collection}/{id_book}", h.AddBookToCollection).Methods("POST", "OPTIONS")
	collections.HandleFunc("/remove-book/{id_collection}/{id_book}", h.RemoveBookFromCollection).Methods("POST", "OPTIONS")
	collections.HandleFunc("/{id}/rate", h.RateCollection).Methods("POST", "OPTIONS")
	collections.HandleFunc("/user/{user_id}", h.GetCollectionsByUserId).Methods("GET", "OPTIONS") // Новый маршрут

	comments := api.PathPrefix("/comments").Subrouter()
	comments.HandleFunc("", h.SaveComment).Methods("POST", "OPTIONS")
	comments.HandleFunc("/{id}", h.DeleteComment).Methods("DELETE", "OPTIONS")
	comments.HandleFunc("/{bookId}", h.GetCommentsByBook).Methods("GET", "OPTIONS")

	return r
}
