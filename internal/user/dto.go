package user

// CreateDto for creating a new user
type CreateDto struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"omitempty"`
	Email     string `json:"email" validate:"required,email"`
}

// UpdateDto  for updating a user, all fields are optional
type UpdateDto struct {
	Firstname string `json:"firstname" validate:"omitempty"`
	Lastname  string `json:"lastname" validate:"omitempty"`
	Email     string `json:"email" validate:"omitempty,email"`
}

// UserDTO for responses, containing all fields
type UserDTO struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
