package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/patterns/services/playercommand"
	"github.com/wyllisMonteiro/patterns/services/playerstrategy"
)

// GetPlayers returns json with players
func GetPlayers(w http.ResponseWriter, req *http.Request) {
	getPlayersWithStrategyPattern(w, req)
	//getPlayersWithCommandPattern(w, req)
}

func getPlayersWithStrategyPattern(w http.ResponseWriter, req *http.Request) {
	var playerType playerstrategy.Sport

	vars := mux.Vars(req)
	switch vars["type"] {
	case "foot":
		playerType = &playerstrategy.Foot{}
	case "basket":
		playerType = &playerstrategy.Basket{}
	}

	handler := playerstrategy.Handler{
		Writer:  w,
		Request: req,
	}

	context := playerstrategy.NewContext(playerType, handler)
	context.Display()
}

func getPlayersWithCommandPattern(w http.ResponseWriter, req *http.Request) {
	sport := &playercommand.SportImpl{
		Handler: playercommand.Handler{
			Writer:  w,
			Request: req,
		},
	}

	var players *playercommand.Players

	vars := mux.Vars(req)
	switch vars["type"] {
	case "foot":
		onFoot := &playercommand.OnFoot{
			Sport: sport,
		}

		players = &playercommand.Players{
			Command: onFoot,
		}
	case "basket":
		onBasket := &playercommand.OnBasket{
			Sport: sport,
		}

		players = &playercommand.Players{
			Command: onBasket,
		}
	}

	players.Display()
}
