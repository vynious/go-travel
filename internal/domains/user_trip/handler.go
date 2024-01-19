package user_trip

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type UserTripHandler struct {
	*UserTripService
}

func NewUserTripHandler(s *UserTripService) *UserTripHandler {
	return &UserTripHandler{
		s,
	}
}

func (h *UserTripHandler) AddUserToTrip(w http.ResponseWriter, r *http.Request) {
	strTID := chi.URLParam(r, "tripId")
	tid, err := strconv.ParseInt(strTID, 10, 64)
	if err != nil {
		http.Error(w, "invalid tid params", http.StatusInternalServerError)
		return
	}

	var userReq AddUserTripRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	uid := userReq.UserId

	ut, err := h.CreateNewUserTrip(r.Context(), uid, tid)
	if err != nil {
		http.Error(w, "failed to create usertrip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := UserTripDetailResponse{
		UserTrip: ut,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *UserTripHandler) GetAllTripsForUserId(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "userId")
	uts, err := h.GetUserTripByUserId(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get user trips", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := UserTripsDetailsResponse{
		UserTrips: uts,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *UserTripHandler) GetAllUsersOnTripId(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "tripId")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid id params", http.StatusInternalServerError)
		return
	}

	uts, err := h.GetUserTripByTripId(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get user trips", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := UserTripsDetailsResponse{
		UserTrips: uts,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *UserTripHandler) DeleteUserFromTripId(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "userId")
	strTID := chi.URLParam(r, "tripId")

	tid, err := strconv.ParseInt(strTID, 10, 64)
	if err != nil {
		http.Error(w, "invalid tid params", http.StatusInternalServerError)
		return
	}

	ut, err := h.DeleteUserTripById(r.Context(), uid, tid)
	if err != nil {
		http.Error(w, "failed to delete user trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := UserTripDetailResponse{
		UserTrip: ut,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
