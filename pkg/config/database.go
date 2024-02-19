package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func initDatabase() {
	var err error
	Database, err = gorm.Open(postgres.Open(Environment.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
