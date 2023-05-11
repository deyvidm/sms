package main

import setuputils "github.com/deyvidm/sms-backend/setupUtils"

func main() {
	cleanupDB := setuputils.SetupDB("DB_FILE")
	defer cleanupDB()
	router := setuputils.SetupRouter()
	router.Run(":8080")
}
