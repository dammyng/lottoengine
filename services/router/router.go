package router

import (
	"lottoengine/internals/persistence"
	"lottoengine/services/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(db persistence.Handler) *mux.Router {
	var r = mux.NewRouter()
	gameHandlers := handlers.NewGameHttpHandler(db)
	r.HandleFunc("/play", gameHandlers.PlayGame).Methods("POST, OPTIONS")
	r.HandleFunc("/games", gameHandlers.GetGames).Methods("GET, OPTIONS")
	r.HandleFunc("/games/{id}", gameHandlers.GetGame).Methods("PUT, OPTIONS")
	r.HandleFunc("/games/{id}", gameHandlers.UpdateGame).Methods("GET, OPTIONS")
	r.HandleFunc("/logs", gameHandlers.GameLogs).Methods("GET, OPTIONS")
	return r
}
