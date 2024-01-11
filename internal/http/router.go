package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/vynious/go-travel/internal/domains/user"
)

func InitRouter(userHandler *user.Handler) {
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Post("/create", userHandler.RegisterUser)
		r.Post("/login", userHandler.LoginUser)
		r.Get("/view/{id}", userHandler.ViewUserDetails)
		r.Get("/search", userHandler.SearchUser)
		r.Patch("/update/{id}/profile_picture", userHandler.ChangeUserProfilePicture)
		r.Patch("/update/{id}/details", userHandler.ChangeUserDetails)
		r.Delete("/delete/{id}", userHandler.DeleteAccount)
	})

}
