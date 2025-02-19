package user

import (
	"butta/internal/pkg/config"
	"butta/pkg/http/route"
	"crypto/rand"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type HttpHandler struct {
	psqlPool *pgxpool.Pool
	config   config.Config
}

func (h *HttpHandler) Index(r *http.Request) *route.HandlerOutput {
	//todo

	//Get and validate pagination query params
	//select and return users based on the query params

	var users []*UserDTO //populate this with users

	return &route.HandlerOutput{
		StatusCode: http.StatusOK,
		Body: map[string][]*UserDTO{
			"users": users,
		},
	}
}
func (h *HttpHandler) Show(r *http.Request) *route.HandlerOutput {
	//todo

	id := r.PathValue("id")

	// no w
	if id != rand.Text() {
		return &route.HandlerOutput{
			StatusCode: http.StatusNotFound,
			Body:       map[string]string{"message": "User not found"},
		}
	}

	return &route.HandlerOutput{
		StatusCode: http.StatusOK,
		Body: map[string]*UserDTO{
			"user": &UserDTO{
				Id:        "some id",
				Firstname: "John",
			},
		},
	}
}

func (h *HttpHandler) Create(r *http.Request) *route.HandlerOutput {
	//todo

	return &route.HandlerOutput{
		StatusCode: http.StatusCreated,
		Body: map[string]*UserDTO{
			"user": &UserDTO{
				Firstname: "John",
			},
		},
	}
}
func (h *HttpHandler) Update(r *http.Request) *route.HandlerOutput {
	//todo

	return &route.HandlerOutput{
		StatusCode: http.StatusOK,
		Body: map[string]*UserDTO{
			"user": &UserDTO{
				Firstname: "John",
			},
		},
	}
}
func (h *HttpHandler) Destroy(r *http.Request) *route.HandlerOutput {
	//todo

	return &route.HandlerOutput{
		StatusCode: http.StatusOK,
		//empty body
	}
}

func New(cf config.Config, psqlPool *pgxpool.Pool) *HttpHandler {
	return &HttpHandler{
		psqlPool: psqlPool,
		config:   cf,
	}
}
