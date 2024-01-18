package travel_entry

import "net/http"

type TravelEntryHandler struct {
	*TravelEntryService
}

func NewTravelEntryHandler(s *TravelEntryService) *TravelEntryHandler {
	return &TravelEntryHandler{
		s,
	}
}

func (h *TravelEntryHandler) EnterTravelEntry(w http.ResponseWriter, r *http.Request) {

}

func (h *TravelEntryHandler) ViewTravelEntry(w http.ResponseWriter, r *http.Request) {

}

func (h *TravelEntryHandler) ViewTravelEntriesUnderTrip(w http.ResponseWriter, r *http.Request) {

}

func (h *TravelEntryHandler) ViewAllTravelEntriesUnderTrip(w http.ResponseWriter, r *http.Request) {

}

func (h *TravelEntryHandler) ViewTravelEntriesUnderTripAndUser(w http.ResponseWriter, r *http.Request) {

}

func (h *TravelEntryHandler) UpdateTravelEntry(w http.ResponseWriter, r *http.Request) {

}

func (h *TravelEntryHandler) DeleteTravelEntry(w http.ResponseWriter, r *http.Request) {

}
