package user

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type Handler struct {
	*Service
}

func NewUserHandler(s *Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userReq RegisterUserReq
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	name := userReq.User.Name
	email := userReq.User.Email
	username := userReq.User.Username

	/*
		password := userReq.Password
		todo: send email and password to firebase for registration
	*/

	user, err := h.CreateNewUser(r.Context(), name, username, email)
	if err != nil {
		// handle error, write error response
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := RegisterUserRes{
		User: user,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var userReq LoginUserRes
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	email := userReq.Email
	password := userReq.Password
	exist, err := h.VerifyUserExistence(r.Context(), email, password)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if exist {

	} else {

	}
	/*

		todo: send email and password to firebase for verification

	*/

	return
}

func (h *Handler) ViewUserDetails(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := h.GetUserById(r.Context(), id)

}

func (h *Handler) SearchUser(w http.ResponseWriter, r *http.Request) {

}
