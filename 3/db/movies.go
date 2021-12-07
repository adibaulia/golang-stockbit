package db

import (
	"fmt"
	"log"

	"github.com/adibaulia/stockbit-test/3/imdb/db/models"
	"github.com/adibaulia/stockbit-test/3/imdb/utils/package/omdbapi"
)

type DB struct {
	dao *omdbapi.OMDBAPI
}

func NewMovieDB(dao *omdbapi.OMDBAPI) *DB { return &DB{dao} }

func (db *DB) Get(ImdbID string) (*models.Movie, error) {
	res, err := db.dao.GetByIMDBID(ImdbID)
	if err != nil {
		return nil, fmt.Errorf("error when getting data source movie, because %v", err)
	}

	result := db.composeToMovieStruct(res)
	return result, nil
}

func (db *DB) Search(name string, page string) ([]models.Movie, error) {

	if name == "" {
		return nil, fmt.Errorf("error when searching movies, name not provided")
	}

	if page == "" {
		page = "1"
	}

	log.Printf("[Search] Searching movie with param name = %v and page = %v", name, page)

	res, err := db.dao.SearchByName(name, page)
	if err != nil {
		return nil, fmt.Errorf("error when getting data source movie, because %v", err)
	}

	log.Printf("[Search] Found %v movies", len(res.Movies))

	return db.composeToMoviesStruct(res.Movies), nil
}

func (db *DB) composeToMoviesStruct(res []omdbapi.Movie) []models.Movie {
	var movies []models.Movie

	for _, movieResponse := range res {
		movies = append(movies, *db.composeToMovieStruct(&movieResponse))
	}

	return movies
}

func (db *DB) composeToMovieStruct(res *omdbapi.Movie) *models.Movie {
	var rating []models.Rating

	for _, r := range res.Ratings {
		rating = append(rating, models.Rating{r.Source, r.Value})
	}
	return &models.Movie{
		Title:      res.Title,
		Year:       res.Year,
		Rated:      res.Rated,
		Released:   res.Released,
		Runtime:    res.Runtime,
		Genre:      res.Genre,
		Director:   res.Director,
		Writer:     res.Writer,
		Actors:     res.Actors,
		Plot:       res.Plot,
		Language:   res.Language,
		Country:    res.Country,
		Awards:     res.Awards,
		Poster:     res.Poster,
		Ratings:    rating,
		Metascore:  res.Metascore,
		ImdbRating: res.ImdbRating,
		ImdbVotes:  res.ImdbVotes,
		ImdbID:     res.ImdbID,
		Type:       res.Type,
		Dvd:        res.Dvd,
		BoxOffice:  res.BoxOffice,
		Production: res.Production,
		Website:    res.Website,
	}
}
