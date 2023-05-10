package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/deyvidm/sms-backend/controllers"
	"github.com/deyvidm/sms-backend/middleware"
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
	if err := models.SetupDB(db); err != nil {
		log.Fatal(err)
	} else {
		log.Infof("DB setup successful")
	}

	router := gin.Default()
	public := router.Group("/user")
	assignPublicRoutes(public)

	private := router.Group("/api")
	private.Use(middleware.AuthJWT())
	assignPrivateRoutes(private)

	router.Run(":8080")
}

func assignPublicRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}

func assignPrivateRoutes(router *gin.RouterGroup) {
	router.GET("/user", controllers.CurrentUser)
	router.POST("/contacts/new", controllers.NewContact)
	router.GET("/contacts", controllers.AllContacts)

	router.POST("/events/new", controllers.NewEvent)
	router.GET("/events", controllers.Events)
}
