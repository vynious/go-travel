// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: connections.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createConnection = `-- name: CreateConnection :one
INSERT INTO connections (
                         party_a,
                         party_b
) VALUES (
          $1, $2
         ) RETURNING party_a, party_b, connected_date
`

type CreateConnectionParams struct {
	PartyA string `json:"party_a"`
	PartyB string `json:"party_b"`
}

func (q *Queries) CreateConnection(ctx context.Context, arg CreateConnectionParams) (Connection, error) {
	row := q.queryRow(ctx, q.createConnectionStmt, createConnection, arg.PartyA, arg.PartyB)
	var i Connection
	err := row.Scan(&i.PartyA, &i.PartyB, &i.ConnectedDate)
	return i, err
}

const deleteConnectionByUserId = `-- name: DeleteConnectionByUserId :one
DELETE FROM connections
WHERE (party_a = $1 AND party_b = $2)
   OR (party_a = $2 AND party_b = $1)
RETURNING party_a, party_b, connected_date
`

type DeleteConnectionByUserIdParams struct {
	PartyA string `json:"party_a"`
	PartyB string `json:"party_b"`
}

func (q *Queries) DeleteConnectionByUserId(ctx context.Context, arg DeleteConnectionByUserIdParams) (Connection, error) {
	row := q.queryRow(ctx, q.deleteConnectionByUserIdStmt, deleteConnectionByUserId, arg.PartyA, arg.PartyB)
	var i Connection
	err := row.Scan(&i.PartyA, &i.PartyB, &i.ConnectedDate)
	return i, err
}

const getConnectionsByUserId = `-- name: GetConnectionsByUserId :many
SELECT
    CASE
        WHEN party_a = $1 THEN party_b
        ELSE party_a
        END AS connected_user_id,
    connected_date
FROM connections
WHERE party_a = $1 OR party_b = $1
`

type GetConnectionsByUserIdRow struct {
	ConnectedUserID sql.NullString `json:"connected_user_id"`
	ConnectedDate   time.Time      `json:"connected_date"`
}

func (q *Queries) GetConnectionsByUserId(ctx context.Context, partyA string) ([]GetConnectionsByUserIdRow, error) {
	rows, err := q.query(ctx, q.getConnectionsByUserIdStmt, getConnectionsByUserId, partyA)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetConnectionsByUserIdRow{}
	for rows.Next() {
		var i GetConnectionsByUserIdRow
		if err := rows.Scan(&i.ConnectedUserID, &i.ConnectedDate); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
