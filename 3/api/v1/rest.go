package api

import (
	"log"
	"net/http"

	"github.com/adibaulia/stockbit-test/3/imdb/db/models"
	"github.com/labstack/echo/v4"
)

type (
	MovieService interface {
		GetMovieByID(ImdbID string) (*models.Movie, error)
		SearchMovieByName(name string, page int) ([]models.Movie, error)
	}

	Rest struct {
		svc MovieService
	}
)

func Route(e *echo.Echo, svc MovieService) {

	r := Rest{svc}

	e.GET("/get", r.GetByIMDBID)
	e.GET("/search", r.SearchMovie)
}

func (r *Rest) GetByIMDBID(c echo.Context) error {
	body := new(models.Request)
	c.Bind(body)

	log.Printf("[REST] incoming request from rest api...")
	data, err := r.svc.GetMovieByID(body.IMDBID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	response := models.Response{
		StatusCode: http.StatusOK,
		Status:     "OK",
		Data:       data,
	}

	return c.JSON(http.StatusOK, response)
}

func (r *Rest) SearchMovie(c echo.Context) error {
	body := new(models.Request)
	c.Bind(body)

	log.Printf("[REST] incoming request from rest api...")
	data, err := r.svc.SearchMovieByName(body.Name, body.Page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	response := models.Response{
		StatusCode: http.StatusOK,
		Status:     "OK",
		Data:       data,
	}

	return c.JSON(http.StatusOK, response)
}
