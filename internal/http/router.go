package http

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	auth "github.com/vynious/go-travel/internal/domains/auth"
	"github.com/vynious/go-travel/internal/domains/connections"
	"github.com/vynious/go-travel/internal/domains/travel_entry"
	"github.com/vynious/go-travel/internal/domains/trip"
	"github.com/vynious/go-travel/internal/domains/user"
	"github.com/vynious/go-travel/internal/domains/user_trip"
)

type AppRouter struct {
	router         chi.Router
	firebaseClient *auth.FBClient
}

func NewAppRouter(
	userHandler *user.UserHandler,
	tripHandler *trip.TripHandler,
	usertripHandler *user_trip.UserTripHandler,
	travelEntryHandler *travel_entry.TravelEntryHandler,
	connectionHandler *connections.ConnectionHandler,
	fbClient *auth.FBClient) *AppRouter {
	r := chi.NewRouter()

	ar := &AppRouter{
		router:         r,
		firebaseClient: fbClient,
	}
	ar.setupRoutes(userHandler, tripHandler, usertripHandler, travelEntryHandler, connectionHandler)
	return ar
}

func (ar *AppRouter) setupRoutes(
	userHandler *user.UserHandler,
	tripHandler *trip.TripHandler,
	usertripHandler *user_trip.UserTripHandler,
	travelEntryHandler *travel_entry.TravelEntryHandler,
	connectionHandler *connections.ConnectionHandler,
) {

	ar.router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

	ar.router.Use(ar.LogRequest)
	// ar.router.Use(ar.AuthenticateRequest)
	ar.router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, "API working!")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		writer.WriteHeader(http.StatusOK)
	})

	ar.router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, "API working!")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		writer.WriteHeader(http.StatusOK)

	})

	ar.router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.RegisterUser)
		r.Post("/token", userHandler.GenerateToken)
		r.Get("/{userId}", userHandler.ViewUserDetails)
		r.Get("/", userHandler.ViewAllUsers)
		r.Get("/search", userHandler.SearchUser)
		r.Patch("/{userId}/profile-picture", userHandler.ChangeUserProfilePicture)
		r.Patch("/{userId}/details", userHandler.ChangeUserDetails)
		r.Delete("/{userId}", userHandler.DeleteAccount)
	})

	ar.router.Route("/trips", func(r chi.Router) {
		r.Post("/", tripHandler.StartTrip)
		r.Get("/{tripId}", tripHandler.ViewTripDetails)
		r.Get("/", tripHandler.ViewAllTrips)
		r.Patch("/{tripId}", tripHandler.ChangeTripDetails)
		r.Delete("/{tripId}", tripHandler.DeleteTrip)
	})

	ar.router.Route("/trip-assignments", func(r chi.Router) {
		r.Post("/{tripId}/users", usertripHandler.AddUsersToTrip)
		r.Delete("/{tripId}/users/{userId}", usertripHandler.DeleteUserFromTripId)
		r.Get("/{tripId}/users", usertripHandler.GetAllUsersOnTripId)
		r.Get("/users/{userId}", usertripHandler.GetAllTripsForUserId)
	})

	// todo missing update the media for the travel entry
	ar.router.Route("/travel-entries", func(r chi.Router) {
		r.Post("/", travelEntryHandler.EnterTravelEntry)
		r.Get("/{entryId}", travelEntryHandler.ViewTravelEntry)
		r.Get("/trips/{tripId}", travelEntryHandler.ViewTravelEntriesUnderTrip)
		r.Get("/users/{userId}/trips/{tripId}", travelEntryHandler.ViewTravelEntriesUnderTripAndUser)
		r.Patch("/{entryId}", travelEntryHandler.UpdateTravelEntry)
		r.Delete("/{entryId}", travelEntryHandler.DeleteTravelEntry)
	})

	ar.router.Route("/connection", func(r chi.Router) {
		r.Post("/{partyA}/{partyB}", connectionHandler.MakeConnection)
		r.Get("/{userId}", connectionHandler.ViewConnections)
		r.Delete("/{userId}", connectionHandler.RemoveConnection)
	})

}

func (ar *AppRouter) LogRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)

		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"time":   time.Since(start),
		}).Info("request handled")
	})
}

func (ar *AppRouter) AuthenticateRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("authorisation")
		authToken = strings.TrimPrefix(authToken, "Bearer ")
		token, err := ar.firebaseClient.VerifyToken(r.Context(), authToken)
		if err != nil {
			http.Error(w, "Invalid ID token", http.StatusUnauthorized)
			return
		}
		key := "uid"
		ctx := context.WithValue(r.Context(), key, token.UID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
