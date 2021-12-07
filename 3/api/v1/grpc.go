package api

import (
	context "context"
	"log"

	"github.com/adibaulia/stockbit-test/3/imdb/db/models"
)

type GRPC struct {
	SVC MovieService
}

func (g *GRPC) mustEmbedUnimplementedMovieServer() {}
func (g *GRPC) GetMovieByIMDBID(ctx context.Context, input *MovieRequest) (*MovieDetail, error) {

	log.Printf("[GRPC] incoming request from GRPC api...")
	data, err := g.SVC.GetMovieByID(input.IMDBID)
	if err != nil {
		return nil, err
	}

	return g.composeToMovieStruct(data), nil
}

func (g *GRPC) SearchMovieByName(in *MovieRequest, stream Movie_SearchMovieByNameServer) error {

	log.Printf("[GRPC] incoming request from GRPC api...")
	data, err := g.SVC.SearchMovieByName(in.Name, int(in.Page))
	if err != nil {
		return err
	}

	for _, m := range data {
		err := stream.Send(g.composeToMovieStruct(&m))
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GRPC) composeToMovieStruct(res *models.Movie) *MovieDetail {
	var rating []*Ratings

	for _, r := range res.Ratings {
		rating = append(rating, &Ratings{Source: r.Source, Value: r.Value})
	}
	return &MovieDetail{
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
		DVD:        res.Dvd,
		BoxOffice:  res.BoxOffice,
		Production: res.Production,
		Website:    res.Website,
	}
}
