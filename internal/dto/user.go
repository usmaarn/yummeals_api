package dto

type CreateUserRequest struct {
	FirstName   string `json:"first_name" validate:"required,min=2,max=50"`
	LastName    string `json:"last_name" validate:"required,min=2,max=50"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=15"`
	Password    string `json:"password" validate:"required,min=8,max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type UserResponse struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
