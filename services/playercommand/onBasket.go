package playercommand

// OnBasket is sport type
type OnBasket struct {
	Sport Sport
}

// execute When executed in asker, return basket players
func (c *OnBasket) execute() {
	c.Sport.basket()
}
