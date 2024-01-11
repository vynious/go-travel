package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"github.com/vynious/go-travel/internal/domains/auth"
	"net/http"
)

type Handler struct {
	*Service
	firebaseClient *auth.Client
}

func NewUserHandler(s *Service, fba *auth.Client) *Handler {
	return &Handler{
		s,
		fba,
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
		send email and password to firebase for registration
	*/
	password := userReq.Password
	fuser, err := h.firebaseClient.CreateNewUser(r.Context(), name, email, password)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}
	uid := fuser.UID

	user, err := h.CreateNewUser(r.Context(), uid, name, username, email)
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
	id := chi.URLParam(r, "id")

	user, err := h.GetUserById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

func (h *Handler) ViewAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.GetAllUser(r.Context())
	if err != nil {
		http.Error(w, "failed to get all users", http.StatusInternalServerError)
		return
	}

	response := AllUserDetailResponse{
		Users: users,
	}

	w.WriteHeader(http.StatusOK)
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
		user, err = h.repository.Queries.GetUserByUsername(r.Context(), username)
	case email != "":
		user, err = h.repository.Queries.GetUserByEmail(r.Context(), email)
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
	id := chi.URLParam(r, "id")

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
	id := chi.URLParam(r, "id")

	var userReq UpdateUserDetailRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tx, err := h.repository.DB.BeginTx(r.Context(), nil)
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
	id := chi.URLParam(r, "id")

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
