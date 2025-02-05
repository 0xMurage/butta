package router

import (
	"butta/pkg/http/response"
	"butta/pkg/http/route"
	"net/http"
)

// MiddlewareFunc is a type that represents a middleware function.
// It takes a http.Handler and returns a http.HandlerFunc.
type MiddlewareFunc func(http.Handler) http.HandlerFunc

// HandlerFuncWithSerialization returns an HTTP handler function with serialization capabilities.
// It takes a serializer for converting the response body into a format suitable for HTTP transmission,
// and a result object containing the handler's output, including status code, cookies, and body.
// The handler function sets the appropriate status code, cookies, headers, and serialized body before writing to the ResponseWriter.
func HandlerFuncWithSerialization(serializer response.Serializer, result *route.HandlerOutput) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Set the status code to the default 200 if no status code is specified
		if result.StatusCode == 0 {
			result.StatusCode = http.StatusOK
		}

		// Set any cookies
		for _, cookie := range result.Cookies {
			http.SetCookie(w, cookie)
		}

		// Set headers
		header := w.Header()
		header.Set("Content-Type", serializer.ContentType())
		header.Set("X-Content-Type-Options", "nosniff")

		var payload []byte
		var err error
		// Serialize the body only if it exists
		if result.Body != nil {
			payload, err = serializer.Serialize(result.Body, result.StatusCode >= 400)
		}

		// Handle serialization errors or write the serialized body
		if err != nil {
			_, err = serializer.WriteError(w, http.StatusInternalServerError, "Unexpected system error encountered.")
		} else {
			w.WriteHeader(int(result.StatusCode))
			_, err = w.Write(payload)
		}

		// Handle write errors
		if err != nil {
			panic(err)
		}
	}
}
