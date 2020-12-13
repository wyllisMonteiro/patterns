package playercommand

// OnFoot is sport type
type OnFoot struct {
	Sport Sport
}

// execute When executed in asker, return foot players
func (c *OnFoot) execute() {
	c.Sport.foot()
}
