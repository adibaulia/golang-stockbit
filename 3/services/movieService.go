package services

import (
	"log"
	"strconv"

	"github.com/adibaulia/stockbit-test/3/imdb/db/models"
)

type (
	MoviesDB interface {
		Get(ImdbID string) (*models.Movie, error)
		Search(name string, page string) ([]models.Movie, error)
	}

	MoviesSvc struct {
		MoviesDB
	}
)

func NewMovieService(db MoviesDB) *MoviesSvc { return &MoviesSvc{db} }

func (svc *MoviesSvc) GetMovieByID(ImdbID string) (*models.Movie, error) {
	log.Printf("[Service] doing some bussiness work in here")
	return svc.Get(ImdbID)
}

func (svc *MoviesSvc) SearchMovieByName(name string, page int) ([]models.Movie, error) {
	log.Printf("[Service] doing some bussiness work in here")

	pageString := strconv.Itoa(page)

	return svc.Search(name, pageString)
}
