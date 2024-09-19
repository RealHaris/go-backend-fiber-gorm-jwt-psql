package handlers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

    "github.com/RealHaris/go-fiber-backend/database"
    "github.com/RealHaris/go-fiber-backend/models"

)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Register(c *fiber.Ctx) error {
    var data map[string]string

    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

    user := models.User{
        Username: data["username"],
        Password: string(password),
    }

    if err := database.DB.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create user"})
    }

    return c.JSON(fiber.Map{"message": "registration successful"})
}

func Login(c *fiber.Ctx) error {
    var data map[string]string

    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    var user models.User
    database.DB.Where("username = ?", data["username"]).First(&user)

    if user.ID == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid credentials"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid credentials"})
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not login"})
    }

    return c.JSON(fiber.Map{"token": tokenString})
}

func Profile(c *fiber.Ctx) error {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userId := claims["user_id"]

    var userInfo models.User
    database.DB.First(&userInfo, userId)

    return c.JSON(fiber.Map{"username": userInfo.Username, "id": userInfo.ID})
}
