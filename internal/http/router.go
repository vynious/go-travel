package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vynious/go-travel/internal/domains/connections"
	"github.com/vynious/go-travel/internal/domains/travel_entry"
	"github.com/vynious/go-travel/internal/domains/trip"
	"github.com/vynious/go-travel/internal/domains/user"
	"github.com/vynious/go-travel/internal/domains/user_trip"
	"net/http"
)

func InitRouter(
	userHandler *user.UserHandler,
	tripHandler *trip.TripHandler,
	usertripHandler *user_trip.UserTripHandler,
	travelEntryHandler *travel_entry.TravelEntryHandler,
	connectionHandler *connections.ConnectionHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "API working!")
		writer.WriteHeader(http.StatusOK)

	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.RegisterUser)
		r.Post("/token", userHandler.GenerateToken)
		r.Get("/{userId}", userHandler.ViewUserDetails)
		r.Get("/", userHandler.ViewAllUsers)
		r.Get("/search", userHandler.SearchUser)
		r.Patch("/{userId}/profile-picture", userHandler.ChangeUserProfilePicture)
		r.Patch("/{userId}/details", userHandler.ChangeUserDetails)
		r.Delete("/{userId}", userHandler.DeleteAccount)
	})

	r.Route("/trips", func(r chi.Router) {
		r.Post("/", tripHandler.StartTrip)
		r.Get("/{tripId}", tripHandler.ViewTripDetails)
		r.Get("/", tripHandler.ViewAllTrips)
		r.Patch("/{tripId}", tripHandler.ChangeTripDetails)
		r.Delete("/{tripId}", tripHandler.DeleteTrip)
	})

	r.Route("/trip-assignments", func(r chi.Router) {
		r.Post("/{tripId}/users", usertripHandler.AddUserToTrip)
		r.Delete("/{tripId}/users/{userId}", usertripHandler.DeleteUserFromTripId)
		r.Get("/{tripId}/users", usertripHandler.GetAllUsersOnTripId)
		r.Get("/users/{userId}", usertripHandler.GetAllTripsForUserId)
	})

	r.Route("/travel-entries", func(r chi.Router) {
		r.Post("/", travelEntryHandler.EnterTravelEntry)
		r.Get("/{entryId}", travelEntryHandler.ViewTravelEntry)
		r.Get("/trips/{tripId}", travelEntryHandler.ViewTravelEntriesUnderTrip)
		r.Get("/users/{userId}/trips/{tripId}", travelEntryHandler.ViewTravelEntriesUnderTripAndUser)
		r.Patch("/{entryId}", travelEntryHandler.UpdateTravelEntry)
		r.Delete("/{entryId}", travelEntryHandler.DeleteTravelEntry)
	})

	r.Route("/connection", func(r chi.Router) {
		r.Post("/{partyA}/{partyB}", connectionHandler.MakeConnection)
		r.Get("/{userId}", connectionHandler.ViewConnections)
		r.Delete("/{userId}", connectionHandler.RemoveConnection)
	})
	return r
}
