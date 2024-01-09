package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	db "github.com/vynious/go-travel/internal/db/sqlc"
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
	var userReq RegisterUserRequest
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
	response := UserDetailResponse{
		User: user,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var userReq LoginUserRequest
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

	w.WriteHeader(http.StatusOK)
	response := UserDetailResponse{
		User: user,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) SearchUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	username := query.Get("username")
	email := query.Get("email")

	var user db.User
	var err error

	switch {
	case username != "":
		user, err = h.repository.queries.GetUserByUsername(r.Context(), username)
	case email != "":
		user, err = h.repository.queries.GetUserByEmail(r.Context(), email)
	default:
		http.Error(w, "No search parameters provided", http.StatusBadRequest)
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := UserDetailResponse{
		User: user,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ChangeUserProfilePicture(w http.ResponseWriter, r *http.Request) {
	/*
	 todo: communicate with s3 to get url, then update the profile_picture field with the new url
	*/
	var url string // replace
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.UpdateUserPictureById(r.Context(), id, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := UserDetailResponse{
		User: user,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ChangeUserDetails(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	var userReq UpdateUserDetailRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tx, err := h.repository.db.BeginTx(r.Context(), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedUser db.User
	var updated bool

	if userReq.Name != nil {
		name := *userReq.Name
		updated = true
		user, err := h.UpdateUserNameById(r.Context(), id, name)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedUser = user
	}

	if userReq.Username != nil {
		username := *userReq.Username
		updated = true
		user, err := h.UpdateUserUsernameById(r.Context(), id, username)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedUser = user
	}

	if userReq.Email != nil {
		email := *userReq.Email
		updated = true
		user, err := h.UpdateUserEmailById(r.Context(), id, email)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedUser = user
	}

	if updated {
		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
	} else {
		tx.Rollback()
		http.Error(w, "no updates to perform", http.StatusBadRequest)
		return
	}

	response := UserDetailResponse{
		User: updatedUser,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.DeleteUserById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := UserDetailResponse{
		User: user,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
