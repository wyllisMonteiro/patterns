package player

import (
	"net/http"

	"github.com/wyllisMonteiro/patterns/models"
	"github.com/wyllisMonteiro/patterns/services"
)

// Basket Implements interface sport.go
type Basket struct{}

// Display Display basket player
func (basket *Basket) Display(handler Handler) {
	players := models.GetBasketPlayers()
	services.WriteJSON(handler.Writer, http.StatusOK, players)
}
