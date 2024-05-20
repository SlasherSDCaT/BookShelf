package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"vk/internal/models"

	"github.com/go-playground/validator/v10"
)

type AuthResponse struct {
	Token  string
	UserID int
}

func (h *Handler) SingUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		newErrorResponse(w, http.StatusBadRequest, errs.Error())
		return
	}

	_, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, id, err := h.service.Authorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("SingIn is ok")
	w.WriteHeader(http.StatusOK)

	auth, err := json.Marshal(AuthResponse{
		Token:  token,
		UserID: id,
	})
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, _ = w.Write(auth)
}

type UserInType struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) SingIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user UserInType

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		newErrorResponse(w, http.StatusBadRequest, errs.Error())
		return
	}

	token, id, err := h.service.Authorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("SingIn is ok")
	w.WriteHeader(http.StatusOK)

	auth, err := json.Marshal(AuthResponse{
		Token:  token,
		UserID: id,
	})
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	_, _ = w.Write(auth)
}
