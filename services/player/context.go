package player

import "net/http"

// Context Get context export
// Design pattern strategy
type Context struct {
	SportType Sport
	Handler   Handler
}

// Handler Stores handler params useful to write response
type Handler struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// NewContext Init new context
func NewContext(sport Sport, handler Handler) *Context {
	return &Context{
		SportType: sport,
		Handler:   handler,
	}
}

// Display players
func (context *Context) Display() {
	context.SportType.Display(context.Handler)
}
