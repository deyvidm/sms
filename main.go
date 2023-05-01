package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deyvidm/sms-backend/controllers"
	"github.com/deyvidm/sms-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env vars and connect to DB
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, cleanup := models.ConnectDB(os.Getenv("DB_FILE"))
	defer cleanup()

	// Auto Migrate the struct
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Error migrating!")
	} else {
		fmt.Println("AutoMigrate successful")
	}

	router := gin.Default()

	public := router.Group("/api")
	public.POST("/login", controllers.SignIn)

	router.Run(":8080")
}
