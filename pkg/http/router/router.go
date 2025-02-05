package router

import (
	"net/http"
)

type CustomRoute struct {
	handler    http.HandlerFunc
	middleware []MiddlewareFunc
}

// ServeHTTP is the HTTP handler for the customRoute, processing incoming requests.
// It checks if there are any middleware to be executed, and if so, it composes them in order and executes them.
// If there are no middleware, it directly calls the handler of the route to process the request.
// Parameters:
//
//	w: The ResponseWriter interface used to write responses.
//	req: The Request object representing the HTTP request received.
func (route *CustomRoute) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Get the total number of middleware.
	total := len(route.middleware)
	// If there are no middleware, directly call the route's handler to process the request.
	if total == 0 {
		route.handler.ServeHTTP(w, req)
		return
	}

	//order reversed so that the execution starts from middleware on first index
	// Start composing the handlers from the last middleware, gradually moving forward.
	handler := route.middleware[total-1](route.handler)
	for i := total - 2; i >= 0; i-- {
		handler = route.middleware[i](handler)
	}

	// Execute the final handler
	// At this point, handler may be the original route handler or the composition of multiple middleware.
	handler.ServeHTTP(w, req)
}

func (route *CustomRoute) Use(middleware ...MiddlewareFunc) *CustomRoute {
	route.middleware = append(route.middleware, middleware...)
	return route
}

// CustomServeMux re-implementation of http server mux with additional methods
type CustomServeMux struct {
	mux        *http.ServeMux
	middleware []MiddlewareFunc
}

// Use add middleware that will be applied to this server mux
func (c *CustomServeMux) Use(middleware ...MiddlewareFunc) {
	c.middleware = append(c.middleware, middleware...)
}

func (c *CustomServeMux) Get(endpoint string, routeHandler http.HandlerFunc) *CustomRoute {

	routed := &CustomRoute{
		handler:    routeHandler,
		middleware: c.middleware,
	}

	c.mux.Handle("GET "+endpoint, routeHandler)

	return routed
}

func (c *CustomServeMux) Post(endpoint string, routeHandler http.HandlerFunc) *CustomRoute {

	routed := &CustomRoute{
		handler:    routeHandler,
		middleware: c.middleware,
	}

	c.mux.Handle("POST "+endpoint, routed)

	return routed
}

func (c *CustomServeMux) Put(endpoint string, routeHandler http.HandlerFunc) *CustomRoute {

	routed := &CustomRoute{
		handler:    routeHandler,
		middleware: c.middleware,
	}

	c.mux.Handle("PUT "+endpoint, routed)

	return routed
}

func (c *CustomServeMux) Delete(endpoint string, routeHandler http.HandlerFunc) *CustomRoute {

	routed := &CustomRoute{
		handler:    routeHandler,
		middleware: c.middleware,
	}

	c.mux.Handle("DELETE "+endpoint, routed)

	return routed
}

func (c *CustomServeMux) Any(endpoint string, routeHandler http.HandlerFunc) *CustomRoute {

	routed := &CustomRoute{
		handler:    routeHandler,
		middleware: c.middleware,
	}

	c.mux.Handle(endpoint, routed)

	return routed
}

func NewServeMux() *CustomServeMux {
	return &CustomServeMux{
		mux:        http.NewServeMux(),
		middleware: []MiddlewareFunc{},
	}
}
