package route

import (
	"net/http"
)

type HandlerOutput struct {
	StatusCode StatusCode
	Body       any
	Cookies    []*http.Cookie
}

type HandlerFunc func(r *http.Request) *HandlerOutput

type HandlerFuncWithParam[T any] func(r *http.Request, param T) *HandlerOutput

func WithParam[T any](handler HandlerFuncWithParam[T], param T) HandlerFunc {
	return func(r *http.Request) *HandlerOutput {
		return handler(r, param)
	}
}
