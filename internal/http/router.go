package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/vynious/go-travel/internal/user"
)

func InitRouter(uh *user.Handler) {
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Post("/create", uh.RegisterUser)
		r.Post("/login", uh.LoginUser)
		r.Get("/view/{id}", uh.ViewUserDetails)
		r.Get("/search", uh.SearchUser)
		r.Patch("/update/{id}/profile_picture", uh.ChangeUserProfilePicture)
		r.Patch("/update/{id}/details", uh.ChangeUserDetails)
		r.Delete("/delete/{id}", uh.DeleteAccount)
	})

}
