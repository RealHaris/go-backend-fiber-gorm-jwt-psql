package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/RealHaris/go-fiber-backend/models"

)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=" + os.Getenv("POSTGRES_HOST") + 
           " user=" + os.Getenv("POSTGRES_USER") + 
           " password=" + os.Getenv("POSTGRES_PASSWORD") + 
           " dbname=" + os.Getenv("POSTGRES_DB") + 
           " port=" + os.Getenv("POSTGRES_PORT") + 
           " sslmode=disable TimeZone=Asia/Shanghai"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database", err)
    }

    log.Println("Database connected successfully")
    DB.AutoMigrate(&models.User{})
}
