package authn

import (
	"butta/internal/pkg/config"
	"butta/pkg/http/request"
	"butta/pkg/http/route"
	"butta/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type HttpHandler struct {
	psqlPool *pgxpool.Pool
	config   config.Config
}

// Login handles the login request.
func (handler *HttpHandler) Login(req *http.Request) *route.HandlerOutput {
	// Decode the request body to extract credentials.
	credentials := &BasicAuthCredentialsDto{}
	err := request.JsonDeserializer(req.Body, credentials, true)

	if err != nil {
		// If an error occurs during deserialization
		logger.Error("login credentials deserialization error", "error", err.Error())

		return &route.HandlerOutput{
			StatusCode: http.StatusBadRequest, //can do better by determining if it's user or server error
			Body:       map[string]string{"message": "Unable to process your request"},
		}
	}

	// Validate the parsed credentials here
	sessionId, err := authWithBasicCredentials(req.Context(), handler.config.Session, handler.psqlPool, credentials)

	if err != nil {
		logger.Error("login credentials doesn't match", "error", err.Error())
		return &route.HandlerOutput{
			StatusCode: http.StatusUnauthorized,
			Body:       map[string]string{"message": "Username or password invalid"},
		}
	}

	// Return a successful response if no errors occur.
	return &route.HandlerOutput{
		StatusCode: http.StatusOK,
		Cookies: []*http.Cookie{
			{Name: "SESSION_COOKIE", Value: sessionId, Secure: true},
		},
	}
}

func New(cf config.Config, psqlPool *pgxpool.Pool) *HttpHandler {
	return &HttpHandler{
		psqlPool: psqlPool,
		config:   cf,
	}
}
