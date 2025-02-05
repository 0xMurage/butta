package response

import (
	"butta/pkg/http/route"
	"net/http"
)

type Serializer interface {
	Serialize(body any, isError bool) ([]byte, error)
	WriteError(w http.ResponseWriter, statusCode route.StatusCode, message string) (int, error)
	ContentType() string
}
