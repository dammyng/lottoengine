package services

import (
	"context"
	"log"
	"lottoengine/internals/persistence"
	"lottoengine/services/router"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type GameHttpService struct {
	Db     persistence.Handler
	Router *mux.Router
}

func (e *GameHttpService) InitializeDb(db persistence.Handler) error {
	e.Db = db
	return nil
}

func (e *GameHttpService) SetRoutes() {
	rr := router.InitRoutes(e.Db)
	e.Router = rr
}

func (e *GameHttpService) StartHttp(port string) error {

	server := &http.Server{
		Addr:           port,
		Handler:        e.Router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Starting HTTP server on port " + server.Addr)

		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
	return nil
}