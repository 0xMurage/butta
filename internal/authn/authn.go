package authn

import (
	"butta/internal/pkg/config"
	"butta/pkg/http/request"
	"butta/pkg/http/route"
	"butta/pkg/logger"
	"butta/pkg/queue"
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

	//todo Validate the parsed credentials here

	//try to authenticate with the creds
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

func (handler *HttpHandler) ForgotPassword(req *http.Request) *route.HandlerOutput {

	var body *ForgotPasswordDto

	err := request.JsonDeserializer(req.Body, body, true)

	if err != nil {
		// If an error occurs during deserialization
		logger.Error("password reset body deserialization error", "error", err.Error())

		return &route.HandlerOutput{
			StatusCode: http.StatusBadRequest, //can do better by determining if it's user or server error
			Body:       map[string]string{"message": "Unable to process your request"},
		}
	}

	//todo validations goes in here. Check if body.username is not empty and user exists?

	//generate a magic password reset link
	magicLink, err := generatePasswordResetLink(handler.config, body.Username)

	if err != nil {
		logger.Error("error generating password reset link", "error", err)
		return &route.HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       map[string]string{"message": "Unable to process your request, try again."},
		}
	}
	// add job to the queue so send the link to the user
	args := &SendPasswordResetLinkJobArgs{
		Email: body.Username,
		Link:  magicLink,
	}

	_, err = queue.With(handler.psqlPool).Queue(req.Context(), args, nil)

	if err != nil {
		logger.Error("error adding password reset mailer job to queue", "error", err)
		return &route.HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       map[string]string{"message": "Unable to process your request, try again."},
		}
	}

	return &route.HandlerOutput{
		StatusCode: http.StatusOK,
	}
}

func New(cf config.Config, psqlPool *pgxpool.Pool) *HttpHandler {
	return &HttpHandler{
		psqlPool: psqlPool,
		config:   cf,
	}
}
