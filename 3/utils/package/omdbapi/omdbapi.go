package omdbapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	baseURL   = "http://www.omdbapi.com/"
	OMDBKey   = "faf7e5bb"
	urlAPIKey = fmt.Sprintf("%v?apikey=%v", baseURL, OMDBKey)
	client    = &http.Client{}
)

type (
	MoviesSearch struct {
		Movies       []Movie `json:"Search,omitempty"`
		TotalResults string  `json:"totalResults,omitempty"`
		Response     string  `json:"Response,omitempty"`
	}

	Movie struct {
		Title      string    `json:"Title,omitempty"`
		Year       string    `json:"Year,omitempty"`
		Rated      string    `json:"Rated,omitempty"`
		Released   string    `json:"Released,omitempty"`
		Runtime    string    `json:"Runtime,omitempty"`
		Genre      string    `json:"Genre,omitempty"`
		Director   string    `json:"Director,omitempty"`
		Writer     string    `json:"Writer,omitempty"`
		Actors     string    `json:"Actors,omitempty"`
		Plot       string    `json:"Plot,omitempty"`
		Language   string    `json:"Language,omitempty"`
		Country    string    `json:"Country,omitempty"`
		Awards     string    `json:"Awards,omitempty"`
		Poster     string    `json:"Poster,omitempty"`
		Ratings    []Ratings `json:"Ratings,omitempty"`
		Metascore  string    `json:"Metascore,omitempty"`
		ImdbRating string    `json:"imdbRating,omitempty"`
		ImdbVotes  string    `json:"imdbVotes,omitempty"`
		ImdbID     string    `json:"imdbID,omitempty"`
		Type       string    `json:"Type,omitempty"`
		Dvd        string    `json:"DVD,omitempty"`
		BoxOffice  string    `json:"BoxOffice,omitempty"`
		Production string    `json:"Production,omitempty"`
		Website    string    `json:"Website,omitempty"`
		Response   string    `json:"Response,omitempty"`
	}
	Ratings struct {
		Source string `json:"Source,omitempty"`
		Value  string `json:"Value,omitempty"`
	}

	OMDBAPI struct {
	}
)

func NewOMDBAPI() *OMDBAPI {
	return &OMDBAPI{}
}

func (m *OMDBAPI) GetByIMDBID(ImdbID string) (*Movie, error) {

	if ImdbID == "" {
		return nil, fmt.Errorf("cannot find movie because IMDB ID is nil")
	}
	url := fmt.Sprintf("%v&i=%v", urlAPIKey, ImdbID)

	var result *Movie
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("[OMDAPI] requesting API to url = %v", url)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *OMDBAPI) SearchByName(name string, page string) (*MoviesSearch, error) {

	if name == "" {
		return nil, fmt.Errorf("cannot find movie because name is nil")
	}
	url := fmt.Sprintf("%v&s=%v&page=%v", urlAPIKey, name, page)

	var result *MoviesSearch

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("[OMDAPI] requesting API to url = %v", url)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
