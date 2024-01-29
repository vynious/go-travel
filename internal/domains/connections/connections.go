package connections

import db "github.com/vynious/go-travel/internal/db/sqlc"

type ConnectionResponse struct {
	PartyA  string
	PartyB  string
	Message string
}

type ConnectionsResponse struct {
	Connections []db.GetConnectionsByUserIdRow
}
