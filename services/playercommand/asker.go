package playercommand

// Players contains command interface
type Players struct {
	Command Command
}

// Display all players using command pattern
func (b *Players) Display() {
	b.Command.execute()
}
