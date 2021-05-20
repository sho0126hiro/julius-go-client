package client

import (
	"context"
)

// Handler structure for handling
type Handler struct {
	CallMessage string   // CallMessage call message
	Do          CallFunc // Do called function
}


type CallFunc func(ctx context.Context, result *Result) error

// RegisteredHandlers Array containing all Handlers
var RegisteredHandlers []Handler

// RegisterHandler Called at initialization
func RegisterHandler(msg string, do CallFunc) {
	RegisteredHandlers = append(RegisteredHandlers, Handler{
		CallMessage: msg,
		Do:          do,
	})
}
