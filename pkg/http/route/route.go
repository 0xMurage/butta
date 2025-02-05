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
