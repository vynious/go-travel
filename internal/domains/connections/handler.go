package connections

import "net/http"

type ConnectionHandler struct {
	*ConnectionService
}

func NewConnectionHandler(s *ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{
		s,
	}
}

func (h *ConnectionHandler) MakeConnection(w http.ResponseWriter, r *http.Request) {

}

func (h *ConnectionHandler) ViewConnection(w http.ResponseWriter, r *http.Request) {

}

func (h *ConnectionHandler) RemoveConnection(w http.ResponseWriter, r *http.Request) {

}
