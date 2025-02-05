package router

import (
	"butta/pkg/http/response"
	"butta/pkg/http/route"
	"net/http"
)

type CustomRoute struct {
	handler                   route.HandlerFunc
	handlerResponseSerializer response.Serializer
	middleware                []MiddlewareFunc
}

// ServeHTTP is the HTTP handler for the customRoute, processing incoming requests.
// It checks if there are any middleware to be executed, and if so, it composes them in order and executes them.
// If there are no middleware, it directly calls the handler of the route to process the request.
// Parameters:
//
//	w: The ResponseWriter interface used to write responses.
//	req: The Request object representing the HTTP request received.
func (route *CustomRoute) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	handler := HandlerFuncWithSerialization(route.handlerResponseSerializer, route.handler(req))

	// Get the total number of middleware.
	total := len(route.middleware)
	// If there are no middleware, directly call the route's handler to process the request.
	if total == 0 {
		handler.ServeHTTP(w, req)
		return
	}

	//order reversed so that the execution starts from middleware on first index
	// Start composing the handlers from the last middleware, gradually moving forward.
	handler = route.middleware[total-1](handler)
	for i := total - 2; i >= 0; i-- {
		handler = route.middleware[i](handler)
	}

	// Execute the final handler
	// At this point, handler may be the original route handler or the composition of multiple middleware.
	handler.ServeHTTP(w, req)
}

// Use adds one or more middleware functions to the current route.
// This method accepts a variadic argument of middleware functions and appends them to the route's middleware slice.
// It returns a pointer to the current CustomRoute instance, enabling method chaining.
func (route *CustomRoute) Use(middleware ...MiddlewareFunc) *CustomRoute {
	// Append the provided middleware functions to the existing middleware slice of the route.
	route.middleware = append(route.middleware, middleware...)
	// Return the current route instance to support method chaining.
	return route
}

func (route *CustomRoute) WithSerializer(serializer response.Serializer) {
	route.handlerResponseSerializer = serializer
}

func NewCustomJsonRoute(routeHandler route.HandlerFunc, middleware []MiddlewareFunc) *CustomRoute {
	return &CustomRoute{
		handler:                   routeHandler,
		middleware:                middleware,
		handlerResponseSerializer: &response.JsonSerializer{},
	}
}

// CustomServeMux re-implementation of http server mux with additional methods
type CustomServeMux struct {
	mux        *http.ServeMux
	middleware []MiddlewareFunc
}

// ServeHTTP is the HTTP handler for the CustomServeMux, processing incoming requests.
// It delegates the request to the appropriate CustomRoute based on the request method and endpoint.
func (c *CustomServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Delegate the request to the underlying http.ServeMux
	c.mux.ServeHTTP(w, req)
}

// Use add middleware that will be applied to this server mux
func (c *CustomServeMux) Use(middleware ...MiddlewareFunc) {
	c.middleware = append(c.middleware, middleware...)
}

func (c *CustomServeMux) Get(endpoint string, routeHandler route.HandlerFunc) *CustomRoute {

	routed := NewCustomJsonRoute(routeHandler, c.middleware)

	c.mux.Handle("GET "+endpoint, routed)

	return routed
}

func (c *CustomServeMux) Post(endpoint string, routeHandler route.HandlerFunc) *CustomRoute {

	routed := NewCustomJsonRoute(routeHandler, c.middleware)

	c.mux.Handle("POST "+endpoint, routed)

	return routed
}

func (c *CustomServeMux) Put(endpoint string, routeHandler route.HandlerFunc) *CustomRoute {

	routed := NewCustomJsonRoute(routeHandler, c.middleware)

	c.mux.Handle("PUT "+endpoint, routed)

	return routed
}

func (c *CustomServeMux) Delete(endpoint string, routeHandler route.HandlerFunc) *CustomRoute {

	routed := NewCustomJsonRoute(routeHandler, c.middleware)

	c.mux.Handle("DELETE "+endpoint, routed)

	return routed
}

func (c *CustomServeMux) Any(endpoint string, routeHandler route.HandlerFunc) *CustomRoute {

	routed := NewCustomJsonRoute(routeHandler, c.middleware)

	c.mux.Handle(endpoint, routed)

	return routed
}

func NewServeMux() *CustomServeMux {
	return &CustomServeMux{
		mux:        http.NewServeMux(),
		middleware: []MiddlewareFunc{},
	}
}
