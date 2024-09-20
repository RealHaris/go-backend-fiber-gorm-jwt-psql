package dto

// RegisterRequest DTO
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterResponse DTO
type RegisterResponse struct {
	Message string `json:"message"`
}

// LoginRequest DTO
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse DTO
type LoginResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
}
