package playercommand

import "net/http"

// Handler Stores handler params useful to write response
type Handler struct {
	Writer  http.ResponseWriter
	Request *http.Request
}
