package login

import (
	"io"
	"jobcord/api"
	"jobcord/config"
	"jobcord/state"
	"jobcord/types"
	"jobcord/utils"
	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
	docs "github.com/infinitybotlist/doclib"
	"github.com/infinitybotlist/eureka/crypto"
	jsoniter "github.com/json-iterator/go"
)

var (
	json         = jsoniter.ConfigCompatibleWithStandardLibrary
	compiledMsgs = api.CompileValidationErrors(LoginReq{})
)

type UserAuth struct {
	Token       string       `json:"token" validate:"required"`
	AccessToken string       `json:"access_token" validate:"required"`
	ID          string       `json:"id" validate:"required"`
	User        *types.IUser `json:"user" validate:"required"`
}

type LoginReq struct {
	Code        string `json:"code" validate:"required"`
	RedirectUri string `json:"redirect_uri" validate:"required"`
	Nonce       string `json:"nonce" validate:"required"`
	ClientID    string `json:"client_id" validate:"required" msg:"Client ID must be the Client ID to login to"`
}

func Docs() *docs.Doc {
	return &docs.Doc{
		Summary:     "Login User",
		Description: "Login user based on code.",
		Req:         LoginReq{},
		Resp:        types.User{},
	}
}

func Route(d api.RouteData, r *http.Request) api.HttpResponse {
	var req LoginReq

	resp, ok := api.MarshalReq(r, &req)

	if !ok {
		return resp
	}

	err := state.Validator.Struct(req)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		return api.ValidatorErrorResponse(compiledMsgs, errors)
	}

	// Get access token from code
	var token struct {
		AccessToken string `json:"access_token"`
	}

	if req.Nonce != "winterwatcher" {
		return api.HttpResponse{
			Json: types.ApiError{
				Error:   true,
				Message: "Your client is outdated and is not supported. Please update your client.",
			},
			Status: http.StatusBadRequest,
		}
	}

	// Find client ID in config
	var client *config.LoginClient

	for _, c := range state.Config.LoginClients {
		if c.ClientID == req.ClientID {
			client = &c
			break
		}
	}

	if client == nil {
		return api.HttpResponse{
			Json: types.ApiError{
				Error:   true,
				Message: "Invalid client ID",
			},
			Status: http.StatusBadRequest,
		}
	}

	if client.RedirectURL != req.RedirectUri {
		state.Logger.Info("Expected redirect URI: "+client.RedirectURL, " but got "+req.RedirectUri)
		return api.HttpResponse{
			Json: types.ApiError{
				Error:   true,
				Message: "Invalid redirect URI",
			},
			Status: http.StatusBadRequest,
		}
	}

	response, err := http.PostForm("https://discord.com/api/v10/oauth2/token", url.Values{
		"client_id":     {client.ClientID},
		"client_secret": {client.ClientSecret},
		"grant_type":    {"authorization_code"},
		"code":          {req.Code},
		"redirect_uri":  {req.RedirectUri},
		"scope":         {"identify guilds"},
	})

	if err != nil {
		state.Logger.Error(err)
		return api.DefaultResponse(http.StatusInternalServerError)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bytes, err := io.ReadAll(response.Body)

		if err != nil {
			state.Logger.Error(err)
			return api.DefaultResponse(http.StatusInternalServerError)
		}

		return api.HttpResponse{
			Status: http.StatusUnauthorized,
			Json: types.ApiError{
				Message: string(bytes),
				Error:   true,
			},
		}
	}

	err = json.NewDecoder(response.Body).Decode(&token)

	if err != nil {
		state.Logger.Error(err)
		return api.DefaultResponse(http.StatusInternalServerError)
	}

	if token.AccessToken == "" {
		return api.HttpResponse{
			Status: http.StatusUnauthorized,
			Json: types.ApiError{
				Message: "Invalid access token",
				Error:   true,
			},
		}
	}

	var userData struct {
		ID string `json:"id"`
	}

	// Get user data from access token in Authorization header
	cli := &http.Client{}

	var httpReq *http.Request
	httpReq, err = http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)

	if err != nil {
		state.Logger.Error(err)
		return api.HttpResponse{
			Json: types.ApiError{
				Error:   true,
				Message: "Failed to create request to Discord to fetch user info",
			},
			Status: http.StatusInternalServerError,
		}
	}

	httpReq.Header.Set("Authorization", "Bearer "+token.AccessToken)

	var httpResp *http.Response

	httpResp, err = cli.Do(httpReq)

	if err != nil {
		state.Logger.Error(err)
		return api.HttpResponse{
			Json: types.ApiError{
				Error:   true,
				Message: "Failed to create request to Discord to fetch user info",
			},
			Status: http.StatusInternalServerError,
		}
	}

	defer httpResp.Body.Close()

	err = json.NewDecoder(httpResp.Body).Decode(&userData)

	if err != nil {
		state.Logger.Error(err)
		return api.HttpResponse{
			Json: types.ApiError{
				Error:   true,
				Message: "Failed to create request to Discord to fetch user info",
			},
			Status: http.StatusInternalServerError,
		}
	}

	var count int

	err = state.Pool.QueryRow(d.Context, "SELECT COUNT(*) FROM users WHERE userID = $1", userData.ID).Scan(&count)

	if err != nil {
		state.Logger.Error(err)
		return api.DefaultResponse(http.StatusInternalServerError)
	}

	var apitoken string
	if count == 0 {
		apitoken = crypto.RandString(255)
		_, err = state.Pool.Exec(d.Context, "INSERT INTO users (userID, token) VALUES ($1, $2)", userData.ID, apitoken)

		if err != nil {
			state.Logger.Error(err)
			return api.DefaultResponse(http.StatusInternalServerError)
		}
	} else {
		var banned bool
		err = state.Pool.QueryRow(d.Context, "SELECT token, banned FROM users WHERE userID = $1", userData.ID).Scan(&apitoken, &banned)

		if err != nil {
			state.Logger.Error(err)
			return api.DefaultResponse(http.StatusInternalServerError)
		}

		if banned {
			return api.HttpResponse{
				Json: types.ApiError{
					Error:   true,
					Message: "You are banned from using this API.",
				},
				Status: http.StatusForbidden,
			}
		}
	}

	user, err := utils.GetDiscordUser(userData.ID)

	if err != nil {
		state.Logger.Error(err)
		return api.DefaultResponse(http.StatusInternalServerError)
	}

	return api.HttpResponse{
		Json: UserAuth{
			User:        user,
			ID:          userData.ID,
			Token:       apitoken,
			AccessToken: token.AccessToken,
		},
	}
}
