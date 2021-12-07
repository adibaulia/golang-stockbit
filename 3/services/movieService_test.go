package services

import (
	"encoding/json"
	"testing"

	"github.com/adibaulia/stockbit-test/3/imdb/config"
	"github.com/adibaulia/stockbit-test/3/imdb/db"
	"github.com/adibaulia/stockbit-test/3/imdb/db/models"
	"github.com/stretchr/testify/assert"
)

var (
	expectedResponseGetMovieJson = `{"Title":"Batman Begins","Year":"2005","Rated":"PG-13","Released":"15 Jun 2005","Runtime":"140 min","Genre":"Action, Adventure","Director":"Christopher Nolan","Writer":"Bob Kane, David S. Goyer, Christopher Nolan","Actors":"Christian Bale, Michael Caine, Ken Watanabe","Plot":"After training with his mentor, Batman begins his fight to free crime-ridden Gotham City from corruption.","Language":"English, Mandarin","Country":"United Kingdom, United States","Awards":"Nominated for 1 Oscar. 13 wins & 79 nominations total","Poster":"https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"8.2/10"},{"Source":"Rotten Tomatoes","Value":"84%"},{"Source":"Metacritic","Value":"70/100"}],"Metascore":"70","imdbRating":"8.2","imdbVotes":"1,369,540","imdbID":"tt0372784","Type":"movie","DVD":"18 Oct 2005","BoxOffice":"$206,852,432","Production":"N/A","Website":"N/A"}`
	c                            = config.GetConnection()
	dao                          = db.NewMovieDB(c.Api)
	api                          = NewMovieService(dao)
)

func TestGetMovieByID(t *testing.T) {

	res, err := api.Get("tt0372784")
	if err != nil {
		panic(err)
	}

	var expectedResult models.Movie

	err = json.Unmarshal([]byte(expectedResponseGetMovieJson), &expectedResult)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, expectedResult, *res, "movie")
}
