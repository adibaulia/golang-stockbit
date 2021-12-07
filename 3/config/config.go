package config

import "github.com/adibaulia/stockbit-test/3/imdb/utils/package/omdbapi"

var Instance *Connection

type (
	Connection struct {
		Api *omdbapi.OMDBAPI
	}
)

func init() {
	//work like connecting to Database
	api := omdbapi.NewOMDBAPI()

	Instance = &Connection{api}
}

func GetConnection() *Connection {
	return Instance
}
