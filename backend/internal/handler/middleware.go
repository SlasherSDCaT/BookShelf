package handler

import (
	"context"
	//"github.com/gorilla/context"
	"log"
	"net/http"
	"strings"
)

const (
	USERROLE = "userRole"
)

func (h *Handler) UserIdentify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		if header == "" {
			newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}
		headerPart := strings.Split(header, " ")
		if len(headerPart) != 2 {
			newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, err := h.service.Authorization.ParseToken(headerPart[1])
		if err != nil {
			newErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}
		user, err := h.service.Authorization.GetUser(userId)
		if err != nil {
			newErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		log.Println("User with role - " + user.Role)
		ctx := context.WithValue(context.Background(), "userRole", user.Role)
		r.WithContext(ctx)
		//context.Set(r, USERROLE, user.Role)
		next.ServeHTTP(w, r)
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://bookshelf-1-ph0o.onrender.com")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
