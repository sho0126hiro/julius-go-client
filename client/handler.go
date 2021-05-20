package client

import (
	"context"
)

// Handler structure for handling
type Handler struct {
	Filter Filter   // ShouldExec Returns whether the handler should be executed or not.
	Do     CallFunc // Do called function
}

type CallFunc func(ctx context.Context, result *Result) error

type Filter func(ctx context.Context, result *Result) bool

// RegisteredHandlers Array containing all Handlers
var RegisteredHandlers []Handler

// RegisterHandler Called at initialization
func RegisterHandler(filter Filter, do CallFunc) {
	RegisteredHandlers = append(RegisteredHandlers, Handler{
		Filter: filter,
		Do:     do,
	})
}
