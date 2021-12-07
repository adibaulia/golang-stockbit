package models

type (
	Request struct {
		IMDBID string `json:"imdbid,omitempty"`
		Name   string `json:"name,omitempty"`
		Page   int    `json:"page,omitempty"`
	}

	Movie struct {
		Title      string   `json:"Title,omitempty"`
		Year       string   `json:"Year,omitempty"`
		Rated      string   `json:"Rated,omitempty"`
		Released   string   `json:"Released,omitempty"`
		Runtime    string   `json:"Runtime,omitempty"`
		Genre      string   `json:"Genre,omitempty"`
		Director   string   `json:"Director,omitempty"`
		Writer     string   `json:"Writer,omitempty"`
		Actors     string   `json:"Actors,omitempty"`
		Plot       string   `json:"Plot,omitempty"`
		Language   string   `json:"Language,omitempty"`
		Country    string   `json:"Country,omitempty"`
		Awards     string   `json:"Awards,omitempty"`
		Poster     string   `json:"Poster,omitempty"`
		Ratings    []Rating `json:"Ratings,omitempty"`
		Metascore  string   `json:"Metascore,omitempty"`
		ImdbRating string   `json:"imdbRating,omitempty"`
		ImdbVotes  string   `json:"imdbVotes,omitempty"`
		ImdbID     string   `json:"imdbID,omitempty"`
		Type       string   `json:"Type,omitempty"`
		Dvd        string   `json:"DVD,omitempty"`
		BoxOffice  string   `json:"BoxOffice,omitempty"`
		Production string   `json:"Production,omitempty"`
		Website    string   `json:"Website,omitempty"`
		Response   string   `json:"Response,omitempty"`
	}

	Rating struct {
		Source string `json:"Source,omitempty"`
		Value  string `json:"Value,omitempty"`
	}

	Response struct {
		StatusCode   int         `json:"statusCode"`
		Status       string      `json:"status"`
		Data         interface{} `json:"data"`
		Message      string      `json:"message"`
		ErrorMessage string      `json:"errorMessage"`
	}
)
