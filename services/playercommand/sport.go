package playercommand

import (
	"net/http"

	"github.com/wyllisMonteiro/patterns/models"
	"github.com/wyllisMonteiro/patterns/services"
)

// SportImpl Implements Sport interface
type SportImpl struct {
	Handler Handler
}

// foot Returns foot player
func (sportImpl *SportImpl) foot() {
	players := models.GetFootPlayers()
	services.WriteJSON(sportImpl.Handler.Writer, http.StatusOK, players)
}

// foot Returns basket player
func (sportImpl *SportImpl) basket() {
	players := models.GetBasketPlayers()
	services.WriteJSON(sportImpl.Handler.Writer, http.StatusOK, players)
}
