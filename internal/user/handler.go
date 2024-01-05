package user

import (
	"net/http"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) (*RegisterUserRes, error) {

	return nil, nil
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) (*LoginUserRes, error) {

	return nil, nil
}
