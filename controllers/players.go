package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/patterns/services/player"
)

// GetPlayers returns json with sample
func GetPlayers(w http.ResponseWriter, req *http.Request) {
	var playerType player.Sport

	vars := mux.Vars(req)
	switch vars["type"] {
	case "foot":
		playerType = &player.Foot{}
	case "basket":
		playerType = &player.Basket{}
	}

	handler := player.Handler{
		Writer:  w,
		Request: req,
	}

	context := player.NewContext(playerType, handler)
	context.Display()
}
