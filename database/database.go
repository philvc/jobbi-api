package database

import (
	"github.com/philvc/jobbi-api/config"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Default(config config.Config) *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.Database.ConnectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Friendship{}, &model.Answer{}, &model.Device{}, &model.Search{}, &model.Offer{}, &model.Company{}, &model.Network{})

	return db
}

func TestDefault() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))

	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	return db
}
