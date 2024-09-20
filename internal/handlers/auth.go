package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/RealHaris/go-fiber-backend/internal/dto"
	"github.com/RealHaris/go-fiber-backend/internal/errors"
	"github.com/RealHaris/go-fiber-backend/internal/models"
	"github.com/RealHaris/go-fiber-backend/internal/services"
)

// Register handler
// @Summary Register a new user
// @Description Register a new user with username and password
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User registration details"
// @Success 201 {object} dto.RegisterResponse
// @Failure 400 {object} errors.ErrorResponse
// @Router /register [post]
func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors.NewErrorResponse("Invalid request"))
		}

		// Validate request
		if err := services.ValidateRegister(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors.NewErrorResponse(err.Error()))
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errors.NewErrorResponse("Could not hash password"))
		}

		user := models.User{
			Username: req.Username,
			Password: string(hashedPassword),
		}

		// Store user in the database
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errors.NewErrorResponse("Could not create user"))
		}

		return c.Status(fiber.StatusCreated).JSON(dto.RegisterResponse{Message: "User registered successfully"})
	}
}

// Login handler
// @Summary Login a user
// @Description Login a user with username and password
// @Accept json
// @Produce json
// @Param user body dto.LoginRequest true "User login details"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Router /login [post]
func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors.NewErrorResponse("Invalid request"))
		}

		// Find user by username
		var user models.User
		if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(errors.NewErrorResponse("Invalid credentials"))
		}

		// Compare passwords
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(errors.NewErrorResponse("Invalid credentials"))
		}

		return c.Status(fiber.StatusOK).JSON(dto.LoginResponse{
			Message: "Login successful",
			Token:   services.GenerateJWT(user.ID),
		})
	}
}

