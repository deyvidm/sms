package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// This function applies/migrates all the datbase models
// It's vital to update MODELS each time we add a new DB table/model
func SetupDB(db *gorm.DB) error {
	MODELS := []interface{}{
		User{},
		Contact{},
		Event{},
		Invite{},
	}

	for _, m := range MODELS {
		log.Info("Auto migrating ", m)
		if err := db.AutoMigrate(m); err != nil {
			return err
		}
	}
	return nil
}

var DB *gorm.DB

func ConnectDB(dbFile string) (*gorm.DB, func()) {
	// Open a database connection
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Return a closure to close the database connection
	cleanup := func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}

	DB = db
	return db, cleanup
}
