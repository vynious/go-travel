package travel_entry

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"net/http"
	"strconv"
)

type TravelEntryHandler struct {
	*TravelEntryService
}

func NewTravelEntryHandler(s *TravelEntryService) *TravelEntryHandler {
	return &TravelEntryHandler{
		s,
	}
}

func (h *TravelEntryHandler) EnterTravelEntry(w http.ResponseWriter, r *http.Request) {
	var userReq NewTravelEntryRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	uid := userReq.UserId
	tid := userReq.TripId
	location := userReq.Location
	description := userReq.Description

	entry, err := h.CreateNewTravelEntry(r.Context(), uid, tid, location, description)
	if err != nil {
		http.Error(w, "failed to create travel entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := TravelEntryDetailResponse{
		TravelEntry: entry,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *TravelEntryHandler) ViewTravelEntry(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid id params", http.StatusInternalServerError)
		return
	}
	entry, err := h.GetTravelEntryById(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get travel entry details", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := TravelEntryDetailResponse{
		TravelEntry: entry,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *TravelEntryHandler) ViewTravelEntriesUnderTrip(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid id params", http.StatusInternalServerError)
		return
	}
	entries, err := h.GetTravelEntriesByTripId(r.Context(), id)

	w.WriteHeader(http.StatusOK)
	response := TravelEntriesDetailResponse{
		TravelEntries: entries,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

}

func (h *TravelEntryHandler) ViewTravelEntriesUnderTripAndUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	uid := query.Get("uid")
	strTID := query.Get("tid")

	tid, err := strconv.ParseInt(strTID, 10, 64)
	if err != nil {
		http.Error(w, "invalid tid params", http.StatusInternalServerError)
		return
	}

	entries, err := h.GetTravelEntriesByUserIdAndTripId(r.Context(), uid, tid)
	if err != nil {
		http.Error(w, "failed to get travel entries", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := TravelEntriesDetailResponse{
		TravelEntries: entries,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

}

func (h *TravelEntryHandler) UpdateTravelEntry(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid id params", http.StatusInternalServerError)
		return
	}

	var userReq UpdateTravelEntryRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tx, err := h.repository.DB.BeginTx(r.Context(), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedEntry db.TravelEntry
	var updated bool

	if userReq.Location != nil {
		location := *userReq.Location
		updated = true
		entry, err := h.UpdateTravelEntryLocationById(r.Context(), id, location)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedEntry = entry
	}

	if userReq.Description != nil {
		description := *userReq.Description
		updated = true
		entry, err := h.UpdateTravelEntryDescriptionById(r.Context(), id, description)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedEntry = entry
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

	response := TravelEntryDetailResponse{
		TravelEntry: updatedEntry,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *TravelEntryHandler) DeleteTravelEntry(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "invalid id params", http.StatusInternalServerError)
		return
	}

	entry, err := h.DeleteTravelEntryById(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get delete travel entry", http.StatusNotFound)
		return
	}

	response := TravelEntryDetailResponse{
		TravelEntry: entry,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
