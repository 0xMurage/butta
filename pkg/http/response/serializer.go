package response

import (
	"butta/pkg/http/route"
	"encoding/json"
	"net/http"
)

type Serializer interface {
	Serialize(body any, isError bool) ([]byte, error)
	WriteError(w http.ResponseWriter, statusCode route.StatusCode, message string) (int, error)
	ContentType() string
}

type JsonSerializer struct {
	Data   any `json:"data,omitempty"`
	Errors any `json:"errors,omitempty"`
}

func (j *JsonSerializer) Serialize(body any, isError bool) ([]byte, error) {
	if isError {
		j.Errors = body
		return json.Marshal(j)
	} else {
		j.Data = body
		return json.Marshal(j)
	}
}
func (j *JsonSerializer) ContentType() string {
	return "application/json"
}
func (j *JsonSerializer) WriteError(writer http.ResponseWriter, statusCode route.StatusCode, message string) (int, error) {
	return WriteJsonError(writer, statusCode, message)
}
