package users

import (
	"jobcord/routes/users/endpoints/get_user"
	"jobcord/routes/users/endpoints/login"

	"github.com/go-chi/chi/v5"
	"github.com/infinitybotlist/eureka/uapi"
)

const tagName = "Users"

type Router struct{}

func (b Router) Tag() (string, string) {
	return tagName, "These API endpoints are regarding users."
}

func (b Router) Routes(r *chi.Mux) {
	uapi.Route{
		Pattern: "/users/{id}",
		OpId:    "get_user",
		Method:  uapi.GET,
		Docs:    get_user.Docs,
		Handler: get_user.Route,
	}.Route(r)

	uapi.Route{
		Pattern: "/users",
		OpId:    "login",
		Method:  uapi.PUT,
		Docs:    login.Docs,
		Handler: login.Route,
	}.Route(r)
}
