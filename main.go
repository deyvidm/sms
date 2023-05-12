package main

import utils "github.com/deyvidm/sms-backend/utils"

func main() {
	cleanupDB := utils.SetupDB(".env", "DB_FILE")
	defer cleanupDB()
	router := utils.SetupRouter()
	router.Run(":8080")
}
