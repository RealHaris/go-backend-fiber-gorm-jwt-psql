package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/RealHaris/go-fiber-backend/internal/database"
	"github.com/RealHaris/go-fiber-backend/internal/handlers"
	"github.com/RealHaris/go-fiber-backend/internal/middleware"
)

// @title Fiber API
// @version 1.0
// @description This is a sample server for a Fiber application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Start() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	// Initialize Fiber app
	app := fiber.New()

	// Connect to database
	db := database.ConnectDB()


  // In your Start function, add CORS middleware
  app.Use(cors.New(cors.Config{
    AllowOrigins: "*", // Adjust as necessary
    AllowHeaders: "Origin, Content-Type, Accept",
  }))

	// Setup routes
	setupRoutes(app, db)

	// Start server
	log.Info().Msg("Server running on port 3000")
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App, db *gorm.DB) {
	// Public routes
	// @Summary Register a new user
	// @Description Register a new user with username and password
	// @Accept json
	// @Produce json
	// @Param user body dto.RegisterRequest true "User registration details"
	// @Success 201 {object} dto.RegisterResponse
	// @Failure 400 {object} errors.ErrorResponse
	// @Router /register [post]
	app.Post("/register", handlers.Register(db))

	// @Summary Login a user
	// @Description Login a user with username and password
	// @Accept json
	// @Produce json
	// @Param user body dto.LoginRequest true "User login details"
	// @Success 200 {object} dto.LoginResponse
	// @Failure 400 {object} errors.ErrorResponse
	// @Failure 401 {object} errors.ErrorResponse
	// @Router /login [post]
	app.Post("/login", handlers.Login(db))

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	

	// Protected routes with JWT middleware
	app.Use(middleware.JWTMiddleware)

	// @Summary Get user profile
	// @Description Get the profile of the logged-in user
	// @Produce json
	// @Success 200 {object} models.User
	// @Failure 401 {object} errors.ErrorResponse
	// @Router /profile [get]
	app.Get("/profile", handlers.Profile(db))
}
