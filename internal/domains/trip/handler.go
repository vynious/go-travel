package trip

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"net/http"
	"strconv"
)

type Handler struct {
	*Service
}

func NewTripHandler(s *Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) StartTrip(w http.ResponseWriter, r *http.Request) {
	var userReq StartTripRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	title := userReq.Title
	country := userReq.Country
	startDate := userReq.StartDate
	endDate := userReq.EndDate

	trip, err := h.CreateNewTrip(r.Context(), title, country, startDate, endDate)
	if err != nil {
		http.Error(w, "failed to create trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := TripDetailResponse{
		Trip: trip,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ViewTripDetails(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid id params", http.StatusInternalServerError)
		return
	}

	trip, err := h.GetTripById(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get trip details", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := TripDetailResponse{
		Trip: trip,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) ChangeTripDetails(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid id params", http.StatusInternalServerError)
		return
	}

	var userReq UpdateTripDetailRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tx, err := h.repository.DB.BeginTx(r.Context(), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedTrip db.Trip
	var updated bool

	if userReq.Title != nil {
		title := *userReq.Title
		updated = true
		trip, err := h.UpdateTripTitle(r.Context(), id, title)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedTrip = trip
	}

	if userReq.Country != nil {
		country := *userReq.Country
		updated = true
		trip, err := h.UpdateTripCountry(r.Context(), id, country)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedTrip = trip
	}

	if userReq.StartDate != nil {
		startDate := *userReq.StartDate
		updated = true
		trip, err := h.UpdateTripStartDate(r.Context(), id, startDate)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedTrip = trip
	}

	if userReq.EndDate != nil {
		endDate := *userReq.EndDate
		updated = true
		trip, err := h.UpdateTripEndDate(r.Context(), id, endDate)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedTrip = trip
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

	response := TripDetailResponse{
		Trip: updatedTrip,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
