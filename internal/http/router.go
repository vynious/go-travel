package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vynious/go-travel/internal/domains/trip"
	"github.com/vynious/go-travel/internal/domains/user"
	"github.com/vynious/go-travel/internal/domains/user_trip"
	"net/http"
)

func InitRouter(userHandler *user.UserHandler, tripHandler *trip.TripHandler, usertripHandler *user_trip.UserTripHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "API working!")
		writer.WriteHeader(http.StatusOK)

	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/create", userHandler.RegisterUser)
		r.Post("/token", userHandler.GenerateToken)
		r.Get("/view/{id}", userHandler.ViewUserDetails)
		r.Get("/view/all", userHandler.ViewAllUsers)
		r.Get("/search", userHandler.SearchUser)
		r.Patch("/update/{id}/profile_picture", userHandler.ChangeUserProfilePicture)
		r.Patch("/update/{id}/details", userHandler.ChangeUserDetails)
		r.Delete("/delete/{id}", userHandler.DeleteAccount)

		r.Get("/view-trips/{id}", usertripHandler.GetAllTripsForUserId)
	})

	r.Route("/trip", func(r chi.Router) {
		r.Post("/create", tripHandler.StartTrip)
		r.Get("/view/{id}", tripHandler.ViewTripDetails)
		r.Get("/view/all", tripHandler.ViewAllTrips)
		r.Patch("/update/{id}", tripHandler.ChangeTripDetails)
		r.Delete("/delete/{id}", tripHandler.DeleteTrip)

		r.Get("/view-users/{id}", usertripHandler.GetAllUsersOnTripId)

	})

	r.Route("/assign-trip", func(r chi.Router) {
		r.Post("/add-user", usertripHandler.AddUserToTrip)
		r.Post("/remove-user", usertripHandler.DeleteUserFromTripId)
	})

	return r
}
