package get_cv

import (
	"jobcord/state"
	"jobcord/types"
	"jobcord/utils"
	"net/http"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	docs "github.com/infinitybotlist/eureka/doclib"
	"github.com/infinitybotlist/eureka/dovewing"
	"github.com/infinitybotlist/eureka/uapi"
)

var (
	sqlCols    = utils.GetCols(types.CV{})
	sqlColsStr = strings.Join(sqlCols, ",")
)

func Docs() *docs.Doc {
	return &docs.Doc{
		Summary:     "Get User CV",
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

func Route(d uapi.RouteData, r *http.Request) uapi.HttpResponse {
	var id = chi.URLParam(r, "id")

	var count int

	err := state.Pool.QueryRow(d.Context, "SELECT COUNT(*) FROM cv WHERE userID = $1", id).Scan(&count)

	if err != nil {
		state.Logger.Error(err)
		return uapi.DefaultResponse(http.StatusInternalServerError)
	}

	if count == 0 {
		return uapi.DefaultResponse(http.StatusNotFound)
	}

	row, err := state.Pool.Query(d.Context, "SELECT "+sqlColsStr+" FROM cv WHERE userID = $1", id)

	if err != nil {
		state.Logger.Error(err)
		return uapi.DefaultResponse(http.StatusInternalServerError)
	}

	var cv = types.CV{}

	err = pgxscan.ScanOne(&cv, row)

	if err != nil {
		state.Logger.Error(err)
		return uapi.DefaultResponse(http.StatusNotFound)
	}

	userObj, err := dovewing.GetDiscordUser(d.Context, id)

	if err != nil {
		state.Logger.Error(err)
		return uapi.DefaultResponse(http.StatusInternalServerError)
	}

	cv.User = userObj

	return uapi.HttpResponse{
		Json: cv,
	}
}
