package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vynious/go-travel/internal/domains/trip"
	"github.com/vynious/go-travel/internal/domains/user"
	"net/http"
)

func InitRouter(userHandler *user.UserHandler, tripHandler *trip.TripHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "API working!")
		writer.WriteHeader(http.StatusOK)

	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/create", userHandler.RegisterUser)
		r.Post("/login", userHandler.LoginUser)
		r.Get("/view/{id}", userHandler.ViewUserDetails)
		r.Get("/view/all", userHandler.ViewAllUsers)
		r.Get("/search", userHandler.SearchUser)
		r.Patch("/update/{id}/profile_picture", userHandler.ChangeUserProfilePicture)
		r.Patch("/update/{id}/details", userHandler.ChangeUserDetails)
		r.Delete("/delete/{id}", userHandler.DeleteAccount)
	})

	r.Route("/trip", func(r chi.Router) {
		r.Post("/create", tripHandler.StartTrip)
		r.Get("/view/{id}", tripHandler.ViewTripDetails)
		r.Get("/view/all", tripHandler.ViewAllTrips)
		r.Patch("/update/{id}", tripHandler.ChangeTripDetails)
		r.Delete("/delete/{id}", tripHandler.DeleteTrip)
	})

	return r
}
