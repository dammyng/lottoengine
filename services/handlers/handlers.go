package handlers

import (
	"lottoengine/internals/persistence"
	"net/http"

	"github.com/dammyng/helpers/httphelpers"
)

type GameHttpHandler struct {
	Db persistence.Handler
}

func NewGameHttpHandler(db persistence.Handler) *GameHttpHandler {
	return &GameHttpHandler{
		Db: db,
	}
}

func (handler *GameHttpHandler) GetGames(w http.ResponseWriter, r *http.Request) {

	OpenCors(&w, r)
	if r.Method == "OPTIONS" {
		response := httphelpers.NewResponseData("")
		httphelpers.RespondWithText(w, http.StatusNoContent, response)
		return
	}
}

func (handler *GameHttpHandler) GetGame(w http.ResponseWriter, r *http.Request) {

	OpenCors(&w, r)
	if r.Method == "OPTIONS" {
		response := httphelpers.NewResponseData("")
		httphelpers.RespondWithText(w, http.StatusNoContent, response)
		return
	}
}

func (handler *GameHttpHandler) UpdateGame(w http.ResponseWriter, r *http.Request) {

	OpenCors(&w, r)
	if r.Method == "OPTIONS" {
		response := httphelpers.NewResponseData("")
		httphelpers.RespondWithText(w, http.StatusNoContent, response)
		return
	}
}

func (handler *GameHttpHandler) PlayGame(w http.ResponseWriter, r *http.Request) {

	OpenCors(&w, r)
	if r.Method == "OPTIONS" {
		response := httphelpers.NewResponseData("")
		httphelpers.RespondWithText(w, http.StatusNoContent, response)
		return
	}
}

func (handler *GameHttpHandler) GameLogs(w http.ResponseWriter, r *http.Request) {

	OpenCors(&w, r)
	if r.Method == "OPTIONS" {
		response := httphelpers.NewResponseData("")
		httphelpers.RespondWithText(w, http.StatusNoContent, response)
		return
	}
}
