package main

import (
	"go-fiber-backend/database"
	"go-fiber-backend/handlers"
	"go-fiber-backend/middleware"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal().Msg("Error loading .env file")
    }

    // Set up Zerolog
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

    // Initialize Fiber app
    app := fiber.New()

    // Initialize database connection
    database.ConnectDB()

    // Auth routes
    app.Post("/register", handlers.Register)
    app.Post("/login", handlers.Login)

    // Protected route
    app.Use(middleware.JWTMiddleware)
    app.Get("/profile", handlers.Profile)

    log.Info().Msg("Starting server on port 3000")
    app.Listen(":3000")
}
