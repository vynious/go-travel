// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCommentStmt, err = db.PrepareContext(ctx, createComment); err != nil {
		return nil, fmt.Errorf("error preparing query CreateComment: %w", err)
	}
	if q.createConnectionStmt, err = db.PrepareContext(ctx, createConnection); err != nil {
		return nil, fmt.Errorf("error preparing query CreateConnection: %w", err)
	}
	if q.createMediaStmt, err = db.PrepareContext(ctx, createMedia); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMedia: %w", err)
	}
	if q.createTravelEntryStmt, err = db.PrepareContext(ctx, createTravelEntry); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTravelEntry: %w", err)
	}
	if q.createTripStmt, err = db.PrepareContext(ctx, createTrip); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTrip: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.createUserTripStmt, err = db.PrepareContext(ctx, createUserTrip); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUserTrip: %w", err)
	}
	if q.deleteCommentByIdStmt, err = db.PrepareContext(ctx, deleteCommentById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCommentById: %w", err)
	}
	if q.deleteConnectionByUserIdStmt, err = db.PrepareContext(ctx, deleteConnectionByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteConnectionByUserId: %w", err)
	}
	if q.deleteMediaByIdStmt, err = db.PrepareContext(ctx, deleteMediaById); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMediaById: %w", err)
	}
	if q.deleteTravelEntryStmt, err = db.PrepareContext(ctx, deleteTravelEntry); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTravelEntry: %w", err)
	}
	if q.deleteTripStmt, err = db.PrepareContext(ctx, deleteTrip); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTrip: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.deleteUserTripStmt, err = db.PrepareContext(ctx, deleteUserTrip); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserTrip: %w", err)
	}
	if q.getAllMediaByEntryIdStmt, err = db.PrepareContext(ctx, getAllMediaByEntryId); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllMediaByEntryId: %w", err)
	}
	if q.getAllMediaByTripIdStmt, err = db.PrepareContext(ctx, getAllMediaByTripId); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllMediaByTripId: %w", err)
	}
	if q.getAllMediaByTripIdAndUserIdStmt, err = db.PrepareContext(ctx, getAllMediaByTripIdAndUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllMediaByTripIdAndUserId: %w", err)
	}
	if q.getAllMediaByUserIdStmt, err = db.PrepareContext(ctx, getAllMediaByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllMediaByUserId: %w", err)
	}
	if q.getAllTravelEntryStmt, err = db.PrepareContext(ctx, getAllTravelEntry); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllTravelEntry: %w", err)
	}
	if q.getAllTravelEntryByTripIdStmt, err = db.PrepareContext(ctx, getAllTravelEntryByTripId); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllTravelEntryByTripId: %w", err)
	}
	if q.getAllTravelEntryByUserIdAndTripIdStmt, err = db.PrepareContext(ctx, getAllTravelEntryByUserIdAndTripId); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllTravelEntryByUserIdAndTripId: %w", err)
	}
	if q.getConnectionsByUserIdStmt, err = db.PrepareContext(ctx, getConnectionsByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetConnectionsByUserId: %w", err)
	}
	if q.getTravelEntryByIdStmt, err = db.PrepareContext(ctx, getTravelEntryById); err != nil {
		return nil, fmt.Errorf("error preparing query GetTravelEntryById: %w", err)
	}
	if q.getTripStmt, err = db.PrepareContext(ctx, getTrip); err != nil {
		return nil, fmt.Errorf("error preparing query GetTrip: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByUsernameStmt, err = db.PrepareContext(ctx, getUserByUsername); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByUsername: %w", err)
	}
	if q.getUserTripsByTripIdStmt, err = db.PrepareContext(ctx, getUserTripsByTripId); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserTripsByTripId: %w", err)
	}
	if q.getUserTripsByUserIdStmt, err = db.PrepareContext(ctx, getUserTripsByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserTripsByUserId: %w", err)
	}
	if q.listTripsStmt, err = db.PrepareContext(ctx, listTrips); err != nil {
		return nil, fmt.Errorf("error preparing query ListTrips: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.updateMediaByIdStmt, err = db.PrepareContext(ctx, updateMediaById); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMediaById: %w", err)
	}
	if q.updateTravelEntryDescriptionStmt, err = db.PrepareContext(ctx, updateTravelEntryDescription); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTravelEntryDescription: %w", err)
	}
	if q.updateTravelEntryLocationStmt, err = db.PrepareContext(ctx, updateTravelEntryLocation); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTravelEntryLocation: %w", err)
	}
	if q.updateTripCountryStmt, err = db.PrepareContext(ctx, updateTripCountry); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTripCountry: %w", err)
	}
	if q.updateTripEndDateStmt, err = db.PrepareContext(ctx, updateTripEndDate); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTripEndDate: %w", err)
	}
	if q.updateTripStartDateStmt, err = db.PrepareContext(ctx, updateTripStartDate); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTripStartDate: %w", err)
	}
	if q.updateTripTitleStmt, err = db.PrepareContext(ctx, updateTripTitle); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTripTitle: %w", err)
	}
	if q.updateUserEmailStmt, err = db.PrepareContext(ctx, updateUserEmail); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserEmail: %w", err)
	}
	if q.updateUserNameStmt, err = db.PrepareContext(ctx, updateUserName); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserName: %w", err)
	}
	if q.updateUserProfilePictureStmt, err = db.PrepareContext(ctx, updateUserProfilePicture); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserProfilePicture: %w", err)
	}
	if q.updateUserUsernameStmt, err = db.PrepareContext(ctx, updateUserUsername); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserUsername: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCommentStmt != nil {
		if cerr := q.createCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCommentStmt: %w", cerr)
		}
	}
	if q.createConnectionStmt != nil {
		if cerr := q.createConnectionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createConnectionStmt: %w", cerr)
		}
	}
	if q.createMediaStmt != nil {
		if cerr := q.createMediaStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMediaStmt: %w", cerr)
		}
	}
	if q.createTravelEntryStmt != nil {
		if cerr := q.createTravelEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTravelEntryStmt: %w", cerr)
		}
	}
	if q.createTripStmt != nil {
		if cerr := q.createTripStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTripStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.createUserTripStmt != nil {
		if cerr := q.createUserTripStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserTripStmt: %w", cerr)
		}
	}
	if q.deleteCommentByIdStmt != nil {
		if cerr := q.deleteCommentByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCommentByIdStmt: %w", cerr)
		}
	}
	if q.deleteConnectionByUserIdStmt != nil {
		if cerr := q.deleteConnectionByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteConnectionByUserIdStmt: %w", cerr)
		}
	}
	if q.deleteMediaByIdStmt != nil {
		if cerr := q.deleteMediaByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMediaByIdStmt: %w", cerr)
		}
	}
	if q.deleteTravelEntryStmt != nil {
		if cerr := q.deleteTravelEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTravelEntryStmt: %w", cerr)
		}
	}
	if q.deleteTripStmt != nil {
		if cerr := q.deleteTripStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTripStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.deleteUserTripStmt != nil {
		if cerr := q.deleteUserTripStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserTripStmt: %w", cerr)
		}
	}
	if q.getAllMediaByEntryIdStmt != nil {
		if cerr := q.getAllMediaByEntryIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllMediaByEntryIdStmt: %w", cerr)
		}
	}
	if q.getAllMediaByTripIdStmt != nil {
		if cerr := q.getAllMediaByTripIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllMediaByTripIdStmt: %w", cerr)
		}
	}
	if q.getAllMediaByTripIdAndUserIdStmt != nil {
		if cerr := q.getAllMediaByTripIdAndUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllMediaByTripIdAndUserIdStmt: %w", cerr)
		}
	}
	if q.getAllMediaByUserIdStmt != nil {
		if cerr := q.getAllMediaByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllMediaByUserIdStmt: %w", cerr)
		}
	}
	if q.getAllTravelEntryStmt != nil {
		if cerr := q.getAllTravelEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllTravelEntryStmt: %w", cerr)
		}
	}
	if q.getAllTravelEntryByTripIdStmt != nil {
		if cerr := q.getAllTravelEntryByTripIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllTravelEntryByTripIdStmt: %w", cerr)
		}
	}
	if q.getAllTravelEntryByUserIdAndTripIdStmt != nil {
		if cerr := q.getAllTravelEntryByUserIdAndTripIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllTravelEntryByUserIdAndTripIdStmt: %w", cerr)
		}
	}
	if q.getConnectionsByUserIdStmt != nil {
		if cerr := q.getConnectionsByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getConnectionsByUserIdStmt: %w", cerr)
		}
	}
	if q.getTravelEntryByIdStmt != nil {
		if cerr := q.getTravelEntryByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTravelEntryByIdStmt: %w", cerr)
		}
	}
	if q.getTripStmt != nil {
		if cerr := q.getTripStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTripStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByUsernameStmt != nil {
		if cerr := q.getUserByUsernameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByUsernameStmt: %w", cerr)
		}
	}
	if q.getUserTripsByTripIdStmt != nil {
		if cerr := q.getUserTripsByTripIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserTripsByTripIdStmt: %w", cerr)
		}
	}
	if q.getUserTripsByUserIdStmt != nil {
		if cerr := q.getUserTripsByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserTripsByUserIdStmt: %w", cerr)
		}
	}
	if q.listTripsStmt != nil {
		if cerr := q.listTripsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTripsStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.updateMediaByIdStmt != nil {
		if cerr := q.updateMediaByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMediaByIdStmt: %w", cerr)
		}
	}
	if q.updateTravelEntryDescriptionStmt != nil {
		if cerr := q.updateTravelEntryDescriptionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTravelEntryDescriptionStmt: %w", cerr)
		}
	}
	if q.updateTravelEntryLocationStmt != nil {
		if cerr := q.updateTravelEntryLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTravelEntryLocationStmt: %w", cerr)
		}
	}
	if q.updateTripCountryStmt != nil {
		if cerr := q.updateTripCountryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTripCountryStmt: %w", cerr)
		}
	}
	if q.updateTripEndDateStmt != nil {
		if cerr := q.updateTripEndDateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTripEndDateStmt: %w", cerr)
		}
	}
	if q.updateTripStartDateStmt != nil {
		if cerr := q.updateTripStartDateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTripStartDateStmt: %w", cerr)
		}
	}
	if q.updateTripTitleStmt != nil {
		if cerr := q.updateTripTitleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTripTitleStmt: %w", cerr)
		}
	}
	if q.updateUserEmailStmt != nil {
		if cerr := q.updateUserEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserEmailStmt: %w", cerr)
		}
	}
	if q.updateUserNameStmt != nil {
		if cerr := q.updateUserNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserNameStmt: %w", cerr)
		}
	}
	if q.updateUserProfilePictureStmt != nil {
		if cerr := q.updateUserProfilePictureStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserProfilePictureStmt: %w", cerr)
		}
	}
	if q.updateUserUsernameStmt != nil {
		if cerr := q.updateUserUsernameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserUsernameStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                     DBTX
	tx                                     *sql.Tx
	createCommentStmt                      *sql.Stmt
	createConnectionStmt                   *sql.Stmt
	createMediaStmt                        *sql.Stmt
	createTravelEntryStmt                  *sql.Stmt
	createTripStmt                         *sql.Stmt
	createUserStmt                         *sql.Stmt
	createUserTripStmt                     *sql.Stmt
	deleteCommentByIdStmt                  *sql.Stmt
	deleteConnectionByUserIdStmt           *sql.Stmt
	deleteMediaByIdStmt                    *sql.Stmt
	deleteTravelEntryStmt                  *sql.Stmt
	deleteTripStmt                         *sql.Stmt
	deleteUserStmt                         *sql.Stmt
	deleteUserTripStmt                     *sql.Stmt
	getAllMediaByEntryIdStmt               *sql.Stmt
	getAllMediaByTripIdStmt                *sql.Stmt
	getAllMediaByTripIdAndUserIdStmt       *sql.Stmt
	getAllMediaByUserIdStmt                *sql.Stmt
	getAllTravelEntryStmt                  *sql.Stmt
	getAllTravelEntryByTripIdStmt          *sql.Stmt
	getAllTravelEntryByUserIdAndTripIdStmt *sql.Stmt
	getConnectionsByUserIdStmt             *sql.Stmt
	getTravelEntryByIdStmt                 *sql.Stmt
	getTripStmt                            *sql.Stmt
	getUserStmt                            *sql.Stmt
	getUserByEmailStmt                     *sql.Stmt
	getUserByUsernameStmt                  *sql.Stmt
	getUserTripsByTripIdStmt               *sql.Stmt
	getUserTripsByUserIdStmt               *sql.Stmt
	listTripsStmt                          *sql.Stmt
	listUsersStmt                          *sql.Stmt
	updateMediaByIdStmt                    *sql.Stmt
	updateTravelEntryDescriptionStmt       *sql.Stmt
	updateTravelEntryLocationStmt          *sql.Stmt
	updateTripCountryStmt                  *sql.Stmt
	updateTripEndDateStmt                  *sql.Stmt
	updateTripStartDateStmt                *sql.Stmt
	updateTripTitleStmt                    *sql.Stmt
	updateUserEmailStmt                    *sql.Stmt
	updateUserNameStmt                     *sql.Stmt
	updateUserProfilePictureStmt           *sql.Stmt
	updateUserUsernameStmt                 *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                     tx,
		tx:                                     tx,
		createCommentStmt:                      q.createCommentStmt,
		createConnectionStmt:                   q.createConnectionStmt,
		createMediaStmt:                        q.createMediaStmt,
		createTravelEntryStmt:                  q.createTravelEntryStmt,
		createTripStmt:                         q.createTripStmt,
		createUserStmt:                         q.createUserStmt,
		createUserTripStmt:                     q.createUserTripStmt,
		deleteCommentByIdStmt:                  q.deleteCommentByIdStmt,
		deleteConnectionByUserIdStmt:           q.deleteConnectionByUserIdStmt,
		deleteMediaByIdStmt:                    q.deleteMediaByIdStmt,
		deleteTravelEntryStmt:                  q.deleteTravelEntryStmt,
		deleteTripStmt:                         q.deleteTripStmt,
		deleteUserStmt:                         q.deleteUserStmt,
		deleteUserTripStmt:                     q.deleteUserTripStmt,
		getAllMediaByEntryIdStmt:               q.getAllMediaByEntryIdStmt,
		getAllMediaByTripIdStmt:                q.getAllMediaByTripIdStmt,
		getAllMediaByTripIdAndUserIdStmt:       q.getAllMediaByTripIdAndUserIdStmt,
		getAllMediaByUserIdStmt:                q.getAllMediaByUserIdStmt,
		getAllTravelEntryStmt:                  q.getAllTravelEntryStmt,
		getAllTravelEntryByTripIdStmt:          q.getAllTravelEntryByTripIdStmt,
		getAllTravelEntryByUserIdAndTripIdStmt: q.getAllTravelEntryByUserIdAndTripIdStmt,
		getConnectionsByUserIdStmt:             q.getConnectionsByUserIdStmt,
		getTravelEntryByIdStmt:                 q.getTravelEntryByIdStmt,
		getTripStmt:                            q.getTripStmt,
		getUserStmt:                            q.getUserStmt,
		getUserByEmailStmt:                     q.getUserByEmailStmt,
		getUserByUsernameStmt:                  q.getUserByUsernameStmt,
		getUserTripsByTripIdStmt:               q.getUserTripsByTripIdStmt,
		getUserTripsByUserIdStmt:               q.getUserTripsByUserIdStmt,
		listTripsStmt:                          q.listTripsStmt,
		listUsersStmt:                          q.listUsersStmt,
		updateMediaByIdStmt:                    q.updateMediaByIdStmt,
		updateTravelEntryDescriptionStmt:       q.updateTravelEntryDescriptionStmt,
		updateTravelEntryLocationStmt:          q.updateTravelEntryLocationStmt,
		updateTripCountryStmt:                  q.updateTripCountryStmt,
		updateTripEndDateStmt:                  q.updateTripEndDateStmt,
		updateTripStartDateStmt:                q.updateTripStartDateStmt,
		updateTripTitleStmt:                    q.updateTripTitleStmt,
		updateUserEmailStmt:                    q.updateUserEmailStmt,
		updateUserNameStmt:                     q.updateUserNameStmt,
		updateUserProfilePictureStmt:           q.updateUserProfilePictureStmt,
		updateUserUsernameStmt:                 q.updateUserUsernameStmt,
	}
}
