package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adibaulia/stockbit-test/3/imdb/api/v1"
	"github.com/adibaulia/stockbit-test/3/imdb/config"
	"github.com/adibaulia/stockbit-test/3/imdb/db"
	"github.com/adibaulia/stockbit-test/3/imdb/services"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	e := echo.New()
	c := config.GetConnection()

	dao := db.NewMovieDB(c.Api)

	svc := services.NewMovieService(dao)

	api.Route(e, svc)

	lis, err := net.Listen("tcp", ":3232")
	if err != nil {
		log.Fatalf("failed to listen port %v, cause: %v", ":3232", err)
		return
	}
	s := grpc.NewServer()
	api.RegisterMovieServer(s, &api.GRPC{svc})
	go s.Serve(lis)

	go func() {
		if err := e.Start(":3033"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
