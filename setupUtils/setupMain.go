package setuputils

import (
	"os"

	"github.com/deyvidm/sms-backend/controllers"
	"github.com/deyvidm/sms-backend/middleware"
	"github.com/deyvidm/sms-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	public := router.Group("")
	assignPublicRoutes(public)

	private := router.Group("/api")
	private.Use(middleware.AuthJWT())
	assignPrivateRoutes(private)

	return router
}

func assignPublicRoutes(router *gin.RouterGroup) {
	router.GET("/ping", controllers.Pong)
	router.POST("/user/register", controllers.Register)
	router.POST("/user/login", controllers.Login)
}

func assignPrivateRoutes(router *gin.RouterGroup) {
	router.GET("/user", controllers.CurrentUser)
	router.POST("/contacts/new", controllers.NewContact)
	router.GET("/contacts", controllers.AllContacts)

	router.POST("/events/new", controllers.NewEvent)
	router.GET("/events", controllers.AllEvents)
}

// envName is the Environment Variable that holds the SQLite3 database filename
// we use this function to set up automated tests as well, allowing us to pass in a separate DB for testing
// it returns a databse cleanup() function that needs to be deferred from main()
func SetupDB(envName string) func() {
	loadEnv()
	dbFilePath := os.Getenv(envName)
	if len(dbFilePath) < 1 {
		log.Warnf("MISSING DB FILENAME: |%s| SQLITE WILL RUN IN AMNESIA MODE; YOUR DATA WILL BE LOST AFTER PROGRAM TERMINATION", dbFilePath)
	}
	db, cleanup := models.ConnectDB(dbFilePath)
	// Auto Migrate the database models
	if err := models.SetupDB(db); err != nil {
		log.Fatal(err)
	} else {
		log.Infof("DB setup successful")
	}
	return cleanup
}

func loadEnv() {
	// Load Env vars and connect to DB
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}
