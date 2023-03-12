package utils

import (
	"jobcord/state"
	"jobcord/types"
	"net/http"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func GetCols(s any) []string {
	refType := reflect.TypeOf(s)

	var cols []string

	for _, f := range reflect.VisibleFields(refType) {
		db := f.Tag.Get("db")
		reflectOpts := f.Tag.Get("reflect")

		if db == "-" || db == "" || reflectOpts == "ignore" {
			continue
		}

		cols = append(cols, db)
	}

	return cols
}

// Using popplio for now until this service grows big enough to warrant its own
func GetDiscordUser(id string) (*types.IUser, error) {
	res, err := http.Get(state.Config.PopplioURL + "/_duser/" + id)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var user types.IUser

	err = json.NewDecoder(res.Body).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
