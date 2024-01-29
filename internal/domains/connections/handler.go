package connections

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ConnectionHandler struct {
	*ConnectionService
}

func NewConnectionHandler(s *ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{
		s,
	}
}

func (h *ConnectionHandler) MakeConnection(w http.ResponseWriter, r *http.Request) {
	a := chi.URLParam(r, "a")
	b := chi.URLParam(r, "b")

	conn, err := h.CreateConnection(r.Context(), a, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ConnectionResponse{
		PartyA:  conn.PartyA,
		PartyB:  conn.PartyB,
		Message: "successfully made connection",
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *ConnectionHandler) ViewConnection(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	conns, err := h.GetUserConnections(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := ConnectionsResponse{
		Connections: conns,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *ConnectionHandler) RemoveConnection(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	targetId := r.URL.Query().Get("targetId")

	conn, err := h.DeleteConnection(r.Context(), userId, targetId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ConnectionResponse{
		PartyA:  conn.PartyA,
		PartyB:  conn.PartyB,
		Message: "successfully removed connection",
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
