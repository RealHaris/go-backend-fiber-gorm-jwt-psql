package services

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"

	"github.com/RealHaris/go-fiber-backend/internal/dto"

)

var validate = validator.New()

// ValidateRegister validates the register request
func ValidateRegister(req dto.RegisterRequest) error {
	if err := validate.Struct(&req); err != nil {
		return errors.New("validation failed")
	}
	return nil
}

// ValidateLogin validates the login request
func ValidateLogin(req dto.LoginRequest) error {
	if err := validate.Struct(&req); err != nil {
		return errors.New("validation failed")
	}
	return nil
}


// GenerateJWT
func GenerateJWT(id uint) (string) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a new claim with the user id
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id

	// Generate the token
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return ""
	}

	return t
}
