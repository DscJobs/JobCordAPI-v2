package ping

import (
	"net/http"

	"jobcord/state"

	docs "github.com/infinitybotlist/eureka/doclib"
	"github.com/infinitybotlist/eureka/uapi"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Hello struct {
	Message string `json:"message"`
	Docs    string `json:"docs"`
	OurSite string `json:"our_site"`
	Status  string `json:"status"`
}

const (
	mainSite   = "https://jobcord.co"
	statusPage = "https://status.jobcord.co"
)

var helloWorld []byte
var helloWorldB Hello

func Setup() {
	var docsSite string = state.Config.APIUrl + "/docs"

	// This is done here to avoid constant remarshalling
	helloWorldB = Hello{
		Message: "Hello world from Jobcord API v4!",
		Docs:    docsSite,
		OurSite: mainSite,
		Status:  statusPage,
	}

	var err error
	helloWorld, err = json.Marshal(helloWorldB)

	if err != nil {
		panic(err)
	}
}

func Docs() *docs.Doc {
	return &docs.Doc{
		Summary:     "Ping Server",
		Description: "This is a simple ping endpoint to check if the API is online. It will return a simple JSON object with a message, docs link, our site link and status page link.",
		Resp:        helloWorldB,
	}
}

func Route(d uapi.RouteData, r *http.Request) uapi.HttpResponse {
	return uapi.HttpResponse{
		Bytes: helloWorld,
	}
}
