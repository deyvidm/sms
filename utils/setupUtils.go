package utils

import (
	"os"

	"github.com/deyvidm/sms-backend/middleware"
	"github.com/deyvidm/sms-backend/models"
	"github.com/deyvidm/sms-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	public := router.Group("")
	routes.AssignPublicRoutes(public)

	private := router.Group("")
	private.Use(middleware.AuthJWT())
	routes.AssignPrivateRoutes(private)

	return router
}

// envName is the Environment Variable that holds the SQLite3 database filename
// we use this function to set up automated tests as well, allowing us to pass in a separate DB for testing
// it returns a databse cleanup() function that needs to be deferred from main()
func SetupDB(envFile, envName string) func() {
	loadEnv(envFile)
	dbFilePath := os.Getenv(envName)
	if len(dbFilePath) < 1 {
		log.Warnf("MISSING DB FILENAME: |%s| SQLITE WILL RUN IN AMNESIA MODE; YOUR DATA WILL BE LOST AFTER PROGRAM TERMINATION", dbFilePath)
	} else {
		log.Infof("Using Database file: %s", dbFilePath)
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

func loadEnv(envFile string) {
	// Load Env vars and connect to DB
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}
