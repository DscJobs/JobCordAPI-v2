package get_user

import (
	"jobcord/api"
	"jobcord/state"
	"jobcord/types"
	"jobcord/utils"
	"net/http"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	docs "github.com/infinitybotlist/eureka/doclib"
	"github.com/infinitybotlist/eureka/dovewing"
)

var (
	sqlCols    = utils.GetCols(types.User{})
	sqlColsStr = strings.Join(sqlCols, ",")
)

func Docs() *docs.Doc {
	return &docs.Doc{
		Summary:     "Get User",
		Description: "Returns a user object from the database.",
		Resp:        types.User{},
		Params: []docs.Parameter{
			{
				Name:        "id",
				In:          "path",
				Description: "The user's ID.",
				Required:    true,
				Schema:      docs.IdSchema,
			},
		},
	}
}

func Route(d api.RouteData, r *http.Request) api.HttpResponse {
	var id = chi.URLParam(r, "id")

	var count int

	err := state.Pool.QueryRow(d.Context, "SELECT COUNT(*) FROM users WHERE userID = $1", id).Scan(&count)

	if err != nil {
		state.Logger.Error(err)
		return api.DefaultResponse(http.StatusInternalServerError)
	}

	if count == 0 {
		return api.DefaultResponse(http.StatusNotFound)
	}

	row, err := state.Pool.Query(d.Context, "SELECT "+sqlColsStr+" FROM users WHERE userID = $1", id)

	if err != nil {
		state.Logger.Error(err)
		return api.DefaultResponse(http.StatusInternalServerError)
	}

	var user = types.User{}

	err = pgxscan.ScanOne(&user, row)

	if err != nil {
		state.Logger.Error(err)
		return api.DefaultResponse(http.StatusNotFound)
	}

	userObj, err := dovewing.GetDiscordUser(d.Context, id)

	if err != nil {
		state.Logger.Error(err)
		return api.DefaultResponse(http.StatusInternalServerError)
	}

	user.User = userObj

	return api.HttpResponse{
		Json: user,
	}
}
