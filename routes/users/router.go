package users

import (
	"jobcord/api"
	"jobcord/routes/users/endpoints/get_user"
	"jobcord/routes/users/endpoints/login"

	"github.com/go-chi/chi/v5"
)

const tagName = "Users"

type Router struct{}

func (b Router) Tag() (string, string) {
	return tagName, "These API endpoints are regarding users."
}

func (b Router) Routes(r *chi.Mux) {
	api.Route{
		Pattern: "/users/{id}",
		OpId:    "get_user",
		Method:  api.GET,
		Docs:    get_user.Docs,
		Handler: get_user.Route,
	}.Route(r)

	api.Route{
		Pattern: "/users",
		OpId:    "login",
		Method:  api.PUT,
		Docs:    login.Docs,
		Handler: login.Route,
	}.Route(r)
}
