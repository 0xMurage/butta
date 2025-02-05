package response

import (
	"butta/pkg/http/route"
	"fmt"
	"net/http"
)

func WriteJsonError(writer http.ResponseWriter, statusCode route.StatusCode, message string) (int, error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("X-Content-Type-Options", "nosniff")

	writer.WriteHeader(int(statusCode))

	return writer.Write([]byte(fmt.Sprintf(`{"errors":{"message":"%s"}}`, message)))
}
