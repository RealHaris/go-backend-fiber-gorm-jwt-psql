package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"

)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// JWTMiddleware validates JWT tokens
func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	c.Locals("user", token)
	return c.Next()
}
