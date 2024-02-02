package travel_entry

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"github.com/vynious/go-travel/internal/domains/media"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"
)

type TravelEntryHandler struct {
	*TravelEntryService
	mediaService *media.MediaService
}

func NewTravelEntryHandler(s *TravelEntryService, m *media.MediaService) *TravelEntryHandler {
	return &TravelEntryHandler{
		s,
		m,
	}
}

func (h *TravelEntryHandler) EnterTravelEntry(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form to get file and fields
	if err := r.ParseMultipartForm(32 << 20); err != nil { // 32MB max memory
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Retrieve other form values
	uid := r.FormValue("UserId")
	strTID := r.FormValue("TripId")
	location := r.FormValue("Location")
	description := r.FormValue("Description")

	tid, err := strconv.ParseInt(strTID, 10, 64)
	if err != nil {
		http.Error(w, "invalid id params", http.StatusInternalServerError)
		return
	}

	// Create the travel entry in the database
	entry, err := h.CreateNewTravelEntry(r.Context(), uid, tid, location, description)
	if err != nil {
		http.Error(w, "failed to create travel entry", http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["media"]
	var wg sync.WaitGroup
	result := make([]*media.MediaResponse, len(files))
	errCh := make(chan error, len(files))

	// loop to create new media
	for i, fileHeader := range files {
		wg.Add(1)
		go func(fh *multipart.FileHeader, idx int) {
			defer wg.Done()

			response, err := h.mediaService.CreateNewMedia(r.Context(), entry.ID, strconv.Itoa(idx))
			if err != nil {
				errCh <- fmt.Errorf(
					"failed to create new media :%w", err)
				return
			}
			result[idx] = response
			errCh <- nil
		}(fileHeader, i)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for i := 0; i < len(files); i++ {
		if err := <-errCh; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Send the response back with x amount of pre-signed urls
	w.WriteHeader(http.StatusCreated)
	response := TravelEntryDetailWithMediaResponse{
		TravelEntry: entry,
		SignedMedia: result,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *TravelEntryHandler) ViewTravelEntry(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "entryId")
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
	data, err := h.mediaService.GetMediasByEntryId(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get signed media urls for travel entry", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := TravelEntryDetailWithMediaResponse{
		TravelEntry: entry,
		SignedMedia: data,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *TravelEntryHandler) ViewTravelEntriesUnderTrip(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "tripId")
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
	uid := chi.URLParam(r, "userId")
	strTID := chi.URLParam(r, "tripId")

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

	// todo:  how to update images in travel entry??

	strId := chi.URLParam(r, "entryId")
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
			if err := tx.Rollback(); err != nil {
				log.Print("rollback error: %w", err)

			}
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
			if err := tx.Rollback(); err != nil {
				log.Print("rollback error: %w", err)
			}
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
		if err := tx.Rollback(); err != nil {
			log.Print("rollback error: %w", err)
		}
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
	strId := chi.URLParam(r, "entryId")
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

	signedUrls, err := h.mediaService.DeleteMediaByEntryId(r.Context(), entry.ID)
	if err != nil {
		http.Error(w, "failed to get signed media urls for travel entry", http.StatusNotFound)
		return
	}

	response := TravelEntryDetailWithMediaResponse{
		TravelEntry: entry,
		SignedMedia: signedUrls,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
