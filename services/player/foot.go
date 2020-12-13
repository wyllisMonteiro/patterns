package player

import (
	"net/http"

	"github.com/wyllisMonteiro/patterns/models"
	"github.com/wyllisMonteiro/patterns/services"
)

// Foot Implements interface sport.go
type Foot struct{}

// Display Display foot player
func (foot *Foot) Display(handler Handler) {
	players := models.GetFootPlayers()
	services.WriteJSON(handler.Writer, http.StatusOK, players)
}
