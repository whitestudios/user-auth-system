package config

import (
	"os"

	"github.com/whitestudios/user-auth-system/internal/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbDir := "./db"
	dbPath := "./db/users.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		err := os.MkdirAll(dbDir, os.ModePerm)

		if err != nil {
			logger.Error("Error creating db dir")
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Error("Error creating db file")
			return nil, err
		}

		file.Close()
	}

	// db connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Error("error opening sqlite: ", err.Error())
		return nil, err
	}

	// Migrate Users
	err = db.AutoMigrate(&user.User{})

	if err != nil {
		logger.Errorf("error migrating sqlite: %v", err.Error())
		return nil, err
	}

	// return the DB
	return db, nil

}
