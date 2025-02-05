package router

import (
	"net/http"
)

// MiddlewareFunc is a type that represents a middleware function.
// It takes a http.Handler and returns a http.HandlerFunc.
type MiddlewareFunc func(http.Handler) http.HandlerFunc
