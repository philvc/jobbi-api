package database

import (
	"github.com/nightborn-be/invoice-backend/config"
	"github.com/nightborn-be/invoice-backend/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Default(config config.Config) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(config.Database.ConnectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Client{}, &model.Invoice{}, &model.Organisation{}, &model.User{})

	return db
}

func TestDefault() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))

	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Client{}, &model.Invoice{}, &model.Organisation{}, &model.User{})

	return db
}
